package main

import (
	"github.com/gonoleks/gonoleks"
)

func main() {
	// Initialize a default Gonoleks app
	app := gonoleks.Default()

	// Define a route for the GET method on the root path '/'
	app.GET("/", func(c *gonoleks.Context) {
		// Send a string response with status code 200
		c.String(200, "Hello, World!")
	})

	// Start the server on port 3000
	app.Run(":3000")
}
