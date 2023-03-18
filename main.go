package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Application Started!")

	// Create a signal handler that waits until the service receives a stop signal
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop

	fmt.Println("Application Stopped!")
}
