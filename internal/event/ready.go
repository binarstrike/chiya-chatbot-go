package event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	eventMap["ready"] = ready
}

var ready = &EventStruct{
	Once: false,
	Fn: func(s *discordgo.Session, _ *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	},
}
