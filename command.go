package gommander

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type (
	Exec func(Ctx)

	Command struct {
		name        string
		description string
		exec        Exec
	}

	Commands map[string]Command

	CommandHandler struct {
		prefix string
		commands Commands
	}
)

/* Make a new command handler. */
func New(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix: prefix,
		commands: make(Commands),
	}
}

/* Get all the commands.
   - Returns a map.
*/
func (h *CommandHandler) GetAll() Commands {
	return h.commands
}

func (h *CommandHandler) Get(name string) (*Exec, bool) {
	cmd, found := h.commands[name]
	return &cmd.exec, found
}

/*
Register a command
*/
func (h *CommandHandler) RegisterOne(name string,description string, exec Exec) {
	cmd := Command{
		name:        name,
		description: description,
		exec:        exec,
	}
	h.commands[name] = cmd

	if len(name) > 1 {
		h.commands[name[:1]] = cmd
	}

	log.Printf("[RegisterOne] Command %s has been registered", name)
}

func (h *CommandHandler) DefaultHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author
	prefix := h.prefix

	if user.Bot {
		return
	}

	content := m.Content
	if len(content) <= len(prefix) {
		return
	}
	if content[:len(prefix)] != prefix {
		return
	}
	content = content[len(prefix):]
	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)
	trigger := strings.ToLower(args[0])
	exec, found := h.Get(trigger)

	if !found {
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Fatalf("Error while getting channel: %v", err)
		return
	}

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		log.Fatalf("Error while getting guild: %v", err)
		return
	}

	ctx := NewCtx(s, guild, channel, user, m.Message, h)
	ctx.Args = args[1:]
	cmd := *exec
	cmd(*ctx)	
}