package main

import (
	"log"

	"github.com/dev-hana/go-mailer/services"
)

func main() {
	h, err := services.NewHandler()
	if err != nil {
		log.Println(err)
		return
	}

	err = h.InitTable()
	if err != nil {
		log.Println(err)
		return
	}

}
