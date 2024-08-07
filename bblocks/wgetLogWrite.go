package bblocks

import "fmt"

/*Cette fonction WriteToWgetLog permet d'écrire des messages dans le fichier de journalisation LogFile. Elle concatène les messages passés en arguments et les écrit dans le fichier.*/
func WriteToWgetLog(a ...any) (n int, e error) {
	str := ""
	for _, v := range a {
		str += fmt.Sprintf("%s", v)
	}
	LogFile.WriteString(str)
	return 0, nil
}
