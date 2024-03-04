package credentials

import (
	"bufio"
	"log"
	"os"
	"strings"

	"day_06/internal/repository"
)

type LogPass struct {
	AdminLogin    string `json:"username" binding:"required"`
	AdminPassword string `json:"password" binding:"required"`
}

func GetCredentials(filePath string) (LogPass, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return LogPass{}, nil
	}
	defer file.Close()

	var credentials LogPass

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		switch parts[0] {
		case "ADMIN_LOGIN":
			credentials.AdminLogin = parts[1]
		case "ADMIN_PASS":
			credentials.AdminPassword = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	return credentials, nil
}

func GetDBConfig(filePath string) (repository.Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return repository.Config{}, err
	}
	defer file.Close()

	var config repository.Config

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		switch parts[0] {
		case "HOST":
			config.Host = parts[1]
		case "PORT":
			config.Port = parts[1]
		case "USERNAME":
			config.Username = parts[1]
		case "PASSWORD":
			config.Password = parts[1]
		case "DBNAME":
			config.DBName = parts[1]
		case "SSLMODE":
			config.SSLMode = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	return config, nil
}
