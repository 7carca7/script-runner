// This script will run all commands in the commands list

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Commands List
	commands := []string{
		"brew update",
		"brew upgrade",
		"brew autoremove",
		"brew cleanup",
		"mas upgrade",
		"softwareupdate -ia --verbose",
	}

	for _, command := range commands {
		fmt.Printf("Executing: %s\n", command)

		// Split the command string into command and arguments
		parts := strings.Fields(command)

		cmd := exec.Command(parts[0], parts[1:]...)

		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Command '%s' failed with error: %s\n", command, err)
			continue
		}

		/*
			    Verbose output
			    output, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Printf("Command '%s' failed with error: %s\n", command, err)
						continue
					}

					fmt.Printf("%s", output)
		*/
	}
}
