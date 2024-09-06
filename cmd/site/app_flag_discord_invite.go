package main

import (
	"github.com/potato-squad-org/site/globals"
	"github.com/urfave/cli/v2"
)

func init() {
	app.Flags = append(app.Flags, discordInviteFlag)
}

var discordInviteFlag = &cli.StringFlag{
	Name:        "discord-invite",
	Usage:       "Discord invite URL",
	Aliases:     []string{"discord-link"},
	EnvVars:     []string{"DISCORD_INVITE_URL"},
	Required:    true,
	Destination: &globals.DiscordInviteURL,
}
