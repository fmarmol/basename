package basename

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParseFile(t *testing.T) {
	testCases := []struct {
		fp       string
		expected *FileInfo
	}{
		{
			"a/b/c.txt",
			&FileInfo{"c", "txt"},
		},
		{
			"a/b/d",
			&FileInfo{"d", ""},
		},
	}

	for _, tc := range testCases {
		assert.DeepEqual(t, ParseFile(tc.fp), tc.expected)
	}
}
