package main

import (
	"net/http"
	"log"

	. "./dao/userDAO"
	. "./config"

	"./router"

)


const _PORT = ":8000"
var config = Config{}
var dao = UsersDAO{}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := router.Init()

	log.Fatal(http.ListenAndServe(_PORT, r))
}

