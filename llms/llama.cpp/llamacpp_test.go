package llamacpp

import (
	"testing"
)

func Test_createCommand(t *testing.T) {
	tests := []struct {
		name   string
		llm    *LLM
		prompt string
		want   string
	}{
		{
			name: "Test 1",
			llm: &LLM{
				llamacppPath: "../llama.cpp",
				modelPath:    "models/mistral/mistral_7b_v1.gguf",
			},
			prompt: "Tell me a dad joke that is not about fish.",
			want:   "../llama.cpp/main -m ../llama.cpp/models/mistral/mistral_7b_v1.gguf -p \"Tell me a dad joke that is not about fish.\" -n 40 -e --log-disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.llm.createCommand(tt.prompt)
			if got != tt.want {
				t.Errorf("createCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
