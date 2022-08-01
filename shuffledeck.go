package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	alphabet := []string{"spades_06", "spades_07", "spades_08", "spades_09", "spades_10", "spades_J", "spades_Q", "spades_K", "spades_A", "heart_02",
	 "heart_03", "heart_04", "heart_05", "heart_06", "heart_07", "heart_08", "heart_09", "heart_J", "heart_Q", "heart_K", "heart_A",
	  "Club_02", "Club_03", "Club_04", "Club_05", "Club_06", "Club_07", "Club_08", "Club_09", "Club_10", "Club_J", "Club_Q", "Club_K", "Club_A",
	   "Diamond_02", "Diamond_03", "spades_02", "spades_03", "spades_04", "spades_05", "Diamond_04", "Diamond_05", "Diamond_06", "Diamond_07", "Diamond_08", "Diamond_09", "Diamond_10", "Diamond_J", "Diamond_Q", "Diamond_K", "Diamond_A"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(alphabet), func(i, j int) {
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
	})
	//fmt.Println(alphabet)
	// sort.Slice(alphabet, func(i, j int) bool {
	// 	return alphabet[i] < alphabet[j]
	// })

	sort.Strings(alphabet)
	customsort(alphabet)
	}
func customsort(slc []string) {
	var slcaft = make([]string, 0)

	for i := 13; i < 26; i++ {
		slcaft = append(slcaft, slc[i])
	}

	for i := 26; i < 38; i++ {
		slcaft = append(slcaft, slc[i])
	}

	
	for i := 38; i < 51; i++ {
		slcaft = append(slcaft, slc[i])
	}
	for i := 0; i < 12; i++ {
		slcaft = append(slcaft, slc[i])

	}

	fmt.Println(slcaft)

}