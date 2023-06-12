package parser

import "testing"

func TestHeader(t *testing.T) {
	header := &Header{}

	header.SetVersion("1.0.0")

	if header.GetVersion() != "1.0.0" {
		t.Fatal("testing error for parser header struct.")
	}
}
