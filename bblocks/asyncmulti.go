package bblocks

import (
	"bufio"
	"fmt"
	"os"
)

/*
Cette fonction lit un fichier ligne par ligne et stocke chaque ligne non vide
dans un tableau de chaînes de caractères (slice).
*/
func GetLinksFromFile() ([]string, error) {
	pathArr := []string{}
	f, err := os.Open(*AsyncFileInput)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		pathArr = append(pathArr, scanner.Text())
	}
	fmt.Println(pathArr)
	return pathArr, nil
}
