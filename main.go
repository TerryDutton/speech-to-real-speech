package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var stutters []string = []string{"well", "you know", "uh", "um", "like", "I mean", "yeah, like", "yeah", "I dunno"}
var speech string = "Ladies and Gentlemen, it is a pleasure for me to be here tonight and address such a great audience. The issue I would like to bring up threatens the prosperity and welfare of the whole nation; however, the majority of the population tends to ignore it and pretend as if it is not a problem at all. Namely, I would like to talk about the risks of obesity."

func main() {
	realSpeech := speechToRealSpeech(speech)
	fmt.Println("Old speech:\n" + speech + "\n\n")
	fmt.Println("REAL speech:\n" + realSpeech)
}

func speechToRealSpeech(speech string) string {
	splitSpeech := strings.Split(speech, " ")
	newSpeech := seedSplitSpeechWithRandomStutters(splitSpeech)
	return strings.Join(newSpeech, " ")
}

func seedSplitSpeechWithRandomStutters(splitSpeech []string) (newSpeech []string) {
	n := len(splitSpeech)
	r := createRandomNumberGenerator()
	for _, word := range splitSpeech[0 : n-1] {
		newSpeech = append(newSpeech, possiblyAddStutter(word, r))
	}
	newSpeech = append(newSpeech, splitSpeech[n-1])
	return
}

func possiblyAddStutter(word string, r *rand.Rand) string {
	isAddStutter := rollPercentageChance(20, r)
	if !isAddStutter {
		return word
	}
	stutter := getRandomStutter(r)

	sentencePauses := []string{",", ";", ":", "..."}
	if thisWordEndsWithOneOf(sentencePauses, word) {
		return fmt.Sprintf("%s %s,%s", word, stutter, possiblyRepeatWord(word, r))
	}

	sentenceEnders := []string{".", "?", "!"}
	if thisWordEndsWithOneOf(sentenceEnders, word) {
		return fmt.Sprintf("%s %s,", word, capitalizeFirstLetter(stutter))
	}

	return fmt.Sprintf("%s, %s,%s", word, stutter, possiblyRepeatWord(word, r))
}

func createRandomNumberGenerator() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	return rand.New(source)
}

func rollPercentageChance(percent int, r *rand.Rand) bool {
	roll := r.Intn(100) + 1
	return percent >= roll
}

func getRandomStutter(r *rand.Rand) string {
	i := r.Intn(len(stutters))
	return stutters[i]
}

func thisWordEndsWithOneOf(suffixes []string, word string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			return true
		}
	}
	return false
}

func possiblyRepeatWord(word string, r *rand.Rand) (repeatedWord string) {
	wordShouldBeRepeated := rollPercentageChance(25, r)
	if wordShouldBeRepeated {
		repeatedWord = " " + word
	}
	return
}

func capitalizeFirstLetter(word string) string {
	firstLetter := string(word[0])
	restOfWord := string(word[1:])
	return strings.ToUpper(firstLetter) + restOfWord
}
