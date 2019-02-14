package main

import (
	"log"

	"github.com/elanq/gap"
)

func main() {
	s, err := gap.GetScreenSize()
	if err != nil {
		log.Println(err)
	}

	app1 := gap.NewApplication("Notion").Right(s)
	app2 := gap.NewApplication("Code").Left(s)

	gap.Resize(app1)
	gap.Resize(app2)
}
