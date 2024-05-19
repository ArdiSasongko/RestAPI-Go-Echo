package main

func main() {
	server := StartServer()
	err := server.Start(":8080")
	if err != nil {
		return
	}
}
