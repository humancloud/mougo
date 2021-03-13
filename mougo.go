package main

import (
	"log"
	"mougo/model"
	"mougo/routers"
	"mougo/utils"
)

func main() {
	// ReadConfig and set settings
	utils.InitConfig()

	// Init Router
	mougo := routers.InitRouter()

	// Init Db
	model.InitDb()

	// Run...
	if err := mougo.Run(utils.HttpPort); err != nil {
		log.Fatalln(err.Error())
	}
}
