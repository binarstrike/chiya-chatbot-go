package command

import (
	"fmt"
	"log"

	dcgo "github.com/bwmarrin/discordgo"
)

func init() {
	Commands["echo"] = EchoCommand
}

var EchoCommand CommandStruct = CommandStruct{
	Ac: &dcgo.ApplicationCommand{
		Name:        "echo",
		Description: "echoing what you type",
		Options: []*dcgo.ApplicationCommandOption{
			{
				Name:        "text",
				Description: "text to be echoing",
				Type:        dcgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},

	Fn: func(s *dcgo.Session, i *dcgo.InteractionCreate, opts CommandOptions) {
		var contentString string

		if u := i.Member.User; u != nil {
			contentString = fmt.Sprintf("<@%s> says: ", u.ID)
		}

		contentString += opts["text"].StringValue()

		err := s.InteractionRespond(i.Interaction, &dcgo.InteractionResponse{
			Type: dcgo.InteractionResponseChannelMessageWithSource,
			Data: &dcgo.InteractionResponseData{
				Content: contentString,
			},
		})
		if err != nil {
			log.Println(err)
		}
	},
}
