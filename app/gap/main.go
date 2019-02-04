package main

import (
	"log"

	"github.com/elanq/gap"
)

func main() {
	resizer := &gap.Resizer{}
	app1 := gap.NewApplication("Notion")
	app2 := gap.NewApplication("ITerm2")

	log.Println(resizer.Left(app1))
	log.Println(resizer.Right(app2))
}
