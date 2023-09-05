package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

const USAGE string = "Search For IPs And Server Names on The Internet"

// Returns the command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = USAGE

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: USAGE,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "instagram.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
