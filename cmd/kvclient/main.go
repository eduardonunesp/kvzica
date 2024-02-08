package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eduardonunesp/kvzika/pkg/client"
	"github.com/urfave/cli/v2"
)

func main() {
	client, err := client.NewClient("0.0.0.0", 5566)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name: "kvclient",
		Commands: []*cli.Command{
			{
				Name:  "set",
				Usage: "Set a key value pair",
				Action: func(c *cli.Context) error {
					key := c.Args().Get(0)
					value := c.Args().Get(1)
					if err := client.SetKeyValue([]byte(key), []byte(value)); err != nil {
						return err
					}
					fmt.Println("Key set successfully", key, value)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "Get a value by key",
				Action: func(c *cli.Context) error {
					key := c.Args().Get(0)
					value, err := client.GetValue([]byte(key))
					if err != nil {
						return err
					}
					log.Println(string(value))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
