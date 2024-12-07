package types

type Agent interface {
	ProcessMessage(message string) string
	GetPersonality() *Personality
}

type Personality struct {
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	Personality string   `json:"personality"`
	Invocations []string `json:"invocations"`
}

type BaseAgent struct {
	Personality *Personality
	MemoryChan  chan string
}
