package controllers

import (
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request){

	fmt.Println("Inside Get All Users method")
	w.Write([]byte("Get All Users"))
}

func UserLogin(w http.ResponseWriter, r *http.Request){

	fmt.Println("Inside User Login method")
	w.Write([]byte("Login User"))
}