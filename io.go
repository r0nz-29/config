package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func getFilePath() string {
	homeDir, err := os.UserHomeDir()
	check(err)
	return filepath.Join(homeDir, FILE_NAME)
}

func createFileIfNotExists() {
	fp := getFilePath()
	if _, err := os.Stat(fp); !errors.Is(err, os.ErrNotExist) {
		return
	}
	_, err := os.Create(fp)
	check(err)
}

func getJsonFromFile() *map[string]string {
	createFileIfNotExists()
	fp := getFilePath()
	file, err := os.ReadFile(fp)
	check(err)
	jsonMap := make(map[string]string)

	if err := json.Unmarshal(file, &jsonMap); err != nil {
		panic(err)
	}

	return &jsonMap
}

func persist(jsonMapPtr *map[string]string) {
	bytes, _ := json.Marshal(*jsonMapPtr)
	fp := getFilePath()

	if err := os.WriteFile(fp, bytes, 0644); err != nil {
		print("cannot write to file!")
		panic(err)
	}
}
