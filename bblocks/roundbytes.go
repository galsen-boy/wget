package bblocks

import (
	"fmt"
	"math"
)

/*
Cette fonction RoundBytes est conçue pour convertir une taille en octets en
une chaîne de caractères représentant la taille en gigaoctets (Go) ou mégaoctets (Mo), arrondie à deux décimales.
*/
func RoundBytes(in_bytes int64) string {
	sizeInMB := float64(in_bytes) / (1024 * 1024) //La taille en mégaoctets (sizeInMB) est divisée par 1024 pour obtenir la taille en gigaoctets.

	// Round to gigabytes if size is 1024 MB or larger
	sizeInGB := math.Round(sizeInMB/1024*100) / 100
	/*Condition pour GB ou MB : La fonction retourne la taille en Go si elle est supérieure ou égale à 1 Go, sinon elle retourne la taille en Mo.*/
	if sizeInGB >= 1 {
		return fmt.Sprintf("%.2f GB", sizeInGB)
	} else {
		return fmt.Sprintf("%.2f MB", sizeInMB)
	}
}
