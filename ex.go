package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"regexp"
)

func main() {
	// Make a GET request to the coronavirus statistics webpage
	resp, err := http.Get("https://www.worldometers.info/coronavirus/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Find the highest number of confirmed cases in the body
	var highestConfirmed int
	startIndex := strings.Index(string(body), "Coronavirus Cases:")
	if startIndex != -1 {
		endIndex := startIndex + strings.Index(string(body)[startIndex:], "</span>")
		confirmedString := string(body)[startIndex:endIndex]
		confirmedString = strings.ReplaceAll(confirmedString, ",", "") // Remove commas from the number string
		confirmedString = strings.ReplaceAll(confirmedString, "Coronavirus Cases:", "")
		re := regexp.MustCompile(`Coronavirus Cases:.*?>([\d,]+)<`)
		fmt.Println(re)
		highestConfirmed, err = strconv.Atoi(strings.TrimSpace(confirmedString))
		if err != nil {
			panic(err)
		}
	}

	// Print the highest number of confirmed cases found on the webpage
	fmt.Printf("The highest number of confirmed cases found on the webpage is %d\n", highestConfirmed)
}