package main

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyAhyp8aYPspZ7jDSt2gu2C_bIPgwUjbpoU"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("the dude abides"))
	if err != nil {
		log.Fatal(err)
	}
	if resp != nil {
		// dereference the pointer to access the struct fields
		candidates := resp.Candidates
		if candidates != nil {
			for _, candidate := range candidates {
				content := candidate.Content
				if content != nil {
					text := content.Parts[0]
					log.Println("Candidate Content Text: ", text)
				}
			}
		} else {
			log.Println("No candidates returned")
		}
	} else {
		log.Println("Response is empty")
	}
}
