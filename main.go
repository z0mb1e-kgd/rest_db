package main

import "log"

func main() {
	handlers := new(Handler)
	srv := new(Server)
	if err := srv.Run(handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
