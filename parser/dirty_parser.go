package parser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// DirtyParser is a raw configuration parser.
type DirtyParser struct {
}

// DirtyParse will parse dirty ssh and sshd configuration file.
func (ctx *DirtyParser) DirtyParse(file *os.File) (error, *Header, []*DirtyItem) {

	reader := bufio.NewReader(file)

	// slice cap seem to max(sshd_config, ssh_config)
	parsedLines := make([]*DirtyItem, 128)

	// read all lines.
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else {
			if err != nil {
				return err, nil, nil
			}
		}

		trLine := strings.TrimLeft(string(line), " ")

		// skip the empty lines and annotates.
		if strings.HasPrefix(trLine, "#") || len(trLine) <= 0 {
			continue
		}

		// create a dirty item and fill it.
		item := DirtyItem{}

		if err := item.Fill(string(line)); err != nil {
			return err, nil, nil
		}

		parsedLines = append(parsedLines, &item)
	}

	return nil, nil, parsedLines
}
