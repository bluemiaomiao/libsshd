package parser

import (
	"errors"
	"strings"
)

// DirtyItem is an empty dirty item.
type DirtyItem struct {
	key   string
	value string
}

// Fill an empty dirty item.
func (ctx *DirtyItem) Fill(line string) error {

	if len(line) <= 0 {
		return errors.New("content line is empty")
	}

	trLine := strings.Trim(line, " ")

	trLineElems := strings.Split(trLine, " ")

	if len(trLineElems) < 2 {
		return errors.New("error line format")
	}

	itemKey := strings.Trim(trLineElems[0], " ")

	start := len(itemKey)
	end := len(trLine) - 1

	itemValue := strings.Trim(trLine[start:end], " ")

	if len(itemValue) <= 0 {
		return errors.New("error line format")
	}

	ctx.key = itemKey
	ctx.value = itemValue

	return nil
}

func (ctx *DirtyItem) GetKey() string {
	return ctx.key
}

func (ctx *DirtyItem) GetValue() string {
	return ctx.value
}
