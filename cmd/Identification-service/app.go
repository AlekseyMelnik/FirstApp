package main

import (
	db "FirstProject/internal/adapters/db/postgres"
	"FirstProject/internal/config"
	"FirstProject/internal/domain/service"
	v1 "FirstProject/internal/handlers/http/v1"
	"FirstProject/pkg/db/postgres"
	"FirstProject/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create Httprouter!")
	router := httprouter.New()
	logger.Info("Create config!")
	cfg := config.GetConfig()
	logger.Info("Create client database!")

	client, err := postgres.NewClient(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)
	if err != nil {
		logger.Fatal(err)
	}
	defer client.Close()
	storage := db.NewUserStorage(client)
	userService := service.NewUserService(storage)
	handler := v1.NewUserHandler(userService)
	handler.Register(router)
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Create  Socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("Socket path: %s", socketPath)
		logger.Info("Listen UnixІ socket")
		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("Server listen %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
