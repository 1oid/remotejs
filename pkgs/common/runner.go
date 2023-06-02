package common

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var cliApp *cli.App

func NewRunner(action cli.ActionFunc) {
	cliApp.Action = action
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cliApp = &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "url",
				Aliases:     []string{"u"},
				Usage:       "open url when open chrome, default blank url",
				Destination: &Vars.Url,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "chrome-path",
				Aliases:     []string{"cp"},
				Usage:       "use specified chrome path",
				Destination: &Vars.ChromePath,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "proxy",
				Usage:       "set proxy for browser",
				Destination: &Vars.Proxy,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "remote-debug-address",
				Usage:       "use remote chrome debugging",
				Destination: &Vars.RemoteDebuggingAddr,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "web-listen",
				Usage:       "web server port",
				Destination: &Vars.WebListenPort,
				Value:       "8088",
			},
		},
	}
}
