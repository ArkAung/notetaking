# Note-Taking App

This is a simple note-taking application backend built with Go and Gin.

## Features

- Create a new note
- Retrieve a list of all notes

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/your-username/note-taking-app.git
cd note-taking-app
```

### Install dependencies:
```bash
go get github.com/gin-gonic/gin
```

### Run the server:
```bash
go run main.go
```
The server will start at http://localhost:8080.

### Endpoints
```bash
POST /create: Create a new note.
GET /notes: Retrieve a list of all notes.
```

### Contributing
Contributions are welcome! Feel free to open issues and pull requests.