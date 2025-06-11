package main

import (
	"github.com/gonoleks/gonoleks"
)

func main() {
	// Initialize custom settings
	// NOTE: The values below represent the default configuration
	settings := &gonoleks.Settings{
		ServerName:                    "Gonoleks",
		MaxProcs:                      0,
		ReduceMemoryUsage:             false,
		GetOnly:                       false,
		CaseInSensitive:               false,
		CacheSize:                     10000,
		DisableCaching:                false,
		CacheMethods:                  "GET,POST,HEAD,OPTIONS",
		MaxCachedParams:               10,
		MaxCachedPathLength:           100,
		ReadBufferSize:                8192 * 2,
		HandleMethodNotAllowed:        false,
		HandleOPTIONS:                 false,
		AutoRecover:                   false,
		MaxRequestBodySize:            4 * 1024 * 1024,
		MaxRouteParams:                1024,
		MaxRequestURLLength:           2048,
		Concurrency:                   512 * 1024,
		Prefork:                       false,
		DisableStartupMessage:         false,
		DisableKeepalive:              false,
		DisableDefaultDate:            false,
		DisableDefaultContentType:     false,
		DisableHeaderNamesNormalizing: true,
		ReadTimeout:                   0,
		WriteTimeout:                  0,
		IdleTimeout:                   0,
		TLSEnabled:                    false,
		TLSCertPath:                   "",
		TLSKeyPath:                    "",
		DisableLogging:                false,
		LogTimeFormat:                 "2006/01/02 15:04:05",
		LogPrefix:                     "",
		LogReportCaller:               false,
		LogReportHost:                 false,
		LogReportIP:                   false,
		LogReportUserAgent:            false,
		LogReportRequestHeaders:       false,
		LogReportRequestBody:          false,
		LogReportResponseError:        false,
	}

	// Initialize a new Gonoleks app with the settings
	app := gonoleks.New(settings)

	// Define a route for the GET method on the root path '/'
	app.GET("/", func(c *gonoleks.Context) {
		// Send a string response with status code 200
		c.String(200, "Hello, World!")
	})

	// Start the server on port 3000
	app.Run(":3000")
}
