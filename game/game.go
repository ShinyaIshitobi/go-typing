package game

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func Game(r io.Reader, w io.Writer, words []string, t time.Duration) (int, error) {

	limit := time.After(t)
	rand.Seed(int64(time.Now().UnixNano()))

	_, err := fmt.Fprintf(w, "start typing game.\ntype the word")
	if err != nil {
		return -1, err
	}

	_, err = fmt.Fprintf(w, "time limit is %d\n", int(t.Seconds()))
	if err != nil {
		return -1, err
	}

	ch := makeCh(r)
	score := 0
	var word string

	for {
		idx := rand.Intn(len(words))
		word = words[idx]
		_, err := fmt.Fprintf(w, ">> %s    ", word)
		if err != nil {
			return -1, err
		}

		select {
		case ans := <-ch:
			score++
			if ans != word {
				score--
				fmt.Fprintf(w, "missed!\n")
			}
		case <-limit:
			_, err = fmt.Fprintf(w, "\n>> game ends! the number of correct answer: %d", score)
			if err != nil {
				return -1, err
			}
			return score, err
		}
	}
}

func makeCh(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}
