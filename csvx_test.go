package csvx

import "testing"

func sliceEq(a, b []string) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParse(t *testing.T) {
	c, err := Parse("test/simple.csv")
	if err != nil {
		t.Error(err)
	}

	headers := []string{"name", "age", "occupation"}
	if !sliceEq(c.Headers, headers) {
		t.Error("incorrect headers")
	}

	if len(c.Rows) != 2 {
		t.Error("rows are invalid")
	}
}

func TestRowHeaders(t *testing.T) {
	c, err := Parse("test/simple.csv")
	if err != nil {
		t.Error(err)
	}

	row := c.Rows[0]
	if !sliceEq(row.Headers, c.Headers) {
		t.Error("headers are invalid")
	}
}
