// This script downloads the image of the day from Bing.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Response struct to map the JSON response from the API
type Response struct {
	Url       string `json:"url"`
	Copyright string `json:"copyright"`
}

func main() {
	// Send a GET request to the API
	resp, err := http.Get("https://bing.biturl.top/?resolution=3840&format=json&index=0&mkt=en-US")
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

	// Split the copyright string into parts
	parts := strings.Split(r.Copyright, ",")
	// Construct the filename using the first part of the copyright string
	filename := "/Users/ernesto/Pictures/bing wallpaper/" + parts[0] + ".jpg"

	// Check if the image already exists in the directory
	if _, err := os.Stat(filename); err == nil {
		// If the image exists, print a message and end execution
		fmt.Printf("The image already exists: \"%s\"\n", parts[0])
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
	fmt.Printf("Image downloaded successfully: %s\n", parts[0])
}
