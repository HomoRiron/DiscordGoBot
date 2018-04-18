package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

var client *discordgo.Session;

func login(token string) (*discordgo.Session) {
	discord,err := discordgo.New("Bot "+token)
	if(err != nil){
		log.Fatal(err)
	}
	return discord
}

func main(){
	client := login("Token Here")
	client.AddHandler(messageCreate)
	err := client.Open()
	errr(err)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
}
func errr(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
	}
}

func messageCreate(session *discordgo.Session,msg *discordgo.MessageCreate){
	var err error
	defer errr(err)
	if msg.Author.ID == session.State.SessionID {
		return
	}
	switch(msg.Content){
	case "Ping":
		_,err = session.ChannelMessageSend(msg.ChannelID,"Pong")
		break
	case "Embed":
		//Embed author
		var author = &discordgo.MessageEmbedAuthor{}
		author.IconURL = ""
		author.Name = ""
		author.ProxyIconURL = ""
		author.URL = ""

		var embed = &discordgo.MessageEmbed{}
		embed.URL = "https://example.com"
		embed.Title = "Embed Test"
		embed.Author = author
		embed.Description = "description"
		_,err = session.ChannelMessageSendEmbed(msg.ChannelID,embed)
		break
	}
}