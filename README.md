# Making a RESTful JSON API in Go

## Overview
This project demonstrates the creation of a RESTful JSON API in Go, focusing on good RESTful design practices. It covers basic web server setup, routing, handling requests, and returning JSON responses.

## Credits:
Shout-out to to Cory Lanou for explanation.

## Requirements
- Go

## Usage
1. Clone the repository.
2. Install the dependencies.
3. Build the project.
4. Run the module.

## Installation
```bash
go get github.com/gorilla/mux
```

## Running the Server
```bash
go run github.com/SlaviXG/Practice-RESTful-JSON-API-in-GO
```

## Endpoints
- `/`: Displays a welcome message.
- `/todos`: Retrieves a list of todos.
- `/todos/{todoId}`: Retrieves details of a specific todo.
- `POST /todos`: Creates a new todo.

## License
This project is licensed under the MIT License.
