# AI Connector - Trustami.AI

This repository aims to provide a simple and easy to use connector to the Trustami.AI API.

### Usage
```golang
package main

import (
	"fmt"
	ai "github.com/trustami/ai-connector.go"
)

func main() {
	token := "YOUR_TOKEN_HERE"
	text := "This is a test text. It is used to test the Trustami.AI API."

	// Detect the language of the text
	langDetector := ai.NewLangDetector(token)

	language, err := langDetector.DetectLanguage(text)
	if err != nil {
		panic(err)
	}

	fmt.Println(language.Language) // en
}
```