package bblocks

import (
	"os"
	"path/filepath"
)

/*
La fonction DetermineFilePath a pour objectif de déterminer le chemin
complet d'un fichier en fonction d'un chemin d'accès fourni par l'utilisateur.
*/
func DetermineFilePath(outputFileName string) string {

	if *New_file_path != "" {
		cleanedPath := filepath.Clean(*New_file_path)                       //Nettoyage du chemin
		homeDir, _ := os.UserHomeDir()                                      //Obtention du répertoire personnel de l'utilisateur
		filePath := filepath.Join(homeDir, cleanedPath[1:], outputFileName) //Construit le chemin complet vers le fichier en utilisant filepath.Join.
		//Mise à jour de la variable globale New_file_path et retour du chemin
		*New_file_path = filePath
		return filePath
	}
	return outputFileName
}
