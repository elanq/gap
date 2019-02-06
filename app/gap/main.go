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

	log.Println(s.Height())
	log.Println(s.Width())
	//	resizer := &gap.Resizer{}
	//	app1 := gap.NewApplication("Notion")
	//	app2 := gap.NewApplication("ITerm2")
	//
	//	log.Println(resizer.Left(app1))
	//	log.Println(resizer.Right(app2))
}
