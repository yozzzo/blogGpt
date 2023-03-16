package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func PostOpenAi(prompt string) string {
	fmt.Println("call_gpt3_api")
	client := openai.NewClient(os.Getenv("OPEN_AI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	fmt.Println("resp")
	fmt.Println(resp)
	fmt.Println("error")
	fmt.Println(err)
	if err != nil {
		return "Error"
	}

	msg := resp.Choices[0].Message.Content
	return msg
}
