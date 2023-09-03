//This script runs all the .go scripts found in the /scripts folder.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Define scriptsFolder at the top for easy modification
var scriptsFolder string = "scripts/"

func main() {
	// Setup logging
	setupLogging()

	// Get all Go scripts
	scriptsToRun := getGoScripts()
	numScripts := len(scriptsToRun)

	// Start timer
	start := time.Now()

	// Execute each script
	for i, script := range scriptsToRun {
		fmt.Printf("Executing %d/%d %s\n", i+1, numScripts, strings.ToUpper(script))
		executeScript(script)
	}

	// Print completion message
	fmt.Printf("Completed %d/%d\n", numScripts, numScripts)

	// Calculate and print elapsed time
	elapsed := time.Since(start)
	log.Printf("Time(S)= %.2f", elapsed.Seconds())
	fmt.Printf("Time(S)= %.2f\n", elapsed.Seconds())
}

// setupLogging configures logging parameters and log file
func setupLogging() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err := os.OpenFile("reg.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	log.SetOutput(logFile)
}

// getGoScripts returns a list of all Go scripts in the scriptsFolder
func getGoScripts() []string {
	files, err := os.ReadDir(scriptsFolder)
	if err != nil {
		log.Fatalf("Error reading scripts directory: %v", err)
	}

	var goScripts []string

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".go" {
			goScripts = append(goScripts, file.Name())
		}
	}
	return goScripts
}

// executeScript runs a given script using the "go run" command and prints its output in real time
func executeScript(script string) {
	cmd := exec.Command("go", "run", filepath.Join(scriptsFolder, script))

	// Redirect stdout and stderr to os.Stdout and os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Printf("Error executing the script: %v", err)
	}
}
