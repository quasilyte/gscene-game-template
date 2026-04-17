package datscan

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/quasilyte/gos"
)

type WalkConfig struct {
	GameDataPath string

	Visit func(f *File) error

	Error func(f *File, err error)
}

type File struct {
	Name string

	AbsPath string
	Path    string

	Arg       float64
	StringArg string

	Kind FileKind
}

type FileKind int

const (
	FileUnknown FileKind = iota

	FileImage
	FileSound
	FileMusicOGG

	FileShaderSource
)

func Errorf(f *File, format string, args ...any) error {
	msg := "[" + f.Path + "] " + fmt.Sprintf(format, args...)
	return errors.New(msg)
}

func parseVolume(s string) (float64, error) {
	volume := 1.0
	if volStart := strings.Index(s, ".vol"); volStart != -1 {
		volEnd := strings.IndexByte(s[volStart+1:], '.')
		if volEnd != -1 {
			s := s[volStart : volStart+volEnd+1]
			s = strings.TrimPrefix(s, ".vol")
			value, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return 1, fmt.Errorf("invalid vol suffix: %v", s)
			}
			volume = float64(value) / 100
		}
	}
	return volume, nil
}

func Walk(config WalkConfig) error {
	f := &File{}

	errorf := func(format string, args ...any) {
		config.Error(f, Errorf(f, format, args...))
	}

	absDir := gos.AbsPath(config.GameDataPath)
	dirPrefix := filepath.Clean(absDir)
	if gos.IsWindows {
		dirPrefix = strings.ReplaceAll(dirPrefix, "\\", "/")
	}
	if !strings.HasSuffix(dirPrefix, "/") {
		dirPrefix += "/"
	}

	fmt.Println(absDir)
	fsys := os.DirFS(absDir)
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		f.Path = p
		f.AbsPath = path.Join(absDir, p)
		f.Kind = FileUnknown
		f.Name = d.Name()
		f.Arg = 0
		f.StringArg = ""

		fileExt := filepath.Ext(f.Name)
		switch fileExt {
		case ".png":
			f.Kind = FileImage

		case ".wav", ".ogg":
			switch fileExt {
			case ".wav":
				f.Kind = FileSound
			case ".ogg":
				f.Kind = FileMusicOGG
			}
			volume, err := parseVolume(f.Name)
			if err != nil {
				config.Error(f, err)
			}
			firstDot := strings.IndexByte(f.Name, '.')
			if firstDot != -1 {
				f.StringArg = f.Name[:firstDot]
			} else {
				f.StringArg = f.Name
			}
			f.Arg = volume

		case ".go":
			switch {
			case strings.HasSuffix(f.Name, ".shader.go"):
				f.Kind = FileShaderSource
			}
		}

		if err := config.Visit(f); err != nil {
			errorf("%s", err.Error())
		}

		return nil
	})

	return err
}

func extractTag(s, skip string) string {
	ext := filepath.Ext(strings.TrimSuffix(s, skip))
	if ext != "" {
		return ext[1:]
	}
	return ""
}
