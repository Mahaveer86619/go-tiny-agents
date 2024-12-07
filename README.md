# Go Tiny Agents 🤖

A Go-based multi-agent system that simulates characters from "The Office" using Google's Gemini AI. Each agent has a unique personality and responds to messages based on their character traits.

## 🌟 Features

- Multiple AI agents with distinct personalities
- Character-based response generation using Gemini 1.5
- Memory sharing between agents
- RESTful API interface
- Configurable personalities via JSON

## 🚀 Getting Started

### Prerequisites

- Go 1.23.4 or higher
- Gemini API key from https://aistudio.google.com
- `.env` file with GOOGLE_GENAI_API_KEY

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/xprilion/go-tiny-agents.git
   cd go-tiny-agents
   ```

2. Create a `.env` file:
   ```bash
   GOOGLE_GENAI_API_KEY=your_api_key_here
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run .
   ```

## 📝 Usage

### API Endpoints

#### POST /message
Send a message to the agents.

**Request Body:**
```json
{
  "message": "Hello, how are you?"
}
```

**Response:**
```json
[
  {
    "name": "Michael Scott",
    "message": "Hello, how are you?"
  }
]
```

## 🤝 Contributing
You can contribute to this project by adding more personalities or improving the existing ones. PRs are welcome!

## 📄 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 👥 Authors

- [@xprilion](https://github.com/xprilion)

## 🙏 Acknowledgments

- [Firebase GenKit](https://github.com/firebase/genkit) for the AI SDK
- [The Office](https://www.imdb.com/title/tt0386676/) for the inspiration
- [Go](https://go.dev/) for the programming language
- [Gemini](https://ai.google.dev/gemini) for the AI model

