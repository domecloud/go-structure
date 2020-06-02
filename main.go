package main

import "dome.cloud/secureapi/routes"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
