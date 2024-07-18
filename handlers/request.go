package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	asciiart "asciiart/functionFiles"
)

const (
	notFound            = http.StatusNotFound
	internalServerError = http.StatusInternalServerError
	methodNotAllowed    = http.StatusMethodNotAllowed
	badRequest          = http.StatusBadRequest
)

var feedback string

func handleError(writer http.ResponseWriter, statusCode int, message string) {
	// Construct the URL for the error page with query parameters
	target := fmt.Sprintf("/error?code=%d&message=%s", statusCode, url.QueryEscape(message))
	http.Redirect(writer, &http.Request{URL: &url.URL{Path: target}}, target, http.StatusSeeOther)
}

func Request(writer http.ResponseWriter, reader *http.Request) {
	if reader.URL.Path != "/" {
		handleError(writer, notFound, "Page not found")
		feedback = "This page does not exist"
		return
	}
	if reader.Method != http.MethodGet {
		handleError(writer, methodNotAllowed, "Method not allowed")
		return
	}
	tmpl := GetTemplate()
	err := tmpl.Execute(writer, Data{Success: false})
	if err != nil {
		handleError(writer, internalServerError, "Internal Server Error")
		fmt.Printf("Error executing template: %s\n", err)
	}
	fmt.Println("GET / - 200 OK") // Log success in the terminal
}

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
		feedback = fmt.Sprintf("Error loading %s banner file", banner)
		fmt.Printf("Error creating map: %s\n", err)
		return
	}

	result := asciiart.DisplayAsciiArt(characterMap, userInput)
	if result == "" {
		handleError(writer, badRequest, "Bad Request")
		feedback = "The Input must NOT contain non-ascii charcters"
		fmt.Println("Character not found")
		return
	}

	tmpl := GetTemplate()
	err = tmpl.Execute(writer, Data{Success: true, UserInput: userInput, Result: result})
	if err != nil {
		handleError(writer, internalServerError, "Internal Server Error")
		fmt.Printf("Error executing template: %s\n", err)
	}
	fmt.Println("POST /ascii-art - 200 OK ") // Log success with input data
}

func ErrorHandler(writer http.ResponseWriter, reader *http.Request) {
	statusCodeStr := reader.URL.Query().Get("code")
	statusCode, err := strconv.Atoi(statusCodeStr)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	message := reader.URL.Query().Get("message")

	err = errorTmpl.Execute(writer, Data{ErrorMessage: message, StatusCode: statusCode, Feedback: feedback})
	if err != nil {
		http.Error(writer, "Error rendering error page", http.StatusInternalServerError)
		fmt.Printf("Error executing error template: %s\n", err)
	}
}
