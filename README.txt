To run this bot, first off all install golang, the programming language

then do
    go get github.com/shadowjonathan/OneDialog`
and
    go get github.com/shadowjonathan/beamer

these will install the required libraries

then grab your user-token (ask people how to do that), place it in a file called Token (just that, no extension) in %GOPATH%/src/github.com/shadowjonathan/beamer/ and do

    cd %GOPATH%/src/github.com/shadowjonathan/beamer/ && go run Beam.go

in the command line, that should start the bot

to test, do ">amion" in any channel, and if it responds, its on 

then you can make OS dialogue boxes with

>tb <face_filename> <text (just type it, make new lines with "\n")>