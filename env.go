package env

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Load(fileName string) error {
	if fileName == "" {
		fileName = ".env"
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			if err == io.EOF {
				break
			}
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, `"'`)

		os.Setenv(key, val)

		if err == io.EOF {
			break
		}
	}

	return nil
}
