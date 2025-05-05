package main

import (
	"fmt"
	"log"
	"os"

	"yaml-parser/yamlparser"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 1 {
		log.Fatal("Необходимо указать один YAML-файл")
	}

	data, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	result, err := yamlparser.ParseYAML(data)
	if err != nil {
		log.Fatalf("Ошибка парсинга YAML: %v", err)
	}

	// Выводим результат
	for m, n := range result {
		fmt.Println(m, n)
	}
}
