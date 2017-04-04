package main

import (
	fmt "fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {

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
}
