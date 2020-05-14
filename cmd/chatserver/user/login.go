package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gtinside/gochat/internal/pkg/modal"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user modal.User
	err := decoder.Decode(&user)
	if err != nil {
		handleError(err, w, "Error decoding user, try again later")
		return
	}
	table := awsSession.Table("UserDetails")
	var result modal.User
	err = table.Get("Email", user.Email).Index("EmailIndex").One(&result)
	if err != nil {
		handleError(err, w, "Invalid username/password")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)) != nil {
		handleError(errors.New("invalid username/password"), w, "Invalid username/password")
		return
	}

	log.Printf("User %v logged in successfully with UserId %v", result.Email, result.UserId)
	u, _ := json.Marshal(&result.Friend)
	response := &modal.Response{Status:modal.COMPLETED, Msg: string(u)}
	res, _ := json.Marshal(response)
	fmt.Fprintf(w, string(res))
}
