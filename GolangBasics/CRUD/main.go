package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// User struct represents a user in the application
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DataService manages the array of users
type DataService struct {
	users []User
}

// InitializeData initializes the data array with some sample users
func (ds *DataService) InitializeData() {
	ds.users = []User{
		{Username: "user1", Password: hashPassword("password1")},
		{Username: "user2", Password: hashPassword("password2")},
	}
}

// GetAllUsers returns all users in the data array
func (ds *DataService) GetAllUsers() []User {
	return ds.users
}


func (ds *DataService) GetUserByUsername(username string) (User, bool) {
	for _, user := range ds.users {
		if user.Username == username {
			return user, true
		}
	}
	return User{}, false
}


func (ds *DataService) CreateUser(user User) {
	user.Password = hashPassword(user.Password)
	ds.users = append(ds.users, user)
}

// UpdateUser updates the password of an existing user
func (ds *DataService) UpdateUser(username, newPassword string) bool {
	for i := range ds.users {
		if ds.users[i].Username == username {
			ds.users[i].Password = hashPassword(newPassword)
			return true
		}
	}
	return false
}

// DeleteUser deletes a user by username
func (ds *DataService) DeleteUser(username string) bool {
	for i, user := range ds.users {
		if user.Username == username {
			// Remove the user from the array
			ds.users = append(ds.users[:i], ds.users[i+1:]...)
			return true
		}
	}
	return false
}

// hashPassword hashes a password using bcrypt
func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func main() {
	dataService := DataService{}
	dataService.InitializeData()

	r := mux.NewRouter()

	r.HandleFunc("/users", GetAllUsersHandler(&dataService)).Methods("GET")
	r.HandleFunc("/users/{username}", GetUserHandler(&dataService)).Methods("GET")
	r.HandleFunc("/users", CreateUserHandler(&dataService)).Methods("POST")
	r.HandleFunc("/users/{username}", UpdateUserHandler(&dataService)).Methods("PUT")
	r.HandleFunc("/users/{username}", DeleteUserHandler(&dataService)).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// GetAllUsersHandler handles GET request to retrieve all users
func GetAllUsersHandler(ds *DataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := ds.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	}
}

// GetUserHandler handles GET request to retrieve a user by username
func GetUserHandler(ds *DataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		user, found := ds.GetUserByUsername(username)
		if !found {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

// CreateUserHandler handles POST request to create a new user
func CreateUserHandler(ds *DataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ds.CreateUser(newUser)
		w.WriteHeader(http.StatusCreated)
	}
}

// UpdateUserHandler handles PUT request to update the password of an existing user
func UpdateUserHandler(ds *DataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		var newPassword struct {
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&newPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		success := ds.UpdateUser(username, newPassword.Password)
		if !success {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteUserHandler handles DELETE request to delete a user by username
func DeleteUserHandler(ds *DataService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		success := ds.DeleteUser(username)
		if !success {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
