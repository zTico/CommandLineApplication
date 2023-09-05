package app

import "github.com/urfave/cli"

//Returns the command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search For IPs And Server Names on The Internet"

	return app
}
