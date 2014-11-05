package csvx

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type CSV struct {
	Headers []string
	Rows    []Row
	Reader  *csv.Reader
}

func NewCSV() *CSV {
	return &CSV{}
}

func (c *CSV) Parse(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	c.Reader = csv.NewReader(file)

	for {
		record, err := c.Reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if c.Headers == nil {
			c.Headers = record
			continue
		}

		row := Row{Headers: c.Headers, Fields: record}

		c.Rows = append(c.Rows, row)
	}

	return
}

type Row struct {
	Headers []string
	Fields  []string
}

func (r Row) String() string {
	return strings.Join(r.Fields, ",")
}

func (r Row) Attributes() map[string]string {
	m := make(map[string]string, len(r.Fields))
	for i, h := range r.Headers {
		m[h] = r.Fields[i]
	}
	return m
}

func Parse(filepath string) (c *CSV, err error) {
	c = NewCSV()
	err = c.Parse(filepath)
	return
}
