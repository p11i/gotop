package main

import (
	"fmt"
	"log"

	"github.com/gotop/internal/app"
)

func main() {
	fmt.Println("Welcome to gotop!")

	// Initialize your application
	myApp := app.NewApp()

	// Run the application
	if err := myApp.Run(); err != nil {
		log.Fatal("Error:", err)
	}
}
