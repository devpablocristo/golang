package tests

import (
	"os"
	"path/filepath"
	"runtime"
)

func LoadTestData() ([]byte, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	filePath := filepath.Join(basePath, "tests-data", "event-data.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
