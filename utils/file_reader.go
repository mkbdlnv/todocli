package utils

import "os"

func ReadFile(filepath string) ([]byte,error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(filepath string, data []byte) (error) {
	return os.WriteFile(filepath, data, os.ModeAppend)
}
