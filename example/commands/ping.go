package commands

import "github.com/adjective-john/gommander"

func PingCommand(ctx gommander.Ctx) {
	ctx.Send("🏓 Pong!")
}