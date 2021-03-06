package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/soulshake/conduit/manager"
	"github.com/soulshake/conduit/version"
)

func run(c *cli.Context) {

	if len(c.StringSlice("repo")) == 0 {
		cli.ShowAppHelp(c)
		log.Fatal("you must specify at least 1 repo")
	}

	tags := c.StringSlice("tag")
	// add default tag if not present
	if len(tags) == 0 {
		tags = []string{"latest"}
	}
    //log.Infof("Using tags %s", tags)

	managerConfig := &manager.ManagerConfig{
		ListenAddr:    c.String("listen"),
		RepoWhitelist: c.StringSlice("repo"),
		Tags:          tags,
		DockerURL:     c.String("docker"),
		TLSCACert:     c.String("tls-ca-cert"),
		TLSCert:       c.String("tls-cert"),
		TLSKey:        c.String("tls-key"),
		AllowInsecure: c.Bool("allow-insecure"),
		AuthUsername:  c.String("auth-username"),
		AuthPassword:  c.String("auth-password"),
		AuthEmail:     c.String("auth-email"),
		Token:         c.String("token"),
		Debug:         c.Bool("debug"),
	}

	m, err := manager.NewManager(managerConfig)
	if err != nil {
		log.Fatal(err)
	}

	m.Run()
}

func main() {
	app := cli.NewApp()

	app.Name = "conduit"
	app.Author = "@ehazlett"
	app.Email = ""
	app.Usage = "docker auto-deployment system"
	app.Version = version.Version
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "Address to listen",
			Value: ":8080",
		},
		cli.StringFlag{
			Name:   "docker, d",
			Usage:  "URL to Docker",
			Value:  "unix:///var/run/docker.sock",
			EnvVar: "DOCKER_HOST",
		},
		cli.StringFlag{
			Name:  "tls-ca-cert",
			Usage: "TLS CA Certificate",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tls-cert",
			Usage: "TLS Certificate",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tls-key",
			Usage: "TLS Key",
			Value: "",
		},
		cli.BoolFlag{
			Name:  "allow-insecure",
			Usage: "Allow insecure communication to daemon",
		},
		cli.StringSliceFlag{
			Name:  "repo, r",
			Usage: "repo for whitelist",
			Value: &cli.StringSlice{},
		},
		cli.StringSliceFlag{
			Name:  "tag",
			Usage: "list of container tags for the repo to deploy",
			Value: &cli.StringSlice{},
		},
		cli.StringFlag{
			Name:   "auth-username, u",
			Usage:  "docker auth username (optional)",
			Value:  "",
			EnvVar: "DOCKER_AUTH_USERNAME",
		},
		cli.StringFlag{
			Name:   "auth-password, p",
			Usage:  "docker auth password (optional)",
			Value:  "",
			EnvVar: "DOCKER_AUTH_PASSWORD",
		},
		cli.StringFlag{
			Name:   "auth-email, e",
			Usage:  "docker auth email (optional)",
			Value:  "",
			EnvVar: "DOCKER_AUTH_EMAIL",
		},
		cli.StringFlag{
			Name:   "token, t",
			Usage:  "webhook token",
			Value:  "",
			EnvVar: "TOKEN",
		},
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
	}

	app.Run(os.Args)

}
