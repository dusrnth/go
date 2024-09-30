package file

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Username     string
	Password     string
	Hostname     string
	Port         string
	Database     string
	QuertTimeout int
}

func loadConfig(filename string) (Config, error) {
	file, err := os.Open("test.txt")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	config := Config{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for len(line) > 0 && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				switch key {
				case "USERNAME":
					config.Username = value
				case "PASSWORD":
					config.Password = value
				case "HOSTNAME":
					config.Hostname = value
				case "PORT":
					config.Port = value
				case "DATABASE":
					config.Database = value
				case "QUERY_TIMEOUT":
					timeout, _ := strconv.Atoi(value)
					config.QuertTimeout = timeout
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return Config{}, err
	}

	return config, nil
}
