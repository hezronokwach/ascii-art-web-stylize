package handlers

import (
	"fmt"
	"net/http"

	asciiart "asciiart/functionFiles"
)

const (
	notFound            = http.StatusNotFound
	internalServerError = http.StatusInternalServerError
	methodNotAllowed    = http.StatusMethodNotAllowed
	badRequest          = http.StatusBadRequest
)
/* 
handleError simplifies sending HTTP error responses with formatted messages.
*/
func handleError(writer http.ResponseWriter, statusCode int, message string) {
	http.Error(writer, fmt.Sprintf("%d %s", statusCode, message), statusCode)
}
/* 
Request handles the GET request for the root path by loading and executing the main template.
*/
func Request(writer http.ResponseWriter, reader *http.Request) {
	if reader.URL.Path != "/" {
		handleError(writer, notFound, "Page not found")
		return
	}
	if reader.Method != http.MethodGet {
		handleError(writer, methodNotAllowed, "Method not allowed")
		return
	}
	temp := GetTemplate()
	temp.Execute(writer, Data{Success: false})
	fmt.Println("GET / - 200 OK") // Log success in the terminal
}

/* 
Post handles the POST request to '/ascii-art' by creating ASCII art based on user input and selected banner.
*/
func Post(writer http.ResponseWriter, reader *http.Request) {
	if reader.Method != http.MethodPost {
		handleError(writer, methodNotAllowed, "Method not allowed")
		return
	}

	userInput := reader.FormValue("text")
	banner := reader.FormValue("banner")
	characterMap, err := asciiart.CreateMap(banner)
	if err != nil {
		handleError(writer, internalServerError, "Internal Server Error")
		fmt.Printf("Error creating map: %s\n", err)
		return
	}

	result := asciiart.DisplayAsciiArt(characterMap, userInput)
	if result == "" {
		handleError(writer, badRequest, "Bad Request")
		fmt.Println("Character not found")
		return
	}

	temp := GetTemplate()
	temp.Execute(writer, Data{Success: true, Result: result, UserInput: userInput})
	fmt.Println("POST /ascii-art - 200 OK ") // Log success with input data
}
