package main

import (
	"fmt"
	"log"
	"net/http"

	. "./config"
	"./dao/articleDAO"
	. "./dao/dbDAO"
	"./dao/userDAO"
	"./router"

	"github.com/rs/cors"
)


const _PORT = ":8000"
var config = Config{}
var dao = DAO{}
var articleDao = articleDAO.ArticleDAO{}
var userDao = userDAO.UsersDAO{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database

	articleDAO.Init(dao)
	userDAO.Init(dao)


	fmt.Println("Connecting to " + dao.Database + "...")
	//dao.Connect()
	fmt.Println("Connected!")
}

func main() {
	r := router.Init()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

	})
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(_PORT, handler))
}

