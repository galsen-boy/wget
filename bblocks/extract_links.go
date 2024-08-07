package bblocks

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/temoto/robotstxt"
	"golang.org/x/net/html"
)

// userAgent : Chaîne représentant l'agent utilisateur pour les requêtes HTTP.
var userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"

/*
La fonction Crawl permet de parcourir les pages web d'un domaine donné,

	tout en respectant les règles définies dans le fichier robots.txt,
	 et en extrayant et suivant les liens présents sur les pages.
*/
func Crawl(urlw string, baseURL *url.URL, discovered map[string]bool, client *http.Client, robots *robotstxt.RobotsData) {
	if _, ok := discovered[urlw]; ok {
		return
	}

	//Analyse de l'URL
	parsedURL, err := url.Parse(urlw)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Vérification du domaine
	if parsedURL.Host != baseURL.Host {
		fmt.Println("Skipping external domain:", urlw)
		return
	}
	//Création et envoi de la requête HTTP
	req, err := http.NewRequest("GET", urlw, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("User-Agent", userAgent)

	Resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer Resp.Body.Close()

	if Resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", Resp.Status)
		return
	}
	//Vérification des règles robots.txt :
	if !robots.TestAgent(urlw, userAgent) {
		fmt.Println("Robot not allowed to crawl:", urlw)
		return
	}
	//Marquage de l'URL comme découverte et affichage
	discovered[urlw] = true
	fmt.Println(urlw)
	//Analyse et extraction des liens HTML
	tokenizer := html.NewTokenizer(Resp.Body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "a", "link", "script", "img":
				for _, attr := range token.Attr {
					if (token.Data == "a" && attr.Key == "href") || (token.Data == "link" && attr.Key == "href") || (token.Data == "script" && attr.Key == "src") || (token.Data == "img" && attr.Key == "src") {
						link := attr.Val
						if strings.HasPrefix(link, "http") {
							Crawl(link, baseURL, discovered, client, robots)
						} else {
							// Resolve relative paths
							linkURL, err := baseURL.Parse(link)
							if err != nil {
								fmt.Println("Error resolving URL:", err)
								continue
							}
							Crawl(linkURL.String(), baseURL, discovered, client, robots)
						}
					}
				}
			case "style":
				for {
					tokenType := tokenizer.Next()
					if tokenType == html.ErrorToken || tokenType == html.EndTagToken && tokenizer.Token().Data == "style" {
						break
					} else if tokenType == html.TextToken {
						cssContent := tokenizer.Token().Data
						cssURLs := ExtractURLsFromCSS(cssContent, baseURL)
						for _, cssURL := range cssURLs {
							Crawl(cssURL, baseURL, discovered, client, robots)
						}
					}
				}
			}
		}
	}
}

/*
Cette fonction extrait les URLs des déclarations url() dans le contenu CSS
et les résout par rapport à l'URL de base si nécessaire.
*/
func ExtractURLsFromCSS(cssContent string, baseURL *url.URL) []string {
	var urls []string

	// Regular expression to match URLs within url() declarations
	re := regexp.MustCompile(`url\(['"]?([^'"]*?)['"]?\)`)

	// Find all matches in the CSS content
	matches := re.FindAllStringSubmatch(cssContent, -1)
	for _, match := range matches {
		url := match[1] // The URL is captured in the second group
		if strings.HasPrefix(url, "(") {
			// Absolute URL
			urls = append(urls, url)
			fmt.Println(urls)
		} else {
			// Relative URL, resolve it relative to the base URL
			linkURL, err := baseURL.Parse(url)
			if err != nil {
				fmt.Println("Error resolving URL:", err)
				continue
			}
			urls = append(urls, linkURL.String())
		}
	}
	return urls
}
