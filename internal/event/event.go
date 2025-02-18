package event

import (
	"github.com/binarstrike/chiya-chatbot-go/internal/command"
	dcgo "github.com/bwmarrin/discordgo"
)

type Event struct {
	session *dcgo.Session
}

type EventStruct struct {
	Once bool
	Fn   any
}

type EventMap = map[string]*EventStruct

var eventMap = make(EventMap)

func NewEvent(s *dcgo.Session) *Event {
	return &Event{s}
}

func (ev *Event) LoadAllEvents() {
	for _, es := range eventMap {
		if es.Once {
			ev.session.AddHandlerOnce(es.Fn)
			continue
		}

		ev.session.AddHandler(es.Fn)
	}
}

func convertToCommandOptions(options []*dcgo.ApplicationCommandInteractionDataOption) command.CommandOptions {
	optionMap := make(command.CommandOptions)
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	return optionMap
}
