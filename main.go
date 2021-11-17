package main

import (
	"log"

	"github.com/RobertoSuarez/monitor-mikrotik/api/device"
	"github.com/RobertoSuarez/monitor-mikrotik/api/procesos"
	"github.com/RobertoSuarez/monitor-mikrotik/config/application"
	"github.com/RobertoSuarez/monitor-mikrotik/monitor"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// for _, v := range monitor.Devs {
	// 	pro := monitor.NewProceso(*v)
	// 	go pro.ListenMikrotik(time.Minute * 2)
	// 	go pro.Run()
	// 	monitor.Pro[v.ID] = pro
	// }

	application := application.New(&application.Config{
		AppFiber: app,
	})

	application.Use(procesos.New(&procesos.Config{}))
	application.Use(device.New(&device.Config{
		Dev: monitor.Devs,
	}))

	// var id int
	// fmt.Scan(&id)

	// procesos[id].EnableReply()
	// contador := Count(0)
	// for v := range procesos[id].Chan() {
	// 	fmt.Println(v)

	// 	if contador() == 2 {
	// 		procesos[id].DisableReply()
	// 	}
	// }

	// fmt.Println("sali del for")

	log.Println(application.AppFiber.Listen(":8080"))
}

func Count(init int) func() int {
	i := init
	return func() int {
		i += 1
		return i
	}
}
