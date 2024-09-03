package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	inputDir := "input"
	outputDir := "output"

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	// Walk through the input directory
	err := filepath.Walk(inputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".ts" {
			// Construct the output file path
			relPath, _ := filepath.Rel(inputDir, path)
			outputPath := filepath.Join(outputDir, filepath.Base(relPath[:len(relPath)-len(filepath.Ext(relPath))]+".mp4"))

			// Prepare the ffmpeg command
			cmd := exec.Command("./ffmpeg.exe", "-i", path, "-map", "0:v", "-map", "0:a:0", "-c:v", "copy", "-c:a", "copy", outputPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Execute the command
			fmt.Printf("Converting %s to %s\n", path, outputPath)
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error converting file %s: %v\n", path, err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", inputDir, err)
	}
}
