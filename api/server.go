package api

import (
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
}

// NewServer creates a new HTTP server and setup routing
func NewServer() (*Server, error) {
	// Initialize a new Fiber app
	app := fiber.New()

	server := &Server{
		app: app,
	}

	// Setup the router
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	// Define a route for the GET method on the root path '/'
	server.app.Get("/api", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World api 👋!")
	})
}

// Start runs the HTTP server on a specific address
func (server *Server) Start() error {
	return server.app.Listen(":3000")
}
