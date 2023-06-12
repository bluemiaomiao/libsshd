package parser

import (
	"errors"
	"strings"
)

// Item is a clear configuration item.
type Item struct {
	annotation string
	key        string
	value      interface{}
}

// fill an empty configuration item.
// Default annotation is "----".
func (ctx *Item) fill(key string, value interface{}) {
	ctx.key = key
	ctx.value = value
	ctx.annotation = "# ----"
}

// fillWithAnnotation fill an empty configuration item with annotation
func (ctx *Item) fillWithAnnotation(key string, value interface{}, annotation string) {
	ctx.key = key
	ctx.value = value
	ctx.annotation = "# " + annotation
}

func (ctx *Item) GetAnnotation() string {
	if ctx.annotation == "----" {
		return ""
	}

	return ctx.annotation[2 : len(ctx.annotation)-1]
}

func (ctx *Item) GetKey() string {
	return ctx.key
}

func (ctx *Item) GetValue() interface{} {
	return ctx.value
}

func (ctx *Item) Fill(line1, line2 string) error {

	line1 = strings.Trim(line1, " ")
	line2 = strings.Trim(line2, " ")

	if len(line1) <= 0 {
		return errors.New("unsupported config format")
	}

	if len(line2) <= 0 {
		return errors.New("empty config item")
	}

	if !strings.HasPrefix(line1, "#") {
		return errors.New("unsupported config format")
	}

	return nil
}
