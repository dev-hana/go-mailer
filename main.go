package main

import (
	"log"

	"github.com/dev-hana/go-mailer/routers"
)

func main() {
	log.Fatal(routers.RunAPI())
}
