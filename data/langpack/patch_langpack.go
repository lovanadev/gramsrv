package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type LanguagePack struct {
	Schema int `json:"schema"`
	Packs  []struct {
		Name      string `json:"name"`
		Languages []struct {
			LangCode        string `json:"lang_code"`
			Name            string `json:"name"`
			NativeName      string `json:"native_name"`
			PluralCode      string `json:"plural_code"`
			Official        bool   `json:"official"`
			Rtl             bool   `json:"rtl,omitempty"`
			StringsCount    int    `json:"strings_count"`
			TranslatedCount int    `json:"translated_count"`
			TranslationsUrl string `json:"translations_url"`
			Version         int    `json:"version"`
			File            string `json:"file"`
			Sha256          string `json:"sha256"`
		} `json:"languages"`
	} `json:"packs"`
}

func fileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	baseDir := `c:\Users\myahm\OneDrive\Documents\Whatsgram\gramsrv\data\langpack`
	jsonPath := filepath.Join(baseDir, "official-language-packs.json")

	b, err := os.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}

	var data LanguagePack
	if err := json.Unmarshal(b, &data); err != nil {
		panic(err)
	}

	versionRe := regexp.MustCompile(`_v(\d+)\.strings$`)
	updatedCount := 0

	for i := range data.Packs {
		pack := &data.Packs[i]
		for j := range pack.Languages {
			lang := &pack.Languages[j]
			if lang.File == "" {
				continue
			}

			// Only process Android for speed, or remove this to process all
			if pack.Name != "android" && pack.Name != "android_x" {
				continue
			}

			absFile := filepath.Join(baseDir, lang.File)
			contentBytes, err := os.ReadFile(absFile)
			if err != nil {
				fmt.Printf("File not found: %s\n", absFile)
				continue
			}
			content := string(contentBytes)

			newContent := strings.ReplaceAll(content, "Telegram", "Whatsgram")
			newContent = strings.ReplaceAll(newContent, "telegram", "whatsgram")
			newContent = strings.ReplaceAll(newContent, "TELEGRAM", "WHATSGRAM")

			if newContent == content {
				continue
			}

			lang.Version++
			newRelFile := versionRe.ReplaceAllString(lang.File, fmt.Sprintf("_v%d.strings", lang.Version))
			if newRelFile == lang.File {
				newRelFile = strings.Replace(lang.File, ".strings", fmt.Sprintf("_v%d.strings", lang.Version), 1)
			}
			newAbsFile := filepath.Join(baseDir, newRelFile)

			if err := os.WriteFile(newAbsFile, []byte(newContent), 0644); err != nil {
				panic(err)
			}
			if newAbsFile != absFile {
				os.Remove(absFile)
			}

			hash, err := fileSHA256(newAbsFile)
			if err != nil {
				panic(err)
			}

			lang.File = newRelFile
			lang.Sha256 = hash

			fmt.Printf("Updated %s -> %s\n", lang.Name, newRelFile)
			updatedCount++
		}
	}

	if updatedCount > 0 {
		outBytes, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile(jsonPath, outBytes, 0644); err != nil {
			panic(err)
		}
		fmt.Printf("Successfully updated %d language packs and official-language-packs.json!\n", updatedCount)
	} else {
		fmt.Println("No changes needed.")
	}
}
