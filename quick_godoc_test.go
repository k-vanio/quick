// Package quick provides a high-performance, minimalistic web framework for Go.
//
// This file contains **unit tests** for various functionalities of the Quick framework,
// including route handling, middleware, static file serving, and request handling.
//
// These tests ensure that the core features of Quick work as expected.
//
// 📌 To run all unit tests, use:
//
//	$ go test -v ./...
//	$ go test -v
package quick

import (
	"github.com/jeffotoni/quick/middleware/cors"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TestExamplePath verifies that a PATCH route returns the expected response
// The test simulates a partial update functionality.
// The will test TestExamplePath(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExamplePath
func TestExamplePath(t *testing.T) {
	q := New()

	// Define uma rota PATCH para atualizar parcialmente um recurso
	q.Patch("/update-partial", func(c *Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.Status(200).String("Feature partially updated!")
	})

	// Simulates a PATCH request
	data, err := q.QuickTest("PATCH", "/update-partial", nil)
	if err != nil {
		t.Errorf("Error running QuickTest: %v", err)
		return
	}

	// Verifica se o código de status HTTP retornado está correto
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200 but got %d", data.StatusCode())
	}

	// Verifica se o corpo da resposta contém a mensagem esperada
	expectedBody := "Feature partially updated!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s' but got '%s'", expectedBody, data.BodyStr())
	}
}

// TestExampleOptions verifies that the OPTIONS route properly returns the allowed methods.
// The test simulates a partial update functionality.
// The will test TestExampleOptions(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExampleOptions
func TestExampleOptions(t *testing.T) {
	q := New()

	// Define a GET route
	q.Get("/example", func(c *Ctx) error {
		return c.Status(200).String("GET is working!")
	})

	// Register OPTIONS for the /example route
	q.Options("/example", func(c *Ctx) error {
		// Define the methods allowed for this resource in the Allow header
		allowedMethods := "GET, POST, PUT, DELETE, PATCH, OPTIONS"
		c.Set("Allow", allowedMethods)
		return c.Status(200).String("OPTIONS request handled")
	})

	// Simulate an OPTIONS request
	data, err := q.QuickTest("OPTIONS", "/example", nil)
	if err != nil {
		t.Errorf("Error executing QuickTest: %v", err)
		return
	}

	// Verify that the HTTP status code is 200
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", data.StatusCode())
	}

	// Get the "Allow" header from the HTTP response
	allowHeader := data.Response().Header.Get("Allow")
	expectedAllow := "GET, POST, PUT, DELETE, PATCH, OPTIONS"
	if allowHeader != expectedAllow {
		t.Errorf("Expected Allow header '%s', but got '%s'", expectedAllow, allowHeader)
	}
}

// TestExampleGetDefaultConfig verifies if GetDefaultConfig() returns the expected default configuration values.
// The will test TestExampleGetDefaultConfig(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExampleGetDefaultConfig
func TestExampleGetDefaultConfig(t *testing.T) {
	// Expected default configuration values
	expected := Config{
		BodyLimit:      2097152, // 2MB
		MaxBodySize:    2097152, // 2MB
		MaxHeaderBytes: 1048576, // 1MB
		RouteCapacity:  1000,    // Maximum number of routes
		MoreRequests:   290,     // Max concurrent requests allowed
	}

	// Get actual configuration
	result := GetDefaultConfig()

	// Verify if the configuration matches the expected values
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("GetDefaultConfig() did not return expected configuration. Expected %+v, got %+v", expected, result)
	}
}

// TestExampleNew verifies if a simple GET route returns the expected response.
// The will test TestExampleNew(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExampleNew
func TestExampleNew(t *testing.T) {
	q := New()

	// Define a simple GET route
	q.Get("/", func(c *Ctx) error {
		// Set response content type
		c.Set("Content-Type", "text/plain")
		// Return a success message
		return c.Status(200).String("Quick in action ❤️!")
	})

	// Simulate a GET request
	data, err := q.QuickTest("GET", "/", nil)
	if err != nil {
		t.Errorf("Error when running QuickTest: %v", err)
		return
	}

	// Validate HTTP status code
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", data.StatusCode())
	}

	// Validate response body
	expectedBody := "Quick in action ❤️!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, data.BodyStr())
	}
}

// TestExampleUse verifies if a middleware (CORS) is correctly applied to the route.
// The will test TestExampleUse(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExampleUse
func TestExampleUse(t *testing.T) {
	q := New()

	// Apply CORS middleware
	q.Use(cors.New())

	// Define a GET route that uses middleware
	q.Get("/use", func(c *Ctx) error {
		// Set response content type
		c.Set("Content-Type", "text/plain")
		// Return success message
		return c.Status(200).String("Quick in action com middleware ❤️!")
	})

	// Simulate a GET request
	data, err := q.QuickTest("GET", "/use", nil)
	if err != nil {
		t.Errorf("Error when running QuickTest: %v", err)
		return
	}

	// Validate HTTP status code
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", data.StatusCode())
	}

	// Validate response body
	expectedBody := "Quick in action com middleware ❤️!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, data.BodyStr())
	}
}

// TestExampleGet verifies if a GET request returns the expected response.
// The will test TestExampleGet(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExampleGet
func TestExampleGet(t *testing.T) {
	q := New()

	// Define a GET route
	q.Get("/hello", func(c *Ctx) error {
		// Set response content type
		c.Set("Content-Type", "text/plain")
		// Return success message
		return c.Status(200).String("Olá, mundo!")
	})

	// Simulate a GET request
	data, err := q.QuickTest("GET", "/hello", nil)
	if err != nil {
		t.Errorf("Error when running QuickTest: %v", err)
		return
	}

	// Validate HTTP status code
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200, but got %d", data.StatusCode())
	}

	// Validate response body
	expectedBody := "Olá, mundo!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, data.BodyStr())
	}
}

