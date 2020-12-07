package utils

import (
	"fmt"
	"os"
)

type FilePartWriter struct {
	Num   int
	Dir   string
	Size  int
	count int
	fp    *os.File
}

func (fw *FilePartWriter) filename() string {
	name := fmt.Sprintf("%s/part-%05d", fw.Dir, fw.Num)
	fw.Num++
	return name
}

func (fw *FilePartWriter) open() error {
	if fw.fp != nil {
		if err := fw.fp.Close(); err != nil {
			return err
		}
	}
	fp, err := os.Create(fw.filename())
	if err != nil {
		return err
	}
	fw.fp = fp
	return nil
}

func (fw *FilePartWriter) Close() error {
	return fw.fp.Close()
}

func (fw *FilePartWriter) Write(b []byte) (int, error) {
	if fw.count > fw.Size {
		if err := fw.open(); err != nil {
			return -1, err
		}
		fw.count = 0
	}
	n, err := fw.fp.Write(b)
	if err != nil {
		return n, err
	}
	fw.count += n
	return n, nil
}

func (fw *FilePartWriter) WriteLine(line string) (int, error) {
	n1, err := fw.Write([]byte(line))
	if err != nil {
		return n1, err
	}
	n2, err := fw.Write([]byte("\n"))
	if err != nil {
		return n1 + n2, err
	}
	return n2, nil
}

func NewFilePartWriter(dirname string, size int) (*FilePartWriter, error) {
	fw := &FilePartWriter{
		Num:   0,
		Dir:   dirname,
		Size:  size,
		count: 0,
		fp:    nil,
	}
	if err := fw.open(); err != nil {
		return nil, err
	}
	return fw, nil
}
