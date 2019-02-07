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

	resizer := &gap.Resizer{}
	app1 := gap.NewApplication("Notion").Left(s)
	app2 := gap.NewApplication("ITerm2").Right(s)

	log.Println(resizer.Do(app1))
	log.Println(resizer.Do(app2))
}
