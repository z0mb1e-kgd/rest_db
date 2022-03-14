package main

func main() {
	router := SetRouter()
	router.Run(serverPort)
}
