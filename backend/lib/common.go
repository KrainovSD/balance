package lib

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"os"
	"strings"
)

func RandomHex(length int) (string, error) {
	bytes := make([]byte, (length+1)/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes)[:length], nil
}

func LoadEnvFile(path string) error {
	var file *os.File
	var err error
	var scanner *bufio.Scanner
	var line string
	var parts []string
	var key string
	var value string

	if file, err = os.Open(path); err != nil {
		return err
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts = strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key = strings.TrimSpace(parts[0])
		value = strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return scanner.Err()
}
