package main

import (
	"fmt"
	"net/http"
	"io"
)

func main(){

	http.HandleFunc("/login",signUp)
	http.HandleFunc("/signup",login)
 err:= http.ListenAndServe(":3000",nil);


 if err!=nil{
	fmt.Println("sorry error occured while server occurances");
 }
 
}

func signUp(w http.ResponseWriter, r *http.Request){
  fmt.Println("backend responded with something")
  io.WriteString(w,"server is singing up with the backend") 

}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("backend responded with something")
	io.WriteString(w,"server is responding successfully witht the following command")
}