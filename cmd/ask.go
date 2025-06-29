package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Ask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: thynk ask \"your question\"")
		return
	}

	token := os.Getenv("HF_TOKEN")
	if token == "" {
		fmt.Println("❌ HF_TOKEN not found in environment.")
		return
	}

	question := args[0]
	payload := map[string]string{"inputs": question}
	jsonData, _ := json.Marshal(payload)

	// 🔁 Replace this URL if using a different model
	url := "https://api-inference.huggingface.co/models/meta-llama/Llama-3.1-8B-Instruct"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("API call failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("❌ API error: %s\n%s\n", resp.Status, string(body))
		return
	}

	// HuggingFace usually returns a list of generated_text responses
	var result []map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON decode error:", err)
		fmt.Println("Raw body:", string(body))
		return
	}

	if len(result) > 0 {
		fmt.Println("🤖 Answer:", result[0]["generated_text"])
	} else {
		fmt.Println("⚠️ No answer received.")
	}
}
