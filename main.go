package main

import (
	"encoding/json"
	"fmt"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfileHandler)
	log.Println("server is running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		_ = fmt.Errorf("there is an error : %v", err)
	}

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "ok")
	if err != nil {
		_ = fmt.Errorf("there is an error : %v", err)

		return
	}
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, "{\"error\":\"method not allowed\"}")

		var t testWriter
		fmt.Fprintf(&t, "username: %s", "mobin")
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", err.Error())))

		return
	}
	var request userservice.RegisterRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}\n", err.Error())))

		return
	}
	mySqlRepo := mysql.New()
	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		AccessTokenExpireDuration, RefreshTokenExpireDuration)
	userSvc := userservice.New(authSvc, mySqlRepo)
	_, registerErr := userSvc.Register(request)
	if registerErr != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, registerErr.Error())))
		writer.Write([]byte("\n"))

		return
	}
	_, err = writer.Write([]byte(`{"message": "user created"}`))
	writer.Write([]byte("\n"))
	if err != nil {
		fmt.Println(err)

		return
	}
}

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	// check request method
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, "{\"error\":\"method not allowed\"}")
		var t testWriter
		fmt.Fprintf(&t, "error: %s not allowed", req.Method)
	}

	// extract request body
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error":"%v"}`, err.Error())))
	}

	// bind to struct from json (unmarshall)
	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)

	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		AccessTokenExpireDuration, RefreshTokenExpireDuration)
	mySqlRepo := mysql.New()
	userSvc := userservice.New(authSvc, mySqlRepo)
	resp, err := userSvc.Login(lReq)

	if err != nil {
		fmt.Fprintf(writer, `{"error":"%v"}`, err.Error())

		return
	}
	data, err = json.Marshal(&resp)
	fmt.Println(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, err.Error())))
		writer.Write([]byte("\n"))

		return
	}

	writer.Write(data)

}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, "{\"error\":\"method not allowed\"}")
	}

	// validate jwt token and retrieve userID from token payload
	authHeader := req.Header.Get("Authorization")

	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		AccessTokenExpireDuration, RefreshTokenExpireDuration)

	claims, err := authSvc.ParseToken(authHeader)
	if err != nil {
		fmt.Fprintf(writer, `{"error":"invalid token"}`)
	}
	mySqlRepo := mysql.New()
	userSvc := userservice.New(authSvc, mySqlRepo)

	resp, err := userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		fmt.Fprintf(writer, `{"error":"%v"}`, err.Error())
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error":"%v"}`, err.Error())))
		writer.Write([]byte("\n"))
		return
	}

	writer.Write(data)

}

type testWriter struct {
	data string
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.data = string(p)
	return len(p), nil
}
