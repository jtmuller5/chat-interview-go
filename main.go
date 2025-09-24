package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

/*
To read input from the command line:
rl := bufio.NewReader(os.Stdin)
input, _ := rl.ReadString('\n')
input = strings.TrimSpace(input)
fmt.Println("You entered:", input)
*/

/*
Generate AI response:
ctx := context.Background()
	result, _ := ai.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text("Explain how AI works in a few words"),
		nil,
	)

fmt.Printf("AI Response: %s\n", result.Text())
*/

// https://ai.google.dev/gemini-api/docs/text-generation#go
var ai *genai.Client
var rl *bufio.Reader

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err == nil {
		ai = client
	}

	// Create readline interface
	rl = bufio.NewReader(os.Stdin)
}

func promptUser(message string) {
	fmt.Print(message + "\n")
	answer, err := rl.ReadString('\n')
	if err != nil {
		log.Printf("Error reading input: %v", err)
		return
	}

	answer = strings.TrimSpace(answer)
	fmt.Printf("Echo: %s\n", answer)
}

func main() {
	promptUser("What message would you like to leave?")
}
