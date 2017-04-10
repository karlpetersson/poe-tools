package main

import (
    "text/template"
    "bytes"
    "strings"
    "github.com/ccbrown/poe-go/api"
)

type message struct {
    Name string
    Type string
    To string
    Price string
}

func parseTemplate(msg message) string {
    templateStr := "@{{.To}} Hi, I would like to buy your {{.Name}} {{.Type}} in Legacy listed for {{.Price}}."
    template, err := template.New("Message").Parse(templateStr)
    if err != nil {
        panic(err)
    }
    var msgBytes bytes.Buffer
    err = template.Execute(&msgBytes, msg)
    if err != nil {
        panic(err)
    }
    return msgBytes.String()
}

func newMessage(item api.Item, stash *api.Stash) string {
    msg := message {
        Name: strings.Trim(item.Name, "<<set:MS>><<set:M>><<set:S>>"),
        Type: item.Type,
        To: stash.LastCharacterName,
        Price: item.Note,
    }

    return parseTemplate(msg)
}