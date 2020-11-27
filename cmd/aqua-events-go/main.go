package main

import (
	"github.com/BryanKMorrow/aqua-events-go/src/webhooksrv"
	"log"
	"os"
)

func main() {
	log.Println("Using the environment variables")
	fatal := checkEnv()
	if fatal {
		log.Fatalln("Environment variables not set, stopping aqua-events-go")
	}
	s := webhooksrv.NewServer()
	s.Start()
}

// checkEnv looks for the required environment variables
func checkEnv() bool {
	fatal := false

	webhook := os.Getenv("SLACK_WEBHOOK")
	if webhook == "" {
		log.Println("Please set the SLACK_WEBHOOK environment variable")
		fatal = true
	}
	return fatal
}
