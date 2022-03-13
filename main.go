package main

import "log"

func main() {
	srv := new(Server)
	if err := srv.Run(); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
