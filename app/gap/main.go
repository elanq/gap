package main

import (
	"log"

	"github.com/elanq/gap"
)

func main() {
	screenSize, err := gap.GetScreenSize()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(screenSize.Height())
	log.Println(screenSize.Width())

}
