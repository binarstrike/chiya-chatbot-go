package command

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	Commands["ping"] = pingCommand
}

var pingCommand = CommandStruct{
	Ac: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "bot will reply pong!",
	},
	Fn: func(s *discordgo.Session, i *discordgo.InteractionCreate, opts CommandOptions) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
		if err != nil {
			log.Println(err)
		}
	},
}
