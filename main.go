package main

import (
	"fmt"
	"github.com/mobin-alz/gameapp/config"
	"github.com/mobin-alz/gameapp/delivery/httpserver"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/service/authservice"
	"github.com/mobin-alz/gameapp/service/userservice"
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

	// Echo instance (engine)
	fmt.Println("Starting echo server")
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappPassword",
			Port:     3306,
			Host:     "localhost",
			DBName:   "gameapp_db",
		},
	}

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)
	server.Serve()

}

//	func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
//		// check request method
//		if req.Method != http.MethodPost {
//			fmt.Fprintf(writer, "{\"error\":\"method not allowed\"}")
//			var t testWriter
//			fmt.Fprintf(&t, "error: %s not allowed", req.Method)
//		}
//
//		// extract request body
//		data, err := io.ReadAll(req.Body)
//		if err != nil {
//			writer.Write([]byte(
//				fmt.Sprintf(`{"error":"%v"}`, err.Error())))
//		}
//
//		// bind to struct from json (unmarshall)
//		var lReq userservice.LoginRequest
//		err = json.Unmarshal(data, &lReq)
//
//		authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
//			AccessTokenExpireDuration, RefreshTokenExpireDuration)
//		mySqlRepo := mysql.New()
//		userSvc := userservice.New(authSvc, mySqlRepo)
//		resp, err := userSvc.Login(lReq)
//
//		if err != nil {
//			fmt.Fprintf(writer, `{"error":"%v"}`, err.Error())
//
//			return
//		}
//		data, err = json.Marshal(&resp)
//		fmt.Println(resp)
//		if err != nil {
//			writer.Write([]byte(fmt.Sprintf(`{"error":"%v"}`, err.Error())))
//			writer.Write([]byte("\n"))
//
//			return
//		}
//
//		writer.Write(data)
//
// }
//
//	func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
//		if req.Method != http.MethodGet {
//			fmt.Fprintf(writer, "{\"error\":\"method not allowed\"}")
//		}
//
//		// validate jwt token and retrieve userID from token payload
//		authHeader := req.Header.Get("Authorization")
//
//		authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
//			AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//		claims, err := authSvc.ParseToken(authHeader)
//		if err != nil {
//			fmt.Fprintf(writer, `{"error":"invalid token"}`)
//		}
//		mySqlRepo := mysql.New()
//		userSvc := userservice.New(authSvc, mySqlRepo)
//
//		resp, err := userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
//		if err != nil {
//			fmt.Fprintf(writer, `{"error":"%v"}`, err.Error())
//			return
//		}
//		data, err := json.Marshal(resp)
//		if err != nil {
//			writer.Write([]byte(
//				fmt.Sprintf(`{"error":"%v"}`, err.Error())))
//			writer.Write([]byte("\n"))
//			return
//		}
//
//		writer.Write(data)
//
// }
func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	repo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, repo)

	return authSvc, userSvc
}
