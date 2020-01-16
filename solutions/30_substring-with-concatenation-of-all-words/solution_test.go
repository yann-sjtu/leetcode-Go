package _0_substring_with_concatenation_of_all_words

import (
	"fmt"
	"testing"
)

func findSubstring(s string, words []string) []int {
	sLen, wordsNum := len(s), len(words)
	if sLen == 0 || wordsNum == 0 {
		return nil
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	wordLen := len(words[0])
	var res []int
	type Stat struct {
		num   int
		index int
	}
	for n := 0; n < wordLen; n++ {
		for i := n; i <= sLen-wordLen*wordsNum; i += wordLen {
			stat := make(map[string]*Stat)
			var j int
			for j = i; j < i+wordLen*wordsNum; j += wordLen {
				word := s[j : j+wordLen]
				if _, ok := wordMap[word]; !ok {
					i = j
					break
				}
				if _, ok := stat[word]; !ok {
					stat[word] = &Stat{1, j}
				} else {
					stat[word].num++
					if stat[word].num > wordMap[word] {
						i = stat[word].index
						break
					}
				}
			}
			if j == i+wordLen*wordsNum {
				res = append(res, i)
			}
		}
	}
	return res
}

func findSubstringV2(s string, words []string) []int {
	sLen, wordsNum := len(s), len(words)
	if sLen == 0 || wordsNum == 0 {
		return nil
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	wordLen := len(words[0])
	var res []int
	for n := 0; n < wordLen; n++ {
		stat := make(map[string]int)
		var count int
		for start := n; start <= sLen-wordLen*wordsNum; start += wordLen {
			for index := start; index < start+wordLen*wordsNum && index <= sLen - wordLen ; index += wordLen {
				word := s[index : index+wordLen]
				if _, ok := wordMap[word]; !ok {
					start = index
					stat = make(map[string]int)
					count = 0
					break
				}
				stat[word]++
				count++
				if stat[word] > wordMap[word] {
					//start remove
					var removeCount int
					for k := start;;k+=wordLen{
						stat[s[k:k+wordLen]]--
						removeCount++
						if s[k:k+wordLen] == word {
							break
						}
					}
					count -= removeCount
					start += removeCount*wordLen
					continue
				}

				if count == wordsNum {
					res = append(res, start)
					stat = make(map[string]int)
					count = 0
					break
				}
			}
		}
	}
	return res
}

func Test_(t *testing.T) {
	s := "ababaab"
	words := []string{"ab", "ba", "ba"}
	fmt.Println(findSubstringV2(s, words))
	s2 := "wordgoodgoodgoodbestword"
	words2 := []string{"word","good","best","good"}
	fmt.Println(findSubstringV2(s2, words2))
}
