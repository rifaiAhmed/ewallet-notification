package main

import (
	"ewallet-notification/cmd"
	"ewallet-notification/helpers"
)

func main() {

	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupMySQL()

	// run grpc
	cmd.ServeGRPC()

	// run http
	// cmd.ServeHTTP()
}
