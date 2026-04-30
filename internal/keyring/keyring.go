package keyring

import (
	"os"
	"path/filepath"
)

var tokenFile = filepath.Join(os.ExpandEnv("$HOME"), ".config", "discordo", "token")

func GetToken() (string, error) {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SetToken(s string) error {
	dir := filepath.Dir(tokenFile)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return os.WriteFile(tokenFile, []byte(s), 0600)
}

func DeleteToken() error {
	return os.Remove(tokenFile)
}
