package utils

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("Test fetch", func(t *testing.T) {
		resp, _ := GetRandomQuote()
		log.Println(resp)
	})
}
