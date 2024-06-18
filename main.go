package main

import (
	
	"log"
)

func main() {

	svc := NewCollegeDetailService("https://universities.hipolabs.com/search?country=India")
	svc = NewLoggingService(svc)

	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":3000"))


}