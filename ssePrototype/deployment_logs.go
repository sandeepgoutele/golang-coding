package main

import (
	"fmt"
	"log"
	"os"
	"time"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"golang.org/x/exp/rand"
)

var LOG_LEVELS = [...]string{"INFO", "WARN", "ERROR", "DEBUG"}

const (
	DATASET_LOGS = "./data"
	INFO         = iota
	WARN
	ERROR
	DEBUG
)

func mock_deployment(deploymentId string) {
	gofakeit.Seed(0)
	filePath := DATASET_LOGS + "/fake_logs.log"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file.")
	}
	defer file.Close()
	log.Printf("Initiating deployment with id: %s", deploymentId)
	log.Printf("Pushing logs to: %s", filePath)
	logger := log.New(file, "", log.LstdFlags)
	for {
		level := LOG_LEVELS[rand.Intn(len(LOG_LEVELS))]
		message := gofakeit.HackerPhrase()
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logMessage := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
		logger.Println(logMessage)

		// Sleep for a random duration between log entries
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
