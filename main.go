package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	generated "github.com/grevian/petstore-demo/api/go"
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

	// Create an HTTP Server that serves our generated specification
	router := generated.NewRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP Server closed unexpected! %v \n", err)
			cancel()
		}
	}()

	// When we receive a signal to shut down, gracefully terminate the HTTP Server
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := srv.Shutdown(shutdownCtx)

	if err != nil {
		fmt.Printf("Application encountered an error: %v \n", err)
	} else {
		fmt.Println("Application Stopped!")
	}
}
