package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var ErrInputOutOfRange = errors.New("input out of range")

func CloseWarningError(subject io.Closer) {
	err := subject.Close()
	if err != nil {
		println("Warning: Couldn't close stream:", err.Error())
	}
}

func PromptOneLine(prompt string) (string, error) {
	defer breakLine()
	var sb strings.Builder
	var stdinReader = bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	for {
		line, isPrefix, err := stdinReader.ReadLine()
		if err != nil {
			return "", err
		}
		sb.Write(line)
		if !isPrefix {
			break
		}
	}
	return sb.String(), nil
}

func MustPromptOneLine(prompt string) string {
	input, err := PromptOneLine(prompt)
	if err == nil {
		panic(err)
	}
	return input
}

func PromptBoolean(prompt string) (bool, error) {
	prompt = fmt.Sprintf("%s [y/N]", prompt)
	answer, err := PromptOneLine(prompt)
	if err != nil {
		return false, err
	}
	return strings.ToLower(answer) == "y", nil
}

func MustPromptBoolean(prompt string) bool {
	answer, err := PromptBoolean(prompt)
	if err != nil {
		panic(err)
	}
	return answer
}

func PromptInteger(prompt string) (int, error) {
	input, err := PromptOneLine(prompt)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(input)
}

func ChooseFromOptions(prompt string, options []string) (int, error) {
	for i, option := range options {
		fmt.Printf("%d) %s\n", i+1, option)
	}
	input, err := PromptInteger(prompt)
	if err == nil && input < 1 || input > len(options) {
		err = ErrInputOutOfRange
	}
	if err != nil {
		return 0, fmt.Errorf("invalid user input: %w", err)
	}
	return input, nil
}

func breakLine() {
	fmt.Println()
}
