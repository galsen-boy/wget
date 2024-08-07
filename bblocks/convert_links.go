package bblocks

import (
	"io"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

/*
Cette fonction prend en entrée un lecteur (io.Reader) pour lire le contenu HTML,
un écrivain (io.Writer) pour écrire le contenu modifié, et une URL de base (*url.URL) pour convertir les liens relatifs en liens absolus.
*/
func ConvertHTMLLinks(input io.Reader, output io.Writer, baseURL *url.URL) error {
	doc, err := html.Parse(input)
	if err != nil {
		return err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link" || n.Data == "img") {
			// Convertir les attributs href ou src
			for i, attr := range n.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					u, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					absURL := baseURL.ResolveReference(u)
					n.Attr[i].Val = absURL.String()
				}
			}
		} else if n.Data == "style" {
			// Convertir les URLs à l'intérieur des éléments style
			cssContent := strings.TrimSpace(getTextContent(n))
			cssContent = convertURLsInCSS(cssContent, baseURL)
			setTextContent(n, cssContent)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	err = html.Render(output, doc)
	if err != nil {
		return err
	}

	return nil
}

/*Fonction auxiliaire qui retourne le texte contenu dans un nœud HTML. */
func getTextContent(n *html.Node) string {
	var result string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			result += c.Data
		}
	}
	return result
}

/*Fonction auxiliaire qui définit le texte d'un nœud HTML.*/
func setTextContent(n *html.Node, content string) {
	n.FirstChild = &html.Node{Type: html.TextNode, Data: content}
	n.LastChild = n.FirstChild
}

func convertURLsInCSS(cssContent string, baseURL *url.URL) string {
	// Expression régulière pour trouver les URLs dans les déclarations url()
	re := regexp.MustCompile(`url\(['"]?([^'"]*?)['"]?\)`)

	// Remplacer les URLs dans le contenu CSS
	modifiedCSS := re.ReplaceAllStringFunc(cssContent, func(match string) string {
		// Extraire l'URL de la correspondance
		urlMatch := re.FindStringSubmatch(match)
		if len(urlMatch) < 2 {
			return match // Pas d'URL trouvée, retourner la correspondance
		}
		originalURL := urlMatch[1]

		// Résoudre l'URL par rapport à l'URL de base
		resolvedURL := baseURL.ResolveReference(&url.URL{Path: originalURL}).String()

		// Retourner l'URL modifiée entourée de la déclaration url()
		return "url('" + resolvedURL + "')"
	})

	return modifiedCSS
}

/*Convertit les URLs dans les déclarations url() dans le contenu HTML.*/
func ConvertURLs(htmlContent []byte) string {
	// Convertir le contenu HTML en chaîne de caractères
	htmlStr := string(htmlContent)

	// Chemin absolu vers le répertoire
	// Définir un motif regex pour trouver les valeurs URL à l'intérieur des fonctions url()
	urlPattern := `url\(['"]?(.*?)['"]?\)`

	// Compiler le motif regex
	re := regexp.MustCompile(urlPattern)

	// Remplacer les valeurs URL correspondantes par la valeur souhaitée
	modifiedHTML := re.ReplaceAllStringFunc(htmlStr, func(match string) string {
		//  Extraire la valeur URL de la chaîne correspondante
		url := re.FindStringSubmatch(match)[1]

		// Remplacer la valeur URL par le chemin absolu souhaité
		absolutePath := "/corndog.io" + url // Remplacer par le chemin du répertoire de travail
		return "url('" + absolutePath + "')"
	})

	return modifiedHTML
}
