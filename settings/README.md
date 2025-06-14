# Settings Example

This project demonstrates an application with customized settings using [Gonoleks][gonoleks_url].

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gonoleks/examples.git
    cd examples/settings
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

Here is an example of customized settings in a Gonoleks application:

```go
package main

import (
	"github.com/gonoleks/gonoleks"
)

func main() {
	settings := &gonoleks.Settings{
		ServerName:                    "Gonoleks",
		MaxProcs:                      0,
		ReduceMemoryUsage:             false,
		GetOnly:                       false,
		CaseInSensitive:               false,
		CacheSize:                     10000,
		DisableCaching:                false,
		CacheMethods:                  "GET,HEAD",
		MaxCachedParams:               10,
		MaxCachedPathLength:           100,
		ReadBufferSize:                8192 * 2,
		HandleMethodNotAllowed:        false,
		HandleOPTIONS:                 false,
		AutoRecover:                   true,
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

	app := gonoleks.New(settings)

	app.GET("/", func(c *gonoleks.Context) {
		c.String(200, "Hello, World!")
	})

	app.Run(":3000")
}
```

## Settings Fields

Here is the full list of settings fields that can be customized:

| Property | Type | Description | Default |
|---|---|---|---|
| **`ServerName`** | string | Server name to send in response headers | `"Gonoleks"` |
| **`MaxProcs`** | int | Controls the number of operating system threads that can execute user-level Go code simultaneously. Zero means use runtime.NumCPU() | `0` |
| **`ReduceMemoryUsage`** | bool | Aggressively reduces memory usage at the cost of higher CPU usage if set to true | `false` |
| **`GetOnly`** | bool | Rejects all non-GET requests if set to true | `false` |
| **`CaseInSensitive`** | bool | Enables case-insensitive routing | `false` |
| **`CacheSize`** | int | Maximum size of LRU cache for routing optimization | `10000` |
| **`DisableCaching`** | bool | Disables LRU caching used to optimize routing performance | `false` |
| **`CacheMethods`** | string | Controls which HTTP methods are cached (comma-separated list) | `"GET,HEAD"` |
| **`MaxCachedParams`** | int | Maximum number of parameters to cache per route. Routes with more parameters than this will not be cached | `10` |
| **`MaxCachedPathLength`** | int | Maximum length of path to cache. Paths longer than this will not be cached | `100` |
| **`ReadBufferSize`** | int | Per-connection buffer size for request reading. This also limits the maximum header size. Increase this buffer for clients sending multi-KB RequestURIs and/or multi-KB headers (e.g., large cookies) | `8192 * 2` |
| **`HandleMethodNotAllowed`** | bool | Enables HTTP 405 Method Not Allowed responses when a route exists but the requested method is not supported, otherwise returns 404 | `false` |
| **`HandleOPTIONS`** | bool | Enables automatic replies to OPTIONS requests when no handlers are explicitly registered for that route | `false` |
| **`AutoRecover`** | bool | Enables automatic recovery from panics during handler execution by responding with HTTP 500 and logging the error without stopping the service | `true` |
| **`MaxRequestBodySize`** | int | Maximum request body size | `4 * 1024 * 1024` |
| **`MaxRouteParams`** | int | Maximum route parameter count | `1024` |
| **`MaxRequestURLLength`** | int | Maximum request URL length | `2048` |
| **`Concurrency`** | int | Maximum number of concurrent connections | `512 * 1024` |
| **`Prefork`** | bool | When enabled, spawns multiple Go processes listening on the same port | `false` |
| **`DisableStartupMessage`** | bool | Suppresses the Gonoleks startup message in console output | `false` |
| **`DisableKeepalive`** | bool | Disables keep-alive connections, causing the server to close connections after sending the first response to client | `false` |
| **`DisableDefaultDate`** | bool | When enabled, excludes the default Date header from responses | `false` |
| **`DisableDefaultContentType`** | bool | When enabled, excludes the default Content-Type header from responses | `false` |
| **`DisableHeaderNamesNormalizing`** | bool | When enabled, prevents header name normalization (e.g., conteNT-tYPE -> Content-Type) | true |
| **`ReadTimeout`** | `time.Duration` | Maximum time allowed to read the full request including body | `0` (unlimited) |
| **`WriteTimeout`** | `time.Duration` | Maximum duration before timing out writes of the response | `0` (unlimited) |
| **`IdleTimeout`** | `time.Duration` | Maximum time to wait for the next request when keep-alive is enabled | `0` (unlimited) |
| **`TLSEnabled`** | bool | Enables TLS (HTTPS) support | `false` |
| **`TLSCertPath`** | string | File path to the TLS certificate | `""` |
| **`TLSKeyPath`** | string | File path to the TLS key | `""` |
| **`DisableLogging`** | bool | Disables HTTP transaction logging | `false` |
| **`LogTimeFormat`** | string | Format string for log timestamps | `"2006/01/02 15:04:05"` |
| **`LogPrefix`** | string | Prefix for log messages | `""` |
| **`LogReportCaller`** | bool | Controls whether the caller information is included in logs | `false` |
| **`LogReportHost`** | bool | Controls whether the host information is included in logs | `false` |
| **`LogReportIP`** | bool | Controls whether client IP addresses are included in logs | `false` |
| **`LogReportUserAgent`** | bool | Controls whether the user agent information is included in logs | `false` |
| **`LogReportRequestHeaders`** | bool | Controls whether request headers are included in logs | `false` |
| **`LogReportRequestBody`** | bool | Controls whether request bodies are included in logs | `false` |
| **`LogReportResponseBody`** | bool | Controls whether response bodies are included in logs | `false` |
| **`LogReportResponseError`** | bool | Controls whether error responses are included in logs | `false` |

<!-- README links -->

[gonoleks_url]: https://github.com/gonoleks/gonoleks