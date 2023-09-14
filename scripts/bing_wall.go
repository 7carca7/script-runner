// This script downloads the image of the day from Bing.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// apiURL is the URL of the TimothyYe/bing-wallpaper API that provides the Bing image of the day.
// Available regions = ["en-US", "en-CA", "en-GB", "en-AU", "en-NZ", "de-DE", "zh-CN", "ja-JP", "random"]
// Available resolutions = ["1920", "1366", "3840"]
// The index is set to 0 to get the most recent image and the market to "en-US" for images from the United States.
var apiURL string = "https://bing.biturl.top/?resolution=3840&format=json&index=0&mkt=en-US"

// Define download folder at the top for easy modification
var downloadFolder string = ""

// Response struct to map the JSON response from the API
type Response struct {
	Url       string `json:"url"`
	Copyright string `json:"copyright"`
}

// This function removes characters that are not allowed in filenames
func sanitizeFilename(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 _-]+")
	if err != nil {
		panic(err)
	}
	output := reg.ReplaceAllString(input, "")
	return output
}

func main() {
	// Send a GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Declare a new Response
	var r Response
	// Decode the JSON response into the Response struct
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(err)
	}

	// Send a GET request to the URL obtained from the previous response
	resp, err = http.Get(r.Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check if the copyright string contains a comma. If it does, split the string at the comma.
	// If it doesn't, check if the string contains a parenthesis and split the string at the parenthesis.
	// This is used to construct the filename for the downloaded image.
	var parts []string
	if strings.Contains(r.Copyright, ",") {
		parts = strings.Split(r.Copyright, ",")
	} else if strings.Contains(r.Copyright, "(") {
		parts = strings.Split(r.Copyright, "(")
	}

	// Sanitize the filename before using it.
	safeFilename := sanitizeFilename(parts[0])
	// Construct the filename using the first part of the copyright string
	filename := downloadFolder + safeFilename + ".jpg"

	// Check if the image already exists in the directory
	if _, err := os.Stat(filename); err == nil {
		// If the image exists, print a message and end execution
		fmt.Printf("The image already exists: \"%s\"\n", strings.TrimSpace(filename))
		return
	}

	// If the image does not exist, create a new file
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Copy the response body (the image) to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	// Print a message indicating that the image was downloaded successfully
	fmt.Printf("Image downloaded successfully: \"%s\"\n", strings.TrimSpace(filename))
}
