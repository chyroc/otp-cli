package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"github.com/xlzd/gotp"
)

func main() {
	app := &cli.App{
		Name:  "otp-cli",
		Usage: "generate otp client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "secret",
				Aliases: []string{"s"},
				Usage:   "otp secret text",
			},
			&cli.StringFlag{
				Name:    "secret-file",
				Aliases: []string{"f"},
				Usage:   "otp secret file",
			},
			&cli.BoolFlag{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "copy to clipboard",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "not output to console",
			},
		},
		Action: func(c *cli.Context) error {
			secret := c.String("secret")
			secretFile := c.String("secret-file")
			isCopy := c.Bool("copy")
			isQuiet := c.Bool("quiet")

			if secret == "" && secretFile == "" {
				return cli.ShowAppHelp(c)
			}

			if secretFile != "" {
				bs, err := ioutil.ReadFile(secretFile)
				if err != nil {
					return err
				}
				secret = string(bs)
			}
			if secret == "" {
				return fmt.Errorf("secret is empty")
			}

			res, err := generate(secret)
			if err != nil {
				return err
			}

			if isCopy {
				if err = clipboard.WriteAll(res); err != nil {
					return err
				}
			}

			if !isQuiet {
				fmt.Println(res)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func generate(secret string) (s string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	return gotp.NewDefaultTOTP(strings.ToUpper(strings.Replace(secret, " ", "", -1))).Now(), nil
}
