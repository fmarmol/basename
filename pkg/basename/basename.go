package basename

import (
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Basename string
	Ext      string
}

func ParseFile(fp string) *FileInfo {
	ext := filepath.Ext(fp)

	base := filepath.Base(fp)
	base = strings.ReplaceAll(base, ext, "")

	ext = strings.TrimLeft(ext, ".")

	return &FileInfo{Basename: base, Ext: ext}
}
