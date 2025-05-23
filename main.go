package main

import "School_Management_System/app"

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}
