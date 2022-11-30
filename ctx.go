package gommander

import (
	"github.com/bwmarrin/discordgo"
)

type Ctx struct {
	Session *discordgo.Session 
	Guild *discordgo.Guild
	Channel *discordgo.Channel 
	User *discordgo.User 
	Message *discordgo.Message
	Args []string

	Handler *CommandHandler
}

func NewCtx(session *discordgo.Session, guild *discordgo.Guild, channel *discordgo.Channel, user *discordgo.User, message *discordgo.Message, handler *CommandHandler) *Ctx {
	ctx := new(Ctx)
	ctx.Session = session 
	ctx.Guild = guild 
	ctx.Channel = channel 
	ctx.User = user 
	ctx.Message = message

	ctx.Handler = handler

	return ctx
}

func (ctx *Ctx) Send(content string) {
	ctx.Session.ChannelMessageSend(ctx.Channel.ID, content)
}