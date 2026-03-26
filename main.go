package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/repository/mysql"
	"github.com/mobin-alz/gameapp/service/userservice"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)

	log.Println("server is running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	fmt.Println("hello")
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
	userSvc := userservice.New(mySqlRepo)
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

type testWriter struct {
	data string
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.data = string(p)
	return len(p), nil
}

func testUserMysqlRepo() {
	mysqlRepo := mysql.New()
	response, err := mysqlRepo.RegisterUser(entity.User{
		PhoneNumber: "09015037617",
		Name:        "Mobin Alizadeh",
	})
	if err != nil {
		_ = fmt.Errorf("error : %v\n", err)
	} else {
		fmt.Println("User created :", response)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(response.PhoneNumber)
	if err != nil {
		_ = fmt.Errorf("error : %v\n", err)
	} else {
		fmt.Println(isUnique)
	}
}
