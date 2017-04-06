package main

import (
	"bufio"
	fmt "fmt"
	"log"

	"io/ioutil"

	"strings"

	"bytes"

	"image/png"

	"time"

	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/shadowjonathan/onedialog"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("FATAL ERROR:\nGOPATH IS NOT SET, PLEASE SET GOPATH TO A VALUE, OR ELSE IMAGE GENERATION WONT WORK!")
	}
	var data []byte
	var err error
	data, err = ioutil.ReadFile("Token")
	if err != nil {
		data, _ = ioutil.ReadFile("Token.txt")
		if len(data) < 1 {
			fmt.Println("Error while getting token file, it probably doesnt exist:", err)
			return
		}
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
var defaultchannel string

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
	fmt.Println("Ready to receive input!\nInput dialog data in this window just like discord, except for the \">tb\" part\nA default channel needs to be set, set one with '>select' in any channel, and command-line calls will select that channel as output.")
	go func() {
		for {
			result := GetInput()
			if defaultchannel == "" && !strings.Contains(result, "select") {
				fmt.Println("Error, default channel not set, do so with '>select' inside one.")
				continue
			}
			ts := strings.Split(result, " ")
			if len(ts) < 2 {
				fmt.Println("Unsufficient arguments.")
			} else {
				face := ts[0]
				text := strings.Join(ts[1:], " ")
				if face == "select" {
					defaultchannel = text
					fmt.Println("Manually selected channel")
					continue
				} else if defaultchannel == "" {
					fmt.Println("Error, default channel not set, do so with '>select' inside one.")
					continue
				}
				PostDialog(face, text, defaultchannel)
			}
		}
	}()
}

func Mess(Ses *discordgo.Session, MesC *discordgo.MessageCreate) {
	if MesC.Author.ID == bbb.dg.State.User.ID {
		if strings.Contains(MesC.Content, ">") {
			ts := strings.Split(MesC.Content, " ")
			if ts[0] == ">amion" {
				bbb.dg.ChannelMessageSend(MesC.ChannelID, "`Hello, master`")
				return
			} else if ts[0] == ">Q" {
				fmt.Println("Issued quit command, stopping...")
				os.Exit(0)
			} else if strings.ToLower(ts[0]) == ">select" {
				defaultchannel = MesC.ChannelID
				bbb.dg.ChannelMessageSend(MesC.ChannelID, "`Selected channel.`")
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
	img, succ := OD.Make(face, text)
	if !succ {
		fmt.Println("Error in draw operation")
		return
	}
	b := new(bytes.Buffer)
	err := png.Encode(b, img)
	if err != nil {
		panic(err)
	}
	_, err = bbb.dg.Channel(channel)
	if err != nil {
		potchan, err := bbb.dg.UserChannelCreate(channel)
		if err == nil {
			channel = potchan.ID
		} else {
			fmt.Println("Unknown error when trying to resolve channel/ID ", channel)
			return
		}
	}
	bbb.dg.ChannelTyping(channel)
	bbb.dg.ChannelFileSend(channel, "Dialog.png", b)
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	response = strings.TrimSpace(response)
	return response
}
