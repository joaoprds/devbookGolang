package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users "))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching  user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user!"))
}
