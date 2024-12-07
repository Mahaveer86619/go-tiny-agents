package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/xprilion/go-tiny-agents/types"
)

type Router struct {
	agents     []types.Agent
	memoryChan chan string
}

// Add this struct for the response format
type AgentResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func NewRouter() *Router {
	return &Router{
		agents:     make([]types.Agent, 0),
		memoryChan: make(chan string, 100),
	}
}

func (r *Router) RegisterAgent(agent types.Agent) {
	r.agents = append(r.agents, agent)
}

func (r *Router) HandleMessage(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request
	var payload struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if payload.Message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}

	// Track which agents should respond
	var respondingAgents []types.Agent

	// Find all appropriate agents based on message content
	for _, agent := range r.agents {
		for _, invocation := range agent.GetPersonality().Invocations {
			if strings.Contains(strings.ToLower(payload.Message), strings.ToLower(invocation)) {
				respondingAgents = append(respondingAgents, agent)
				break // Break inner loop to avoid adding same agent multiple times
			}
		}
	}

	// If we found matching agents, let them all respond
	if len(respondingAgents) > 0 {
		responses := make([]AgentResponse, 0)
		for _, agent := range respondingAgents {
			response := AgentResponse{
				Name:    agent.GetPersonality().Name,
				Message: agent.ProcessMessage(payload.Message),
			}
			responses = append(responses, response)
		}

		// Set content type header
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses)
		return
	}

	// Default to Michael Scott if no specific agents matched
	for _, agent := range r.agents {
		if agent.GetPersonality().Name == "Michael Scott" {
			response := []AgentResponse{{
				Name:    "Michael Scott",
				Message: agent.ProcessMessage(payload.Message),
			}}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Error case
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]AgentResponse{{
		Name:    "System",
		Message: "Error: Michael Scott not found",
	}})
}
