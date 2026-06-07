package internal

import (
	"fmt"
)

// CompareConfigsDeep recursively compares two configs
func compareConfigsDeep(config1, config2 map[string]interface{}, path string, differences *[]string) {
	// Get all unique keys
	allKeys := make(map[string]bool)
	for k := range config1 {
		allKeys[k] = true
	}
	for k := range config2 {
		allKeys[k] = true
	}

	for key := range allKeys {
		currentPath := path + key
		val1, exists1 := config1[key]
		val2, exists2 := config2[key]

		// Check key existence
		if !exists1 {
			// Key missing in first config
			*differences = append(*differences, fmt.Sprintf("  [%s] missing in first config", currentPath))
			continue
		}
		if !exists2 {
			// Key missing in second config
			*differences = append(*differences, fmt.Sprintf("  [%s] missing in second config", currentPath))
			continue
		}

		// If both values are maps, compare recursively
		map1, isMap1 := val1.(map[string]interface{})
		map2, isMap2 := val2.(map[string]interface{})

		// Also check if they are Config type
		if !isMap1 {
			if config1, ok := val1.(Config); ok {
				map1 = map[string]interface{}(config1)
				isMap1 = true
			}
		}
		if !isMap2 {
			if config2, ok := val2.(Config); ok {
				map2 = map[string]interface{}(config2)
				isMap2 = true
			}
		}

		if isMap1 && isMap2 {
			// Recursively check missing keys inside
			compareConfigsDeep(map1, map2, currentPath+".", differences)
		}
	}
}

func CompareAllConfigs(configs []Config, fileNames []string) bool {
	if len(configs) < 2 {
		fmt.Println("At least 2 configs required for comparison")
		return false
	}

	baseConfig := configs[0]
	baseFileName := fileNames[0]
	allMatch := true

	for i := 1; i < len(configs); i++ {
		var differences []string
		// Convert Config to map[string]interface{}
		compareConfigsDeep(map[string]interface{}(baseConfig), map[string]interface{}(configs[i]), "", &differences)

		if len(differences) > 0 {
			fmt.Printf("\n=== Differences between %s and %s ===\n", baseFileName, fileNames[i])
			for _, diff := range differences {
				fmt.Println(diff)
			}
			allMatch = false
		} else {
			fmt.Printf("\n✅ %s is identical to %s\n", fileNames[i], baseFileName)
		}
	}

	if allMatch {
		fmt.Println("\n✅ All configurations are identical")
	} else {
		fmt.Println("\n❌ Differences found in configurations")
	}
	return allMatch
}
