package main

import (
	"FirstProject/internal/user"
	"FirstProject/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create Httprouter!")
	router := httprouter.New()

	handler := user.NewHandler(logger)
	handler.Register(router)
	start(router)
}
func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.Serve(listener)
}
