package main

import (
	"fmt"
	readtoml "goprograms/ReadToml"
	getapistatus "goprograms/getAPIstatus"
)

func main() {
	fmt.Println("Server Started (+)")
	getapistatus.APIServiceStatusMain()
	readtoml.GetTomlValues()
	fmt.Println("Server Ended (-)")
}
