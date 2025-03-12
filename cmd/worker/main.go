package main

import (
	"fmt"
	"os"
	"time"
)

const statusFile = "status.log"

func readStatus() {
	for {
		data, err := os.ReadFile(statusFile)
		if err != nil {
			fmt.Println("Error reading status file:", err)
			continue
		}

		if len(data) > 0 {
			fmt.Println("\n🚀 Live Delivery Updates 🚀")
			fmt.Println(string(data))

			// Clear file after reading to avoid duplicates
			os.WriteFile(statusFile, []byte{}, 0644)
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Println("📦 Delivery Tracking System Running...")
	readStatus()
}
