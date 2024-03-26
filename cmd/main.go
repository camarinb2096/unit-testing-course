package main

import (
	"camarinb2096/unit-testing-course/internal/config/server"
)

func main() {

	Server := server.NewServer()
	Server.Routes()
	Server.Start()

}
