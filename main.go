package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"time"

	game "github.com/ShinyaIshitobi/typing-game/game"
)

func main() {
	file, err := os.Open("game/words.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	words2D, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var words []string
	for _, word := range words2D {
		words = append(words, word[0])
	}

	limit := flag.Int("limit", 20, "time limit")
	flag.Parse()

	timeLimit := time.Duration(*limit) * time.Second
	_, err = game.Game(os.Stdin, os.Stdout, words, timeLimit)
	if err != nil {
		log.Fatal(err)
	}
}
