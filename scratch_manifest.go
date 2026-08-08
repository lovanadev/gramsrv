package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Rarity struct {
	Kind string `json:"kind"`
	Permille *int `json:"permille"`
}

type Model struct {
	Name string `json:"name"`
	Rarity Rarity `json:"rarity"`
}

type UpgradeAttributeSet struct {
	GiftID int64 `json:"gift_id"`
	Models []Model `json:"models"`
}

func main() {
	b, err := os.ReadFile("C:\\Users\\myahm\\OneDrive\\Documents\\Whatsgram\\gramsrv\\data\\official-gifts\\manifest.json")
	if err != nil {
		panic(err)
	}

	var root struct {
		UpgradeAttributeSets []UpgradeAttributeSet `json:"upgrade_attribute_sets"`
	}
	if err := json.Unmarshal(b, &root); err != nil {
		panic(err)
	}

	for _, entry := range root.UpgradeAttributeSets {
		seen := make(map[string]int)
		for _, m := range entry.Models {
			seen[m.Name]++
			if seen[m.Name] == 2 {
				fmt.Printf("GiftID %d has duplicate model name: %s\n", entry.GiftID, m.Name)
			}
		}
	}
}
