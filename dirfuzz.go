//usr/bin/go run $0 $@ ; exit

// DIRFUZZ Written by NOCHHACKS in GOlang //
// You may need to increase your resource usage. //
// For Linux, I used "ulimit 614400" //

package main

import (
	"net/http"
	"log"
	"os"
	"bufio"
	"fmt"
	"strconv"
	"github.com/TwinProduction/go-color"
	"flag"
)

func main () {
	// INITIALISE VARIABLES //
	var url string
	var fullUrl string
	var code string
	var output string
	var filePath string
	var fileTextLines []string
	var verbose bool
	var client http.Client

	// GET FLAGS & VALUES //
	flag.StringVar(&url, "u", "http://X.X.X.X/", "Specify base URL, include http(s)://")
	flag.StringVar(&filePath, "w", "wordlist.txt", "Specify path to directory wordlist")
	flag.BoolVar(&verbose, "v", false, "Toggle verbose setting, ouputs non-200 (OK) responses.")
	flag.Parse()

	// GET LENGTH OF URL INPUT //
	runeVersion := []rune(url)
	urlLength := len(runeVersion)

	// CHECK FOR DEFAULT FLAG VALUE //
	if (url == "http://X.X.X.X/") {
		print("\n")
		flag.PrintDefaults()
		print("\n")
		os.Exit(0)	
	}

	// VALIDATE URL LENGTH //
	if (urlLength < 7) {
		print("\n")
		flag.PrintDefaults()
		print("\n")
		os.Exit(0)
	}

	// INITIATE FUZZING IF URL INCLUDES HTTP(S) //
	if (url[0:7] == "http://" || url[0:8] == "https://") {

		output = "Scanning: " + url
		print("\n")
		fmt.Println(color.Colorize(color.Blue, output))
		output = "List: " + filePath
		fmt.Println(color.Colorize(color.Blue, output))
		
		// OPEN FILEPATH //
		readFile, err := os.Open(filePath)
		// CATCH ERRORS //
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		
		// SCAN LINES IN TEXT FILE //
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		// APPEND ALL LINES TO fileTextLines //
		for fileScanner.Scan() {
			fileTextLines = append(fileTextLines, fileScanner.Text())
		}

		// CLOSE FILE //
		readFile.Close()

		// MAIN FOR LOOP TO CHECK FOR DIRECTORIES //
		print("\n")
		for _, slug := range fileTextLines {

			// CREATE FULL URL //
			if url[len(url)-1:] != "/" {
				fullUrl = string(url + "/" + slug)
			} else {
				fullUrl = string(url + slug)
			}

			// INITIATE HTTP REQUEST //
			response, error := client.Get(fullUrl)

			// CATCH ERRORS //
			if error != nil {
				log.Fatal("Connection Failed! \n", error)
			}

			defer response.Body.Close()

			// CONVERT STATUSCODE RESPONSE TO STR //
			code = strconv.Itoa(response.StatusCode)

			// FORMAT RESPONSE ACCORDINGLY //
			if response.StatusCode == http.StatusOK {
				output = "[+] " + fullUrl + " " + code + " " + http.StatusText(response.StatusCode)
				fmt.Println(color.Colorize(color.Green, output))
			} else {
				if verbose {
					output = "[X] /" + slug + " " + code + " " + http.StatusText(response.StatusCode)
					fmt.Println(color.Colorize(color.Red, output))
				}
			}
		}
		print("\n")
	} else {
		print("\n")
		flag.PrintDefaults()
		print("\n")
		os.Exit(0)
	}
}
