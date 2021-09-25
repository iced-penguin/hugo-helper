package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Prompt struct {
	Categories []string
}

func NewPrompt(categories []string) Prompt {
	return Prompt{
		Categories: categories,
	}
}

func (p Prompt) Input() (filename, category string) {
	filename = p.inputFilename()
	category = p.inputCategory()
	return
}

func (p Prompt) inputFilename() string {
	return inputEach("Enter file name (without extension)")
}

func (p Prompt) inputCategory() string {
	categoriesMap := make(map[string]string)
	for i, c := range p.Categories {
		categoriesMap[strconv.Itoa(i+1)] = c
	}
	key := inputEach(getMsgCategory(p.Categories))
	category, ok := categoriesMap[key]
	if !ok {
		return p.inputCategory()
	}
	return category
}

func inputEach(msg string) (in string) {
	fmt.Println(msg)
	fmt.Print(">>> ")
	fmt.Scan(&in)
	return
}

func getMsgCategory(categories []string) string {
	var s []string
	for i, c := range categories {
		s = append(s, fmt.Sprintf("%d: %s", i+1, c))
	}
	return fmt.Sprintf("Choose category (%s)", strings.Join(s, ", "))
}
