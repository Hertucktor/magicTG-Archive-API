package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main(){
	var language string


	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:        "lang",
				Value:       "en",
				Usage:       "language for the card name",
				Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "Vincent"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "de"  {
				//execute import with german foreign name filter
				fmt.Println("Moin", name)
			} else {
				//execute import with english name filter
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err)
	}
}