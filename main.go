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
var routerName string

// main entrypoint
func main() {
	parseFlags()
	var router http.Handler
	switch routerName {
	case "mux":
		router = getMuxHandler()
	case "gin":
		router = getGinHandler()
	case "echo":
		router = getEchoHandler()
	}
	log.Printf("starting frontend server application on %d port for %s folder with %s router", port, folder, routerName)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// parseFlags get commandline params
func parseFlags() {
	flag.IntVar(&port, "port", 8000, "port serve frontend application")
	flag.StringVar(&folder, "folder", ".", "folder frontend application")
	flag.StringVar(&routerName, "router", "mux", "router name: mux, gin, echo")
	flag.Parse()
}

// getMuxHandler Mux router
func getMuxHandler() http.Handler {
	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.Dir(folder)))
	return r
}

// getGinHandler Gin router
func getGinHandler() http.Handler {
	r := gin.Default()
	r.Static("/", folder)
	return r
}

// getEchoHandler Echo router
func getEchoHandler() http.Handler {
	r := echo.New()
	r.Static("/", folder)
	return r
}
