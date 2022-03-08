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
			&FileInfo{"a/b", "c", "txt"},
		},
		{
			"a/b/d",
			&FileInfo{"a/b", "d", ""},
		},
	}

	for _, tc := range testCases {
		assert.DeepEqual(t, ParseFile(tc.fp), tc.expected)
	}
}

func TestFullPath(t *testing.T) {
	path := "a/b/c.txt"
	assert.Equal(t, path, ParseFile(path).FullPath())
}

func TestChangeExt(t *testing.T) {
	path := "a/b/c.txt"
	fi := ParseFile(path)
	fi.Ext = "md"
	assert.Equal(t, "a/b/c.md", fi.FullPath())
}
