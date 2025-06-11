# Hello World Example

This project demonstrates a simple "Hello, World!" application using [Gonoleks][gonoleks_url].

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gonoleks/examples.git
    cd examples/hello-world
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `http://localhost:3000`.

## Example

Here is an example "Hello, World!" in a Gonoleks application:

```go
package main

import (
	"github.com/gonoleks/gonoleks"
)

func main() {
	app := gonoleks.New()

	app.GET("/", func(c *gonoleks.Context) {
		c.String(200, "Hello, World!")
	})

	app.Run(":3000")
}
```

<!-- README links -->

[gonoleks_url]: https://github.com/gonoleks/gonoleks