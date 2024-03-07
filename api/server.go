package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"sqlBoiler/util"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	s := &Server{router: gin.Default()}
	s.router.GET("/persons", GetAllPersons)
	s.router.GET("/persons/:iin", GetPersonByIIN)
	s.router.POST("/persons", CreatePerson)
	return s
}

func (s *Server) Start() error {
	config, _ := util.LoadConfig(".")
	log.Printf("Сервер запущен по адресу: %s", config.ServerAddress)
	return s.router.Run(config.ServerAddress)
}
