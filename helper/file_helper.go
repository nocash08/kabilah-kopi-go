package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file *multipart.FileHeader, uploadDir string) (string, error) {
	// Create the upload directory if it doesn't exist
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate unique filename
	filename := filepath.Join(uploadDir, file.Filename)

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	// Return the relative path to be stored in database
	return filename, nil
}

func DeleteFile(filepath string) error {
	// Check if file exists first
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil // File doesn't exist, just return nil
	}

	if err := os.Remove(filepath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	return nil
}
