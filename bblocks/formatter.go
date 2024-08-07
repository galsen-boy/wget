package bblocks

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*Formater une taille de fichier en utilisant des unités lisibles telles que GB, MB, ou KB.*/
func FormatSize(size int64) string {
	if size > 1e9 {
		return fmt.Sprintf("%dGB", size/(1024*1024*1024))
	} else if size > 1e6 {
		return fmt.Sprintf("%dMB", size/(1024*1024))
	} else {
		return fmt.Sprintf("%dKB", size/1024)
	}
}

// Créer une nouvelle instance de RateLimiter avec une limite de vitesse spécifiée pour les téléchargements.
func NewLimiter(downloadSpeed int) *RateLimiter {
	return &RateLimiter{
		limit:  float64(downloadSpeed),
		burst:  1024 * 1024, // Allow bursts up to 1 MB
		tokens: 0,
	}
}

// Limit returns the allowed download speed in bytes per second.
func (r *RateLimiter) Limit() float64 {
	return r.limit
}

// Retourner la vitesse limite actuelle du téléchargeur.
// Reserve calculates the time required to download a given number of bytes.
func (r *RateLimiter) Reserve(bytes int) time.Duration {
	requiredTokens := float64(bytes)
	timeRequired := time.Duration(requiredTokens / r.limit * float64(time.Second))
	return timeRequired
}

/*
Parser une chaîne de caractères représentant une limite de taux
(ex. "500K" ou "2M") et retourner la valeur en octets.
*/
func ParseRateLimit(rateLimit string) (int, error) {
	numericPart := rateLimit[:len(rateLimit)-1]
	suffix := rateLimit[len(rateLimit)-1]

	value, err := strconv.ParseFloat(numericPart, 64)
	if err != nil {
		return 0, err
	}

	switch strings.ToUpper(string(suffix)) {
	case "K":
		value *= 1024
	case "M":
		value *= 1024 * 1024
	}

	return int(value), nil
}
