package main

func main() {
	// setup llm to localclient
	client := llamacpp.New("../llama.cpp", "models/mistral/mistral_7b_v1.gguf")

	messages := []struct {
		Username string
		Text     string
	}{
		{
			Username: "RuggMatt",
			Text:     "hey pedro? where should i buy a pizza from?",
		},
		{
			Username: "Archeious",
			Text:     "Run the server in a container with the -always-restart flag.",
		},
		{
			Username: "Archeious",
			Text:     "That is dirty and janky but it works.",
		},
		{
			Username: "Archeious",
			Text:     "damn it is opinionated.",
		},
		{
			Username: "Archeious",
			Text:     "LC delivers",
		},
		{
			Username: "Archeious",
			Text:     "But if you want good/sbobby pizza then Firebird is the bestest.",
		},
		{
			Username: "Archeious",
			Text:     "LOL FB is like 100 yards west of the Purple Turtle.",
		},
		{
			Username: "Archeious",
			Text:     "Apparently my job is to distract you.",
		},
		{
			Username: "Archeious",
			Text:     "Pedro, why is Purple Turtle the best?",
		},
		{
			Username: "soy_llm_bot",
			Text:     "Hello, my name is Pedro_el_asistente I am here to help you.",
		},

	messageHistory := []llms.MessageContent{llms.TextParts(llms.ChatMessageTypeSystem, pedroPrompt+"\nHere is the twitch chat history for you to respond to:")}
	var twitchChatHistory []string
	for _, message := range messages {
		twitchChatHistory = append(twitchChatHistory, fmt.Sprintf("%s: %s", message.Username, message.Text))
	}
	messageHistory = append(messageHistory, llms.TextParts(llms.ChatMessageTypeHuman, twitchChatHistory...))

	// pass prompts to llm
	client.GenerateContent("Tell me a dad joke that is not about fish.")

	// and compare the results from llm
}
