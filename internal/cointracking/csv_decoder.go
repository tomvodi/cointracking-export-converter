package cointracking

import (
	"bufio"
	"io"
	"strings"
)

type CsvDecoder struct {
	scanner         *bufio.Scanner
	firstRowScanned bool
}

func (c *CsvDecoder) GetCSVRow() ([]string, error) {
	if !c.scanner.Scan() {
		if c.scanner.Err() != nil {
			return nil, c.scanner.Err()
		}
		return nil, io.EOF
	}

	row := c.scanner.Text()
	if !c.firstRowScanned {
		c.firstRowScanned = true
		return ctCsvHeaders, nil
	}

	fields := strings.Split(row, "\",\"")
	for i := 0; i < len(fields); i++ {
		fields[i] = strings.ReplaceAll(fields[i], `"`, "")
		fields[i] = strings.TrimSpace(fields[i])
	}

	return fields, nil
}

func (c *CsvDecoder) GetCSVRows() ([][]string, error) {
	var lines [][]string
	for {
		row, err := c.GetCSVRow()
		if err != nil {
			return lines, nil
		}
		lines = append(lines, row)
	}
}

func NewCsvDecoder(reader io.Reader) *CsvDecoder {
	return &CsvDecoder{
		scanner: bufio.NewScanner(reader),
	}
}
