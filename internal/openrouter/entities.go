package openrouter

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ThoughtsExtraBody struct {
	Reasoning struct {
		Enabled bool `json:"enabled"`
	} `json:"reasoning"`
}

type ThoughtsReasoning struct {
}

type ThoughtsEnabled struct {
}

type RequestBody struct {
	Model     string            `json:"model"`
	Messages  []Message         `json:"messages"`
	ExtraBody ThoughtsExtraBody `json:"extra_body"`
}
