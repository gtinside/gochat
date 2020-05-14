package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gtinside/gochat/internal/pkg/modal"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var awsSession = modal.GetDynamoDBSession()

// registerAction takes in a user request, validates the input and register the user
func RegisterAction(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user modal.User
	err := decoder.Decode(&user)
	if err != nil {
		handleError(err, w, "Error decoding user")
		return
	}

	table := awsSession.Table("UserDetails")
	//Step 1: Check if user exist by looking up for email id
	var results []modal.User
	err = table.Get("Email", user.Email).Index("EmailIndex").All(&results)
	if err != nil {
		log.Print(fmt.Errorf("error validating user request, err: %w", err))
		handleError(err, w, "Error validating request ")
		return
	}

	if len(results) != 0 {
		handleError(errors.New("invalid Request"), w, fmt.Sprintf("User with username: %s already exists", user.Email))
		return
	} else {
		//Step 2: Register the user, but lets process the password
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			handleError(err, w, "Error processing the password")
		}
		user.Password = string(hash)
		user.UserId = uuid.New().String()
		err = table.Put(user).Run()
		if err != nil {
			handleError(err, w, "Error processing the password")
			return
		}
		log.Printf("User %v created successfully with UserId %v", user.Email, user.UserId)
		response := &modal.Response{Status:modal.COMPLETED, Msg: ""}
		result, _ := json.Marshal(response)
		fmt.Fprintf(w, string(result))
		return
	}
}

// GetAllUsers : Return the list of users
func GetAllUsers(w http.ResponseWriter, req *http.Request)  {
	var users []modal.Friend
	table := awsSession.Table("UserDetails")
	err := table.Scan().All(&users)
	if err != nil {
		handleError(err, w, " Error retrieving users")
	}
	friends, _ := json.Marshal(users)
	response := &modal.Response{Status:modal.COMPLETED, Msg: string(friends)}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	return
}

// GetUser : Return user details fora given userId
func GetUser(w http.ResponseWriter, req *http.Request)  {
	decoder := json.NewDecoder(req.Body)
	var user modal.Friend
	err := decoder.Decode(&user)
	if err != nil {
		handleError(err, w, " Error retrieving users")
		return
	}
	var friend modal.Friend
	table := awsSession.Table("UserDetails")
	err = table.Get("UserId", user.UserId).Index("UserIdIndex").One(&friend)
	if err != nil {
		handleError(err, w, " Error retrieving users")
		return
	}
	friends, _ := json.Marshal(friend)
	response := &modal.Response{Status:modal.COMPLETED, Msg: string(friends)}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}


// HandleError : Generic method to handle errors
func handleError(err error, w http.ResponseWriter, msg string) {
	log.Print(err)
	response := &modal.Response{Status:modal.ERROR, Msg: msg}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
