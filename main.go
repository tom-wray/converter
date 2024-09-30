package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	sourceParent = "docs"
	sourceDir    = "docs"
	outputDir    = "./output"
)

func main() {
	// Clean output directory
	err := os.RemoveAll(outputDir)
	if err != nil {
		fmt.Println("Error removing output directory:", err)
		return
	}

	// Walk through source directory and process files
	err = filepath.Walk(sourceDir, processFile)
	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}

// processFile handles each file or directory encountered during the filepath.Walk
// It creates an index for directories and updates markdown files
func processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return createIndex(path, outputDir)
	}

	if filepath.Ext(path) == ".md" {
		return updateFile(path, info, outputDir)
	}

	return nil
}

// createIndex generates an index.md file for each directory, except the root
func createIndex(path string, outputDir string) error {
	if path == sourceDir {
		// Do not create index.md in the root folder
		return nil
	}

	title := filepath.Base(path)
	content := fmt.Sprintf("---\ntitle: %s\n---", title)

	indexPath := filepath.Join(outputDir, path, "index.md")
	err := os.MkdirAll(filepath.Dir(indexPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", filepath.Dir(indexPath), err)
	}

	err = os.WriteFile(indexPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error creating index file %s: %w", indexPath, err)
	}

	fmt.Println("Created:\t index.md in", indexPath)
	return nil
}

// updateFile processes each markdown file, adding front matter and updating its content
func updateFile(path string, info os.FileInfo, outputDir string) error {
	title := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	parent := ""

	if dir := filepath.Dir(path); dir != sourceParent {
		parent = filepath.Base(dir)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", path, err)
	}

	newContent := fmt.Sprintf("---\ntitle: %s\n%slayout: %s\n---\n%s", title, getIndent(parent), "default", string(content))

	outputPath := filepath.Join(outputDir, path)
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", filepath.Dir(outputPath), err)
	}

	err = os.WriteFile(outputPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing file %s: %w", outputPath, err)
	}

	fmt.Println("Updated:\t", outputPath)
	return nil
}

// getIndent returns the parent front matter field if a parent exists, otherwise an empty string
func getIndent(parent string) string {
	if parent != "" {
		return fmt.Sprintf("parent: %s\n", parent)
	}
	return ""
}
