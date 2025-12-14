package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
func (mr MapReduce) wordCountMapper(text string) []KeyValue {
	re := regexp.MustCompile(`[^a-zA-Z]+`)
	clean := re.ReplaceAllString(strings.ToLower(text), " ")

	words := strings.Fields(clean)
	var result []KeyValue

	for _, word := range words {
		result = append(result, KeyValue{Key: word, Value: 1})
	}

	return result
}

func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return KeyValue{Key: key, Value: sum}
}

func (mr MapReduce) Run(input []string) map[string]int {
	var wg sync.WaitGroup
	mapChannel := make(chan []KeyValue)

	// Mapping
	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			mapChannel <- mr.wordCountMapper(text)
		}(line)
	}

	// Closing channel
	go func() {
		wg.Wait()
		close(mapChannel)
	}()

	// Shuffling
	grouped := make(map[string][]int)
	for kvs := range mapChannel {
		for _, kv := range kvs {
			grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
		}
	}

	// Reducing
	results := make(map[string]int)
	for key, values := range grouped {
		kv := mr.wordCountReducer(key, values)
		results[kv.Key] = kv.Value
	}

	return results
}