// TestExamplePost verifies if a POST request returns the expected response.
// The will test TestExamplePost(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExamplePost
func TestExamplePost(t *testing.T) {
	q := New()

	// Define a POST route
	q.Post("/create", func(c *Ctx) error {
		// Set response content type
		c.Set("Content-Type", "text/plain")
		// Return success message
		return c.Status(201).String("Recurso criado!")
	})

	// Simulate a POST request
	data, err := q.QuickTest("POST", "/create", nil)
	if err != nil {
		t.Errorf("Error when running QuickTest: %v", err)
		return
	}

	// Validate HTTP status code
	if data.StatusCode() != 201 {
		t.Errorf("Expected status 201, but got %d", data.StatusCode())
	}

	// Validate response body
	expectedBody := "Recurso criado!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, data.BodyStr())
	}
}

// TestExamplePut verifies if a PUT request updates the resource and returns the expected response.
// The will test TestExamplePut(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExamplePut
func TestExamplePut(t *testing.T) {
	q := New()

	// Define a PUT route for updating a resource
	q.Put("/update", func(c *Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.Status(200).String("Recurso atualizado!")
	})

	// Simulate a PUT request
	data, err := q.QuickTest("PUT", "/update", nil)
	if err != nil {
		t.Fatalf("QuickTest failed: %v", err)
	}

	// Check if the status code is correct
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200, but received %d", data.StatusCode())
	}

	// Check if the response body is correct
	expectedBody := "Recurso atualizado!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but received '%s'", expectedBody, data.BodyStr())
	}
}

// TestExampleDelete verifies if a DELETE request correctly deletes a resource and returns the expected response.
// The will test TestExampleDelete(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestExampleDelete
func TestExampleDelete(t *testing.T) {
	q := New()

	// Define a DELETE route for deleting a resource
	q.Delete("/delete", func(c *Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.Status(200).String("Recurso deletado!")
	})

	// Simulate a DELETE request
	data, err := q.QuickTest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatalf("QuickTest failed: %v", err)
	}

	// Check if the status code is correct
	if data.StatusCode() != 200 {
		t.Errorf("Expected status 200, but received %d", data.StatusCode())
	}

	// Check if the response body is correct
	expectedBody := "Recurso deletado!"
	if data.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but received '%s'", expectedBody, data.BodyStr())
	}
}

// TestServeHTTP verifies if dynamic route parameters are correctly handled in a GET request.
// The will test TestServeHTTP(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestServeHTTP
func TestServeHTTP(t *testing.T) {
	q := New()

	// Define a GET route with a dynamic parameter
	q.Get("/users/:id", func(c *Ctx) error {
		return c.Status(200).String("User Id: " + c.Params["id"])
	})

	// Simulate a GET request with a user ID
	res, err := q.QuickTest("GET", "/users/42", nil)
	if err != nil {
		t.Fatalf("QuickTest failed: %v", err)
	}

	// Check if the status code is correct
	expectedStatus := 200
	if res.StatusCode() != expectedStatus {
		t.Errorf("Expected status %d, but got %d", expectedStatus, res.StatusCode())
	}

	// Check if the response body is correct
	expectedBody := "User Id: 42"
	if res.BodyStr() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, res.BodyStr())
	}
}

// TestGetRoute verifies if the registered routes are correctly retrieved.
// The will test TestGetRoute(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestGetRoute
func TestGetRoute(t *testing.T) {
	q := New()

	// Define multiple routes
	q.Get("/users/:id", func(c *Ctx) error {
		return c.Status(200).String("User ID: " + c.Params["id"])
	})
	q.Post("/users", func(c *Ctx) error {
		return c.Status(201).String("User created")
	})

	// Retrieve the registered routes
	routes := q.GetRoute()

	// Check if the expected number of routes exists
	expectedNumRoutes := 2
	if len(routes) != expectedNumRoutes {
		t.Errorf("Expected %d routes, but got %d", expectedNumRoutes, len(routes))
	}

	// Define expected routes
	expectedRoutes := map[string]string{
		"GET":  "/users/:id",
		"POST": "/users",
	}

	// Check if the routes match the expected values
	for _, route := range routes {
		pattern := route.Pattern
		if pattern == "" {
			pattern = route.Path
		}

		expectedPattern, exists := expectedRoutes[route.Method]
		if !exists {
			t.Errorf("Unexpected HTTP method: %s", route.Method)
		} else if pattern != expectedPattern {
			t.Errorf("Expected pattern for %s: %s, but got %s", route.Method, expectedPattern, route.Pattern)
		}
	}
}

// TestQuick_ExampleListen verifies if the Quick server starts correctly and serves responses.
// The will test TestQuick_ExampleListen(t *testing.T)
//
// Run:
//
//	$ go test -v -run ^TestQuick_ExampleListen
func TestQuick_ExampleListen(t *testing.T) {

	// start Quick
	q := New()

	// Define a simple GET route
	q.Get("/", func(c *Ctx) error {
		return c.Status(200).String("Hello, Quick!")
	})

	// Start the server using ListenWithShutdown on a dynamic port
	server, shutdown, err := q.ListenWithShutdown(":0")
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer shutdown()

	// Wait for the server to start correctly
	time.Sleep(500 * time.Millisecond)

	// Make an HTTP request using the returned dynamic port
	resp, err := http.Get("http://" + server.Addr + "/")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer resp.Body.Close()

	// Checks if the returned status is correct
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", resp.StatusCode)
	}
}
