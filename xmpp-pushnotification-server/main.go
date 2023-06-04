package main

import (
	"fmt"

	"xmpp-pushnotification-server/app"
	"xmpp-pushnotification-server/app/httpserve"
)

func main() {
	fmt.Println("Staring code")
	app.Start()

	httpserve.StartHttpServer()
}
