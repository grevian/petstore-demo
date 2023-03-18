package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	openapi "github.com/grevian/petstore-demo/api/go"
	"github.com/grevian/petstore-demo/service"
)

func main() {
	fmt.Println("Application Started!")

	// Create a signal handler that waits until the service receives a stop signal
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt)
		<-stop
		cancel()
	}()

	// Construct an API implementation using the default generated implementation
	PetApiService := service.NewPetAPIService()
	PetApiController := openapi.NewPetApiController(PetApiService)

	StoreApiService := service.NewStoreApiService()
	StoreApiController := openapi.NewStoreApiController(StoreApiService)

	UserApiService := service.NewUserApiService()
	UserApiController := openapi.NewUserApiController(UserApiService)

	router := openapi.NewRouter(PetApiController, StoreApiController, UserApiController)
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the HTTP Server
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// If our signal handler was invoked, gracefully shutdown the HTTP server
	<-ctx.Done()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownCancel()
	err := srv.Shutdown(shutdownCtx)
	if err != nil {
		fmt.Println("An error occurred while shutting down the server: " + err.Error())
	}

	fmt.Println("Application Stopped!")
}
