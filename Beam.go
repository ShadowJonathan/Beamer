package main

import (
	fmt "fmt"

	"io/ioutil"

	"strings"

	"bytes"

	"image/png"

	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shadowjonathan/onedialog"
)

func main() {
	data, err := ioutil.ReadFile("Token")
	if err != nil {
		fmt.Println("Error while getting token file, it probably doesnt exist:", err)
		return
	}
	fmt.Println("LOADING BOT WITH", string(data))
	Initialize(string(data))
}

type Version struct {
	Major               byte
	Minor               byte
	Build               byte
	Experimental        bool
	ExperimentalVersion byte
}

type Bot struct {
	dg *discordgo.Session
}

// Vars after this

var bbb *Bot
var err error

// Functions after this

func Initialize(Token string) {
	bbb = &Bot{}

	bbb.dg, err = discordgo.New(Token)
	if err != nil {
		fmt.Println("Discord Session error, check token, error message: " + err.Error())
		return
	}

	bbb.dg.AddHandler(Ready)
	bbb.dg.AddHandler(Mess)

	bbb.dg.Open()
	fmt.Println("DG OPENED")
	for {
		time.Sleep(10 * time.Second)
	}
}

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Discord: Ready message received\nSH: I am '" + r.User.Username + "'!\nSH: My User ID: " + r.User.ID)
}

func Mess(Ses *discordgo.Session, MesC *discordgo.MessageCreate) {
	if MesC.Author.ID == bbb.dg.State.User.ID {
		if strings.Contains(MesC.Content, ">") {
			ts := strings.Split(MesC.Content, " ")
			if ts[0] == ">amion" {
				bbb.dg.ChannelMessageSend(MesC.ChannelID, "`Hello, master`")
				return
			}
			if len(ts) > 2 {
				if ts[0] == ">tb" {
					face := ts[1]
					text := strings.Join(ts[2:], " ")
					bbb.dg.ChannelMessageDelete(MesC.ChannelID, MesC.ID)
					PostDialog(face, text, MesC.ChannelID)
				}
			}
		}
	}
}

func PostDialog(face, text, channel string) {
	img := OD.Make(face, text)
	b := new(bytes.Buffer)
	err := png.Encode(b, img)
	if err != nil {
		panic(err)
	}
	bbb.dg.ChannelFileSend(channel, "Dialog.png", b)
}
