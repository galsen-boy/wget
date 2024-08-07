package bblocks

import (
	"io"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

/*
La fonction DownloadWithStandardProgressBar permet de télécharger des données tout
en affichant une barre de progression standard et en respectant une éventuelle limite de débit.
*/
func DownloadWithStandardProgressBar(body io.Reader, file *os.File, limiter *RateLimiter, totalSize int64, bar *progressbar.ProgressBar) error {
	var reader io.Reader = body // Utilise le lecteur `body` initialement

	//  Applique le limiteur de débit si fourni
	if limiter != nil {
		reader = NewRateLimitedReader(body, limiter) // Enveloppe le lecteur `body` avec le limiteur de débit
	}

	// Copie les données du lecteur vers le fichier avec la barre de progression
	_, err := io.Copy(io.MultiWriter(file, bar), reader)
	if err != nil {
		return err
	}

	return nil
}

/*Cette fonction crée un nouveau lecteur enveloppé (rateLimitedReader) qui applique une limite de débit.*/
func NewRateLimitedReader(r io.Reader, limiter *RateLimiter) io.Reader {
	return &rateLimitedReader{
		reader:  r,
		limiter: limiter,
	}
}

// Un type qui contient un lecteur et un limiteur de débit.
// rateLimitedReader is a rate-limited reader that enforces the rate limit while reading.
type rateLimitedReader struct {
	reader  io.Reader
	limiter *RateLimiter
}

// Read reads data from the underlying reader and enforces the rate limit.
func (rlr *rateLimitedReader) Read(p []byte) (n int, err error) {
	// Réserve des jetons basés sur la taille du tampon
	n, err = rlr.reader.Read(p)
	if err != nil {
		return n, err
	}
	if rlr.limiter != nil {
		//Calcule le temps nécessaire pour lire les données et réserve des jetons
		timeRequired := rlr.limiter.Reserve(n)
		time.Sleep(timeRequired)
	}
	return n, nil
}
