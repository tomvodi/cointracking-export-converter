package cointracking

import (
	"bufio"
	"github.com/gocarina/gocsv"
	"io"
	"strings"
)

type ctCsvDecoder struct {
	scanner         *bufio.Scanner
	firstRowScanned bool
}

func (c *ctCsvDecoder) GetCSVRow() ([]string, error) {
	if !c.scanner.Scan() {
		if c.scanner.Err() != nil {
			return nil, c.scanner.Err()
		} else {
			return nil, io.EOF
		}
	}

	row := c.scanner.Text()
	if !c.firstRowScanned {
		c.firstRowScanned = true
		return ctCsvHeaders, nil
	}

	fields := strings.Split(row, ",")
	for i := 0; i < len(fields); i++ {
		fields[i] = strings.ReplaceAll(fields[i], `"`, "")
		fields[i] = strings.TrimSpace(fields[i])
	}

	return fields, nil
}

func (c *ctCsvDecoder) GetCSVRows() ([][]string, error) {
	var lines [][]string
	for {
		row, err := c.GetCSVRow()
		if err == nil {
			lines = append(lines, row)
		}
		return lines, nil
	}
}

func NewCsvDecoder(reader io.Reader) gocsv.SimpleDecoder {
	return &ctCsvDecoder{
		scanner: bufio.NewScanner(reader),
	}
}
