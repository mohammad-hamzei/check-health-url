package main

import (
	"fmt"
	"git.tashilcar.com/core/tashilcar/app"
	_ "git.tashilcar.com/core/tashilcar/docs"
	"git.tashilcar.com/core/tashilcar/domain/usecases"
	"github.com/jasonlvhit/gocron"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
)

func main() {
	//go executeCronJob()
	//	@title			Go + Gin Tashilcar API
	//	@version		1.0
	//	@description	This is a sample server tashilcar server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

	//	@contact.name	API Support
	//	@contact.url	http://www.swagger.io/support
	//	@contact.email	support@swagger.io

	//	@license.name	MIT
	//	@license.url	https://opensource.org/licenses/MIT

	//	@host						localhost:8080
	//	@BasePath					/
	//	@query.collection.format	multi

	fmt.Println("** ABOUT TO START APPLICATION **")
	app.StartApplication()

	//wait for ctrl+c to exit

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	//block until signal is received
	<-ch

	fmt.Println("** ABOUT TO STOP APPLICATION **")
	app.StopApplication()
}

func myTask() {
	usecases.CheckHealth()
}
func executeCronJob() {
	gocron.Every(1).Second().Do(myTask)
	<-gocron.Start()
}
