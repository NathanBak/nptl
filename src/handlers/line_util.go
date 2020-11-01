package handlers

import (
	"bufio"
	"io"
)

// readLine reads and returns the next line and the newline marker.  If an io.EOF error is returned
// it means there are no more lines to be read.  Any other error should be taken as an error.  The
// line will never be nil for a successful invocation, but may be empty.  The newline will never be
// nil but may be empty.
func readLine(br *bufio.Reader) ([]byte, []byte, error) {

	line, err := br.ReadSlice('\n')

	newline := []byte{'\n'}

	if err == bufio.ErrBufferFull {
		// Handle the case where "\r\n" straddles the buffer.
		if len(line) > 0 && line[len(line)-1] == '\r' {
			newline = []byte{'\r', '\n'}
			line = line[:len(line)-1]

		}
		return line, newline, nil
	}

	if err == io.EOF {
		return line, []byte{}, err

	}

	if err != nil {
		return line, newline, err
	}

	if line[len(line)-1] == '\n' {
		drop := 1
		if len(line) > 1 && line[len(line)-2] == '\r' {
			drop = 2
			newline = []byte{'\r', '\n'}
		}
		line = line[:len(line)-drop]
	}

	return line, newline, err
}
