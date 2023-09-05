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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "instagram.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  USAGE,
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search For Server Names on The Internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host) //name server
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
