package main

import (
	"fmt"

	"github.com/ThatsOuttaPocket/data"
	"github.com/ThatsOuttaPocket/helpers"
)

func main() {
	fmt.Println("From main")
	helpers.Say_hello()
	fmt.Println(data.Filler_words)
}
