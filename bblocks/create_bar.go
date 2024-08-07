package bblocks

import "github.com/schollz/progressbar/v3"

/*
Cette fonction crée et configure une barre de progression pour afficher
l'avancement d'une opération en fonction de la taille totale spécifiée.
*/
func CreateProgressBar(totalSize int64) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(
		int(totalSize),
		progressbar.OptionSetPredictTime(true),   //Active la prédiction du temps restant.
		progressbar.OptionSetElapsedTime(true),   //Affiche le temps écoulé.
		progressbar.OptionEnableColorCodes(true), //Active les codes couleur pour la barre de progression.
		progressbar.OptionShowBytes(true),        //Affiche la progression en octets.
		progressbar.OptionSetWidth(15),           //Définit la largeur de la barre de progression à 15 caractères.
		progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."), //Définit une description personnalisée pour la barre de progression.
		progressbar.OptionSetRenderBlankState(true),                                     // Affiche la barre de progression même si l'état est vide.
		progressbar.OptionSetWidth(35),                                                  //Redéfinit la largeur de la barre de progression à 35 caractères.
		progressbar.OptionSetDescription(FormatSize(totalSize)),                         //Redéfinit la description avec une fonction FormatSize qui devrait formater la taille totale.
		progressbar.OptionShowCount(),                                                   //Affiche le compteur de progression.
		progressbar.OptionSetTheme(progressbar.Theme{
			//Définit un thème personnalisé pour la barre de progression avec des caractères spécifiques pour différentes parties de la barre.
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	return bar
}
