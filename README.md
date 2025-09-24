## Outline
Create a CLI chat bot to take messages for a bakery. The bot should:
1. Ask the user what message they would like to leave
2. Ask for their first and last name

### CLI Shortcuts
- Pressing "q" should exit the chat
- Pressing "c" should print out the formatted conversation
- Pressing "m" should print out the user's message using the following format:

```
[First Last]: Message
```

For example:
[John Doe]: I was wondering what time the bakery opened.

## Task 1 Assumptions
(No generative AI)
- The user's response to each prompt will be correct
- The user will not give more information than what is requested

## Task 2 Assumptions
(Will need to use Generative AI)
- The user will respond like a normal person
- There's no guarantee what the user will say

## Setup
1. `go mod tidy` (to install dependencies)
2. Add `GEMINI_API_KEY={your_api_key}` to `.env`
3. `go run main.go` or `go build -o chatbot main.go && ./chatbot`