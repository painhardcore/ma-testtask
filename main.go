package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/painhardcore/MillionAgents/internal/httpsearcher"
)

type searcher interface {
	Search(string, string)
	GetTotal() int64
	Wait()
}

func main() {
	// cmd flags
	k := flag.Int64("k", 5, "Concurrent jobs")
	s := flag.String("word", "Go", "String to search count for")
	flag.Parse()

	httpSearcher := httpsearcher.NewWithLimit(*k)

	total := searchAndCount(httpSearcher, *s, os.Stdin)
	fmt.Printf("Total: %d\n", total)

}

func searchAndCount(searcher searcher, word string, file *os.File) int64 {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "quit" {
			break
		}
		searcher.Search(scanner.Text(), word)
	}
	searcher.Wait()
	return searcher.GetTotal()
}
