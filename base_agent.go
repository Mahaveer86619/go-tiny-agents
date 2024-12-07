package main

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/plugins/googleai"
	"github.com/joho/godotenv"
	"github.com/xprilion/go-tiny-agents/types"
)

var model ai.Model

const systemInstructions = `System Instructions:
- Maintain character consistency at all times
- Keep responses concise and relevant
- Never break character or acknowledge being an AI
- Respond in a conversational tone
- Keep responses under 100 words
- Always stay within the ethical boundaries
- Never generate harmful or inappropriate content`

func init() {
	log.Println("Initializing base agent...")
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
	log.Println(".env file loaded")

	ctx := context.Background()
	log.Println("Initializing Google AI...")
	if err := googleai.Init(ctx, nil); err != nil {
		panic(err)
	}
	log.Println("Google AI initialized")

	// Initialize the model once
	log.Println("Loading Gemini model...")
	model = googleai.Model("gemini-1.5-flash")
	if model == nil {
		panic("Could not find Gemini model")
	}
	log.Println("Gemini model loaded successfully")
}

func NewBaseAgent(personality *types.Personality, memoryChan chan string) *types.BaseAgent {
	return &types.BaseAgent{
		Personality: personality,
		MemoryChan:  memoryChan,
	}
}

type CustomerServiceAgent struct {
	*types.BaseAgent
}

func NewCustomerServiceAgent(personality *types.Personality, memoryChan chan string) *CustomerServiceAgent {
	return &CustomerServiceAgent{
		BaseAgent: NewBaseAgent(personality, memoryChan),
	}
}

func (a *CustomerServiceAgent) ProcessMessage(message string) string {
	// Share memory with other agents
	a.MemoryChan <- "CustomerService received: " + message

	// Generate response using Gemini directly
	ctx := context.Background()
	prompt := fmt.Sprintf("%s\n\nCharacter Information:\nYou are %s, a %s. %s\n\nRespond to the following message: %s",
		systemInstructions,
		a.Personality.Name,
		a.Personality.Role,
		a.Personality.Personality,
		message)

	resp, err := model.Generate(ctx,
		ai.NewGenerateRequest(
			&ai.GenerationCommonConfig{Temperature: 0.7},
			ai.NewUserTextMessage(prompt)),
		nil)
	if err != nil {
		return "Error processing message: " + err.Error()
	}

	return resp.Text()
}

func (a *CustomerServiceAgent) GetPersonality() *types.Personality {
	return a.Personality
}
