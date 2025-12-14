package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/api_server"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.uber.org/zap"
)

func main() {
	// Start setup for docker
	//if err := utils.StartDockerCompose("../database-setup/docker-compose.yaml"); err != nil {
	//	// Decide if we want to fail hard or continue.
	//	// Since the user wants it to start "as soon as application start",
	//	// failure here likely means DB won't be ready.
	//	utils.Log.Warn("Proceeding despite docker start failure (maybe it is already running?)", zap.Error(err))
	//} else {
	//	// Give it a moment? MongoDB might take a few seconds to accept connections even if container is up.
	//	// The db.go logic has a retry/timeout mechanism, so strictly sleeping might not be necessary,
	//	// but a small pause can reduce initial connection error logs.
	//	time.Sleep(2 * time.Second)
	//}

	ginServer := api_server.NewGinServer()
	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        ginServer.GetHandler(),
		ReadTimeout:    time.Duration(20) * time.Second,
		WriteTimeout:   time.Duration(20) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		utils.Log.Log(zap.ErrorLevel, "error starting server", zap.Error(err))
	}
	utils.Log.Log(zap.InfoLevel, "server started on port", zap.Int("port", 8080))
}
