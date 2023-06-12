package parser

import (
	"os"
)

// Parser is a fast configuration parser.
type Parser struct {
}

// Parse will parse libsshd configuration file.
func (ctx *Parser) Parse(file *os.File) (error, *Header, []*Item) {
	return nil, &Header{}, nil
}
