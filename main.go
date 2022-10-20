package main

import (
	"fmt"

	"github.com/Aberry-byte/ThatsOuttaPocket/data"
	"github.com/Aberry-byte/ThatsOuttaPocket/helpers"
)

func main() {
	fmt.Println("From main")
	helpers.Say_hello()
	fmt.Println(data.Filler_words)
}
