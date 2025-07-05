package main

import "github.com/BeGeos/go-echo/cmd/core"

func main() {
	// run setup based on environment
	e := core.App.Setup()

	e.Logger.Fatal(e.Start(":1323"))
}
