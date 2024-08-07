package bblocks

import (
	"time"
)

/*
Cette fonction utilise le package time pour obtenir la date et l'heure actuelles,
puis formate et affiche cette information en fonction du paramètre start
*/
func DisplayDate(start bool) {
	dt := time.Now()

	x := dt.Format("2006-01-02 15:04:05")

	if start {
		outputFunc("start at " + x + "\n")
	} else {
		outputFunc("finished at " + x + "\n")
	}
}
