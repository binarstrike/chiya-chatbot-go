package command

import (
	"log"

	dcgo "github.com/bwmarrin/discordgo"
)

type CommandHandlerFunc = func(s *dcgo.Session, i *dcgo.InteractionCreate, opts CommandOptions)

type CommandStruct struct {
	Ac *dcgo.ApplicationCommand
	Fn CommandHandlerFunc
}

type CommandMap map[string]CommandStruct

type CommandOptions = map[string]*dcgo.ApplicationCommandInteractionDataOption

var Commands = make(CommandMap)

func DeploySlashCommand(s *dcgo.Session, appId string) {
	for name, command := range Commands {
		_, err := s.ApplicationCommandCreate(appId, "", command.Ac)
		if err != nil {
			log.Printf("Error deploying '%s' slash command, error: %v", name, err)
		}

		log.Printf("Command deployed: %s", name)
	}
}
