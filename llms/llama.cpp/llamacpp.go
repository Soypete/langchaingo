package llamacpp

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/llms"
)

type LLM struct {
	CallbacksHandler callbacks.Handler
	llamacppPath     string
	modelPath        string // subdirectory of llamacppPath
}

func (m *LLM) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {

	if m.CallbacksHandler != nil {
		m.CallbacksHandler.HandleLLMGenerateContentStart(ctx, messages)
	}

	opts := &llms.CallOptions{}
	for _, opt := range options {
		opt(opts)
	}

	fmt.Println(messages)

	// // If o.client.GlobalAsArgs is true
	// if m.client.GlobalAsArgs {
	// 	// Then add the option to the args in --key=value format
	// 	m.appendGlobalsToArgs(*opts)
	// }

	// for _, msg := range messages {
	// 	part := msg.Parts[0]
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	resp := &llms.ContentResponse{
		Choices: []*llms.ContentChoice{
			{
				Content: "hello world",
			},
		},
	}

	if m.CallbacksHandler != nil {
		m.CallbacksHandler.HandleLLMGenerateContentEnd(ctx, resp)
	}

	return resp, nil
	return nil, nil
}

func (m *LLM) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return m.runExecutable(prompt)
}

func New(llamaPath, modelPath string) *LLM {
	return &LLM{
		llamacppPath: llamaPath,
		modelPath:    modelPath,
	}
}

func (m *LLM) createCommand(prompt string) string {
	cmdStr := m.llamacppPath + "/main"
	modelArgs := " -m " + m.llamacppPath + "/" + m.modelPath
	promtArgs := " -p " + "\"" + prompt + "\""

	return cmdStr + modelArgs + promtArgs + " -n 40 -e --log-disable"
}

// TODO: add options
func (m *LLM) runExecutable(prompt string) (string, error) {

	cmd := exec.Command(m.createCommand(prompt)) // make sure to pass options here
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)

	}
	defer stdout.Close()
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run llama.cpp | %w", err)
	}

	var resp []byte
	_, err = stdout.Read(resp)
	if err != nil {
		return "", fmt.Errorf("failed to read stdout | %w", err)
	}

	return string(resp), nil
}
