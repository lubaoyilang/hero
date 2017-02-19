package hero

import (
	"bytes"
	"testing"
)

func TestExecCommand(t *testing.T) {
	// test for panic
	execCommand("")
	execCommand("ls")
}

func TestFormatUint(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{
		{in: 0, out: "0"},
		{in: 1, out: "1"},
		{in: 100, out: "100"},
		{in: 101, out: "101"},
	}

	buffer := new(bytes.Buffer)

	for _, c := range cases {
		FormatUint(c.in, buffer)
		if buffer.String() != c.out {
			t.Fail()
		}
		buffer.Reset()
	}
}
