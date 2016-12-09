package main

import (
	"log"

	"github.com/bketelsen/micro-employees/db"
	"github.com/bketelsen/micro-employees/handler"
	proto "github.com/bketelsen/micro-employees/proto/employees"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("gophertrain.employees"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g docker@tcp(127.0.0.1:3306)/user",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()
	db.Init()

	proto.RegisterEmployeesHandler(service.Server(), new(handler.Employee))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
