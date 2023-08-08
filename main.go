package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Input struct {
	Sentence string
	Key      string
}

type Document struct {
	sentences []string
	key       string
}

func splitStringWithRegex(input, pattern string) []string {
	regex := regexp.MustCompile(pattern)
	parts := regex.Split(input, -1) // -1 means no limit on the number of splits
	return parts
}

func getTerms(input Input) []string {
	tokens := splitStringWithRegex(input.Sentence, "\\s")
	regex := regexp.MustCompile("[^a-zA-Z]")
	for i, item := range tokens {
		tokens[i] = regex.ReplaceAllString(strings.ToLower(item), "")
	}
	return tokens
}

func getDocument(input Input) Document {
	return Document{
		sentences: getTerms(input),
		key:       input.Key,
	}
}

func main() {
	inputs := getInputs("data/sentences.json")
	for _, input := range inputs {
		document := getDocument(input)
		fmt.Println(document.key)
	}
}

func getInputs(filepath string) []Input {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var inputs []Input
	decoder := json.NewDecoder(file)

	decoder.Decode(&inputs)
	return inputs
}
