package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	dir := "C:\\Users\\myahm\\OneDrive\\Documents\\Whatsgram\\gramsrv\\data\\langpack"
	
	oldStr := []byte("telegram.org")
	newStr := []byte("whatsgram.org")

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		
		if bytes.Contains(content, oldStr) {
			newContent := bytes.ReplaceAll(content, oldStr, newStr)
			err = os.WriteFile(path, newContent, 0644)
			if err != nil {
				return err
			}
			fmt.Printf("Updated %s\n", path)
		}
		
		// Case-insensitive replace for Telegram.org (if any)
		oldStr2 := []byte("Telegram.org")
		newStr2 := []byte("Whatsgram.org")
		if bytes.Contains(content, oldStr2) {
			content, _ = os.ReadFile(path)
			newContent := bytes.ReplaceAll(content, oldStr2, newStr2)
			err = os.WriteFile(path, newContent, 0644)
			if err != nil {
				return err
			}
			fmt.Printf("Updated %s\n", path)
		}
		
		return nil
	})
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
