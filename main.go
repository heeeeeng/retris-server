package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	initURL()

	fmt.Println("---------Server Start!---------")
	fmt.Println("Port: ", 10000)
	log.Fatal(http.ListenAndServe(":10000", nil))
}


func initURL() {
	http.HandleFunc("/new_game", createNewGame)
	http.HandleFunc("/user_ctrl", handleUserCtrl)
}

func createNewGame(w http.ResponseWriter, r *http.Request) {

}

func handleUserCtrl(w http.ResponseWriter, r *http.Request) {

}


type Server struct {

}

func NewServer() *Server {

	return nil
}

func (srv *Server) Start() {}

func (srv *Server) Stop() {}


