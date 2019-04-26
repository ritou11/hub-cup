package main

import (
  "os"
  "path"
  "bufio"
  "fmt"
  "strings"
  "errors"
  cli "gopkg.in/urfave/cli.v1"
  "github.com/juju/loggo"

  "hub-cup/libhub"
)

var logger = loggo.GetLogger("")

func main() {
  homedir, _ := os.UserHomeDir()
  app := &cli.App{
		Name:      "hub-sync",
		UsageText: `hub-sync <what> [<from>]`,
		Usage:     "Make your github forks catch up with origins",
		Version:   "0.1.0",
		HideHelp:  true,
		Flags: []cli.Flag{
      &cli.StringFlag{
        Name: "token, t",
        Usage: "Github token, see https://github.com/settings/tokens",
      },
			&cli.StringFlag{
        Name: "token-file",
        Usage: "`path` to your Github token file",
        Value: path.Join(homedir, ".hub-cup"),
      },
      &cli.BoolFlag{Name: "force, f", Usage: "As if {git push --force}"},
      &cli.BoolFlag{Name: "dry-run, n", Usage: "Don't actually update"},
			&cli.BoolFlag{Name: "debug", Usage: "print debug messages"},
			&cli.BoolFlag{Name: "help, h", Usage: "print the help"},
		},
		Action:  func (c *cli.Context) error {
      if c.Bool("help") || len(c.Args()) == 0 {
        cli.ShowAppHelpAndExit(c, 0)
      }
      if c.Bool("debug") {
    		loggo.ConfigureLoggers("<root>=DEBUG;libhub=DEBUG")
    	} else {
    		loggo.ConfigureLoggers("<root>=INFO;libhub=INFO")
    	}
      what := c.Args()[0]
      from := ""
      if len(c.Args()) > 1 {
        from = c.Args()[1]
      }
      logger.Debugf("what: %s; from: %s;", what, from)
      token := c.String("token")
      token_file := c.String("token-file")
      if len(token) == 0 {
        if _, err := os.Stat(token_file); os.IsNotExist(err) {
          fmt.Println(err)
          fmt.Println("Token needed!")
          return errors.New("no-token")
        }
        f, err := os.Open(token_file)
        defer f.Close()
        if err != nil {
          return err
        }
        reader := bufio.NewReader(f)
        token, err = reader.ReadString('\n')
        if err != nil {
          return err
        }
        token = strings.TrimSpace(token)
      }
      if len(token) == 0 {
        fmt.Println("Token file error!")
        return errors.New("token-file-error")
      }
      logger.Debugf("Token: %s", token)
      hc := libhub.New(token)
      err := hc.Cup(what, from, c.Bool("force"), c.Bool("dry-run"))
      if err != nil {
        logger.Errorf(err.Error())
        return err
      }
      return nil
    },
		Authors: []cli.Author{{Name: "Nogeek", Email: "ritou11@gmail.com"}},
	}

	app.Run(os.Args)
}
