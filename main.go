package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const MODEL_NAME = "gemini-pro"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	api_key := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// clear the screen
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Print("----------------------------------------------")
	fmt.Print("\nWelcome to the Gemini AI REPL by CODELIKESURAJ\n")
	fmt.Print("\nUsage:\n")
	fmt.Print("\tType 'EXIT' or 'QUIT' to end the program\n")
	fmt.Print("----------------------------------------------\n\n")

	runRepl(ctx, client.GenerativeModel(MODEL_NAME).StartChat())
}

func runRepl(ctx context.Context, cs *genai.ChatSession) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Me > ")
		scanner.Scan()
		cleaned := strings.TrimSpace(scanner.Text())
		if cleaned == "" {
			fmt.Print("\nGemini AI > please enter a valid non-empty message!")
			continue
		}

		switch cleaned {
		case "EXIT", "QUIT":
			fmt.Print("\nGemini AI > Session ended!!!\n\n")
			fmt.Print("----------------------------------------------\n")
			os.Exit(0)
		}

		fmt.Print("\nGemini AI > Loading...")

		prompt := genai.Text(cleaned)
		resp, err := cs.SendMessage(ctx, prompt)
		fmt.Print("\r                      \r")
		if err != nil {
			fmt.Printf("Gemini AI > Oops, something went wrong: %s\n\n", err.Error())
		} else {
			printResp(resp)
		}
	}
}

func printResp(resp *genai.GenerateContentResponse) {
	fmt.Println("Gemini AI >")
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				msg := strings.Split(strings.ReplaceAll(fmt.Sprintf("%s\n", part), "**", ""), " ")
				l := len(msg)
				for i := 0; i < l; i++ {
					time.Sleep(time.Millisecond * 75)
					s := strings.TrimSpace(msg[i])
					if len(s) > 0 {
						fmt.Printf("%s", msg[i])
					}
					if i < l-1 {
						fmt.Print(" ")
					}
				}
			}
		}
	}

	fmt.Println()
}
