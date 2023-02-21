package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"github.com/xlzd/gotp"
)

const version = "v0.3.0"

func main() {
	app := &cli.App{
		Name:  "otp-cli",
		Usage: "generate otp client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "secret",
				Aliases: []string{"s"},
				Usage:   "otp secret text",
				EnvVars: []string{"OTP_SECRET"},
			},
			&cli.StringFlag{
				Name:    "secret-file",
				Aliases: []string{"f"},
				Usage:   "otp secret file",
				EnvVars: []string{"OTP_SECRET_FILE"},
			},
			&cli.StringFlag{
				Name:    "scope",
				Usage:   "otp scope",
				EnvVars: []string{"OTP_SCOPE"},
			},
			&cli.BoolFlag{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "copy to clipboard",
				EnvVars: []string{"OTP_COPY"},
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "not output to console",
				EnvVars: []string{"OTP_QUIET"},
			},
		},
		Action: func(c *cli.Context) error {
			secret := c.String("secret")
			secretFile := c.String("secret-file")
			isCopy := c.Bool("copy")
			isQuiet := c.Bool("quiet")
			scope := c.String("scope")

			if scope != "" {
				bs, _ := ioutil.ReadFile(getScopeSecretFile(scope))
				if len(bs) != 0 {
					secret = string(bs)
				}
			}

			if secret == "" && secretFile == "" {
				return cli.ShowAppHelp(c)
			}

			if secretFile != "" {
				bs, err := os.ReadFile(secretFile)
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
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "show otp-cli version",
				Action: func(c *cli.Context) error {
					fmt.Println("otp-cli version " + version)
					return nil
				},
			},
			{
				Name:  "set-scope",
				Usage: "set scope secret",
				// otp-cli set-scope --name github --secret xxx
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Usage: "scope name", Required: true},
					&cli.StringFlag{Name: "secret", Usage: "scope secret", Required: true},
				},
				Action: func(c *cli.Context) error {
					name := c.String("name")
					secret := c.String("secret")
					if name == "" || secret == "" {
						return cli.ShowCommandHelp(c, "set-scope")
					}

					secretFile := getScopeSecretFile(name)
					if err := os.MkdirAll(filepath.Dir(secretFile), 0755); err != nil {
						return err
					}
					if err := os.WriteFile(secretFile, []byte(secret), 0644); err != nil {
						return err
					}
					log.Printf("set scope %s success", name)
					return nil
				},
			},
			{
				Name:  "del-scope",
				Usage: "delete scope secret",
				// otp-cli del-scope --name github
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Usage: "scope name", Required: true},
				},
				Action: func(c *cli.Context) error {
					name := c.String("name")
					if name == "" {
						return cli.ShowCommandHelp(c, "del-scope")
					}

					secretFile := getScopeSecretFile(name)
					if err := os.Remove(secretFile); err != nil {
						return err
					}
					log.Printf("delete scope %s success", name)
					return nil
				},
			},
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

func getWorkDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home + "/.local/otp-cli"
}

func getScopeSecretFile(scope string) string {
	return fmt.Sprintf("%s/.otp-cli/%s.secret.txt", getWorkDir(), scope)
}
