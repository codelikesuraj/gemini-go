package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const MODEL_NAME = "gemini-pro"

var EXIT_CMD = []string{"EXIT", "QUIT", "exit", "quit"}

func main() {
	clearScreen()

	fmt.Print("-----------------------------------------------")
	fmt.Print("\nWelcome to the Gemini AI REPL by CODELIKESURAJ\n")
	fmt.Print("\nCommands:\n")
	fmt.Print("\t'QUIT', 'EXIT': end the program\n")
	fmt.Print("\t'UPDATE_KEY':   update API key\n")
	fmt.Print("------------------------------------------------\n\n")

	api_key, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok || api_key == "" {
		updateApiKey()
		api_key = os.Getenv("GEMINI_API_KEY")
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

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

		if slices.Contains(EXIT_CMD, cleaned) {
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

func updateApiKey() {
	clearScreen()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your gemini API key > ")
		scanner.Scan()
		text := scanner.Text()
		trimmed := strings.TrimSpace(text)

		switch {
		case trimmed == "":
			clearScreen()
			fmt.Print("Error: please enter a valid key!\n")
			continue
		case slices.Contains(EXIT_CMD, trimmed):
			os.Exit(0)
		}

		clearScreen()

		os.Environ()
		os.Setenv("GEMINI_API_KEY", trimmed)
		break
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
