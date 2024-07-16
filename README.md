# ASCII Art Web Server
![](https://www.parallels.com/blogs/ras/app/uploads/2022/02/Best-Practices-for-Server-Performance-Monitoring.jpg)

## Description

This Go web application serves as an ASCII art generator. Users can input text and select a banner style, then the server processes this input to generate ASCII art. The application is built with the Go programming language and uses HTML templates for rendering web pages.

## Authors
- [Hezron Okwach](https://github.com/hezronokwach) 
- [Anne Maina](https://github.com/nyagooh)
- [Brian Bantu](https://github.com/Bantu-art)
## Usage

### Requirements

- Go (version 1.15 or higher recommended)
- Access to a terminal or command prompt

### Running the Application

1. Clone the repository:
   ```bash
   git clone https://learn.zone01kisumu.ke/git/hokwach/ascii-art-web
   cd ascii-art-web
   ```
2. Run the server:
   ```bash
    go run .  
    ```
3. Open a web browser and visit:
    ```bash
   http://localhost:8090
    ```

## Implementation Details

### Endpoints

  GET /: Serves the main page where users can input text and choose a banner style.

  POST /submit: Receives data from the form on the main page (text and a selected banner), generates ASCII art based on the input, and returns the result on the page.

### Features

   Text Input: Users can enter the text they want to convert to ASCII art. The text must be ascii characters otherwise they will encounter bad request error message.
    Banner Selection: Users can select the desired ASCII art style via a select dropdown i.e standard, shadow, thinkertoy.

   Dynamic Results Display: The results of the ASCII art generation are displayed on the main page after the POST request, without needing to navigate to a new page.

  ### Algorithm

The ASCII art generation algorithm involves several key steps, from receiving user input to processing it into ASCII art. Here’s a detailed breakdown:

1. **Receiving User Input**:
   - The main page (`GET /`) presents users with an HTML form where they can enter the text they wish to transform and select a banner style (e.g., standard, shadow, or thinkertoy).
   - Users submit their input and selected banner via a `POST` request to `/submit`.

2. **Form Validation**:
   - When the server receives the `POST` request, it first validates the input to ensure that the text is not empty and that the selected banner is one of the predefined options.
   - If the validation fails, the server responds with a `400 Bad Request` status code, and the user is prompted to correct their input.

3. **Processing the Input**:
   - If the input passes validation, the server uses the selected banner style to retrieve a corresponding character map. Each banner style has a different set of character designs for representing text.
   - The server processes each character of the input text through the chosen banner's character map, converting it to a stylized ASCII representation.

4. **Generating ASCII Art**:
   - The processed characters are then assembled into full lines of ASCII art, maintaining the appropriate spacing and line breaks based on the dimensions of the banner characters.
   - This step involves looping through each line of text input and each character, applying the character map, and concatenating the results into complete lines of ASCII art.

5. **Displaying the Result**:
   - After processing, the generated ASCII art is embedded into the response HTML using Go templates.
   - The ASCII art is displayed on the same main page `/submit` page.
6. **Error Handling**:
   - Throughout the process, the server checks for potential errors, such as file access errors (if character maps are stored in files) or rendering issues.
   - Any internal server errors trigger a `500 Internal Server Error` response, with appropriate logging for debugging purposes.

### HTTP Status Codes

- `200 OK`: Returned when pages and POST requests are processed successfully.
- `404 Not Found`: Returned when a requested resource (e.g., a non-existent template or banner) is not found.
- `400 Bad Request`: Returned when the client sends an incorrect request.
- `500 Internal Server Error`: Returned when the server encounters an error that it cannot handle.

### Technologies Used

- **Go**: Primary programming language for server-side logic.
- **HTML Templates**: Used for rendering the user interface.
- **Net/HTTP Package**: Used for handling HTTP requests and responses.

## Contributing

Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.