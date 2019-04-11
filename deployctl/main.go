package main

import (
	"os"

	"github.com/urfave/cli"
)

var version = "1.3.0"

func main() {

	app := cli.NewApp()
	app.Name = "deployctl"
	app.Usage = "SYROS deploy CLI"
	app.Author = "Stefan Prodan"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "Download URL for the config.tar.gz file",
			EnvVar: "DCTL_CONFIG_URL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "promote",
			Usage:  "Promote containers from one environment to another",
			Action: componentPromote,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ticket, t",
					Usage: "JIRA ticket ID, if specified the deploy log will be posted on the ticket",
				},
				cli.StringSliceFlag{
					Name:  "environment, e",
					Usage: "Target environment, multiple values accepted",
				},
				cli.StringSliceFlag{
					Name:  "component, c",
					Usage: "Docker service, multiple values accepted",
				},
				cli.StringFlag{
					Name:  "tag",
					Usage: "If a tag is specified this exact docker image tag will be deployed",
				},
			},
		},
		{
			Name:   "reload",
			Usage:  "Reload containers configuration",
			Action: componentReload,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ticket, t",
					Usage: "JIRA ticket ID, if specified the deploy log will be posted on the ticket",
				},
				cli.StringSliceFlag{
					Name:  "environment, e",
					Usage: "Target environment, multiple values accepted",
				},
				cli.StringSliceFlag{
					Name:  "component, c",
					Usage: "Docker service, multiple values accepted",
				},
			},
		},
		{
			Name:   "rollback",
			Usage:  "Rollback containers to previous version",
			Action: componentRollback,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ticket, t",
					Usage: "JIRA ticket ID, if specified the rollback log will be posted on the ticket",
				},
				cli.StringSliceFlag{
					Name:  "environment, e",
					Usage: "Target environment, multiple values accepted",
				},
				cli.StringSliceFlag{
					Name:  "component, c",
					Usage: "Docker service, multiple values accepted",
				},
			},
		},
		{
			Name:   "migrate",
			Usage:  "Migrate PostgreSQL, Kafka or OpenTSDB",
			Action: componentMigrate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ticket, t",
					Usage: "JIRA ticket ID, if specified the rollback log will be posted on the ticket",
				},
				cli.StringSliceFlag{
					Name:  "environment, e",
					Usage: "Target environment, multiple values accepted",
				},
				cli.StringSliceFlag{
					Name:  "component, c",
					Usage: "Component name",
				},
			},
		},
	}

	app.Run(os.Args)
}
