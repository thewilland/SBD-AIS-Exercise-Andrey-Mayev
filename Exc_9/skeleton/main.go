package main

import (
	"exc9/mapred"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Main function
func main() {
	// todo read file
	data, err := os.ReadFile("res/meditations.txt")
	if err != nil {
		panic(err)
	}

	text := strings.Split(string(data), "\n")

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)

	// todo print your result to stdout

	// Alphabetical version
	// // extracting keys
	// keys := make([]string, 0, len(results))
	// for k := range results {
	// 	keys = append(keys, k)
	// }

	// // sorting keys alphabetically
	// sort.Strings(keys)

	// // printing in order
	// for _, k := range keys {
	// 	fmt.Printf("%s: %d\n", k, results[k])
	// }

	// Frequency version
	// map → slice
	pairs := make([]mapred.KeyValue, 0, len(results))
	for word, count := range results {
		pairs = append(pairs, mapred.KeyValue{Key: word, Value: count})
	}

	// sorting by frequency (descending)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	// top words
	topN := 100

	// making sure we don't go out of bounds
	if len(pairs) < topN {
		topN = len(pairs)
	}

	// printing top N
	for i := 0; i < topN; i++ {
		fmt.Printf("%3d. %-15s %d\n", i+1, pairs[i].Key, pairs[i].Value)
	}
}
