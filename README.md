## Task 

Create a CLI chat bot to take messages for a bakery. The bot should:
1. Ask the user what message they would like to leave
2. Ask for their first and last name
3. Repeat the message back to the user at the end of the conversation to confirm it is correct
3a. If yes, say thank you and end the call
3b. If no, clear the current message and start over

### CLI Shortcuts
- Pressing "q" should exit the chat
- Pressing "c" should print out the formatted conversation
- Pressing "m" should print out the user's message using the following format:

```
[First Last]: Message
```

For example:
[John Doe]: I was wondering what time the bakery opened.

## Setup
1. `go mod tidy` (to install dependencies)
2. Add `GEMINI_API_KEY={your_api_key}` to `.env`
3. `go run main.go` or `go build -o chatbot main.go && ./chatbot`

## Starting Assumptions
- The user's response to each prompt will be correct
- The user will not give more information than what is requested