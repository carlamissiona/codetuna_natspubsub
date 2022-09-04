package main

import (
	"net/http"
    "codetunapubsub/nats"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
 
func main() {
	// Echo instance
	e := echo.New() 
 
    nats.Serve()
    
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
     
        nats.Pubs( []byte("{ 'data' : 'server has published  to nats '}}") )
		return c.String(http.StatusOK, "Each request to this route will be forwarded to Nats server and processed by the Nats subscriber!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
} 

