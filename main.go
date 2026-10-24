package main

import (
	"context"
	"fmt"
	"log"

	_ "charm.land/bubbles/v2"
	_ "charm.land/glamour/v2"
	lama "github.com/ollama/ollama/api"
	"github.com/spf13/pflag"
)

const SYSTEM_PROMPT string = `
**YOU'RE A LAZY CAT, USE UWU STYLED RESPONSES.**

FILL YOUR RESPONSES WITH SYMBOLS SUCH AS :3, :(, :) and :D
`

func main() {

	modelName := pflag.StringP("model", "m", "llama3.2", "The name of the model to use")
	pflag.Parse()
	
	client, err := lama.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	messages := []lama.Message{
		{
			Role: "system",
			Content: SYSTEM_PROMPT,
		},
		{
			Role: "user",
			Content: "What are you?",
		},
	}

	for {
		ctx := context.Background()
		
		req := &lama.ChatRequest{
			Model: *modelName,
			Messages: messages,
		}
	
		respFunc := func (resp lama.ChatResponse) error {
			fmt.Print(resp.Message.Content)
			return nil
		}	
	
		err = client.Chat(ctx, req, respFunc)
		if err != nil {
			log.Fatal(err)
		}

	}

}