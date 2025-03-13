package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"blog-api/internal/delivery/http/handler"
	"blog-api/internal/delivery/http/router"
	"blog-api/internal/delivery/http/server"
	"blog-api/internal/repository"
	"blog-api/internal/usecase"
	httpPkg "blog-api/pkg/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize repository
	postRepo := repository.NewMemoryPostRepository()

	// Initialize use case
	postUseCase := usecase.NewPostUseCase(postRepo)

	// Initialize handler
	postHandler := handler.NewPostHandler(postUseCase)

	// Initialize router
	routerConfig := httpPkg.RouterConfig{
		PostHandler: postHandler,
	}
	httpRouter := router.NewRouter(routerConfig)

	// Initialize and start server
	serverConfig := server.NewServerConfig()
	httpServer := server.NewServer(httpRouter, serverConfig)

	// Handle graceful shutdown
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	log.Printf("Server starting on port %s...\n", serverConfig.Port)

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}
}
