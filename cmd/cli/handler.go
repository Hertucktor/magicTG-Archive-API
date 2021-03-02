package main

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func ReadFromCLI() (string, string,string, error) {
	var inputCardNumber string
	var inputSetName string
	var language string

	app := cli.NewApp()
	reader := bufio.NewReader(os.Stdin)

	app.Flags = []cli.Flag {
		&cli.StringFlag{
			Name: "lang",
			Value: "en",
			Usage: "language for cli",
			Required: true,
		},
	}

	app.Action = func(c *cli.Context) error {

		if c.String("lang") == "de" {
			language = "de"
			fmt.Print("Trage die Kartennummer ein: ")
			inputCardNumber, _ = reader.ReadString('\n')
			fmt.Print("Trage den Set Namen ein: ")
			inputSetName, _ = reader.ReadString('\n')
		} else {
			language = "en"
			fmt.Print("Enter the card number: ")
			inputCardNumber, _ = reader.ReadString('\n')
			fmt.Print("Enter the set name: ")
			inputSetName, _ = reader.ReadString('\n')
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error().Err(err)
		return "", "", "", err
	}

	return language, inputCardNumber, inputSetName, err
}