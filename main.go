// Package main frontend server application
package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var port int
var folder string
var handlerName string

// main entrypoint
func main() {
	parseFlags()
	var handler http.Handler
	switch handlerName {
	case "mux":
		handler = getMuxHandler()
	case "gin":
		handler = getGinHandler()
	case "echo":
		handler = getEchoHandler()
	}
	log.Printf("starting frontend server application on %d port for %s folder", port, folder)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

// parseFlags get commandline params
func parseFlags() {
	flag.IntVar(&port, "port", 8000, "port serve frontend application")
	flag.StringVar(&folder, "folder", "./public", "folder frontend application")
	flag.StringVar(&handlerName, "handler", "mux", "handler name: mux, gin, echo")
	flag.Parse()
}

// getMuxHandler Mux Handler
func getMuxHandler() http.Handler {
	h := http.NewServeMux()
	h.Handle("/", http.FileServer(http.Dir(folder)))
	return h
}

// getGinHandler Gin Handler
func getGinHandler() http.Handler {
	h := gin.Default()
	h.Static("/", folder)
	return h
}

// getEchoHandler Echo handler
func getEchoHandler() http.Handler {
	h := echo.New()
	h.Static("/", folder)
	return h
}
