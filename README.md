# Command handler for DiscordGo (For Go 1.18)
Gommander is a legacy command handler for [DiscordGo](https://github.com/bwmarrin/discordgo)

# Installation
```sh
go get github.com/adjective-john/gommander 
```

Or 

```go 
import "github.com/adjective-john/gommander"
```

# Getting Started
```go
handler := gommander.New("~")

/* 
`s` is your discordgo session.

- DefaultHandler is the default handler that will handle and execute commands.
*/
s.AddHandler(handler.DefaultHandler)

/*
Registering commands.

RegisterOne(...) - Register a command.
*/
handler.RegisterOne("ping", "Says pong!", PingCommand)

/*
Simple ping command.
*/
func PingCommand(ctx gommander.Ctx) {
    ctx.Send("Pong!")
}
```

Check [example](https://github.com/ApastaDev/gommander/tree/main/example) directory.