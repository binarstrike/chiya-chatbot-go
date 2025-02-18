package event

import (
	"github.com/binarstrike/chiya-chatbot-go/internal/command"
	"github.com/bwmarrin/discordgo"
)

func init() {
	eventMap["slash-command"] = slashCommandHandler
}

var slashCommandHandler = &EventStruct{
	Once: false,
	Fn: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()

		if cs, ok := command.Commands[data.Name]; ok {
			opts := convertToCommandOptions(data.Options)
			cs.Fn(s, i, opts)
		}
	},
}
