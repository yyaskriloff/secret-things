package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func findPoint(f *os.File, offset int64) (int64, error) {
	_, err := f.Seek(offset, io.SeekStart)
	if err != nil {
		return -1, err
	}

	reader := bufio.NewReader(f)
	data, err := reader.ReadBytes('\n')

	if err != nil {
		return -1, err
	}

	newlinePos := offset + int64(len(data)) - 1
	return newlinePos, nil

}

func Parse(path string) (map[string]string, error) {
	vars := make(map[string]string)
	filePath, _ := filepath.Abs(path)
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var offset int64 = 0

	for {
		// break case
		pos, err := findPoint(f, offset)
		if err != nil && err == io.EOF {
			return vars, nil
		} else if err != nil {
			return nil, err
		}

		buffer := make([]byte, pos-offset)

		bytes, err := f.ReadAt(buffer, offset)

		if err != nil {
			return nil, err
		}

		convertedString := string(buffer[:bytes])

		name, value, found := strings.Cut(convertedString, "=")

		if !found {
			return nil, errors.New("ENV var misconfigured")
		}

		trimedQuote := strings.Trim(value, "'")
		trimeParenthesis := strings.Trim(trimedQuote, "\"")

		vars[name] = trimeParenthesis

		offset += pos + 1
	}

}

func WriteEnv(vars map[string]string, fileName string) {

	var b bytes.Buffer
	for key, value := range vars {
		fmt.Fprintf(&b, "%s=\"%s\"\n", key, value)
	}

	filePath, _ := filepath.Abs(fileName)
	// potential issue if fail to write to file we just trunicated and lost all vars
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create/truncate file: %v", err)
	}
	defer f.Close()

	_, err = f.Write(b.Bytes())
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}
