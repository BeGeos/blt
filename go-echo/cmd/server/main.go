package main

import (
	"github.com/BeGeos/go-echo/internal/app"
)

func main() {
	// run setup based on environment
	e := app.Setup()

	e.Logger.Fatal(e.Start(":1323"))
}
