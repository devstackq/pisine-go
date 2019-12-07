package piscine

func SplitWhiteSpaces(str string) []string {
	// count word, check -> each letter, and - when  will be - spaces - countword +1
	// new var for - concat each letter - in one word
	wordCount := 0

	var arrSentence []string
	letterConcat := "" // h+e+l+...
	//arrSentence = make(string(letterConcat), wordCount)
	arrRunes := []rune(str) // reunes from - strings
	//  split letter, after if - ' space' -> wordcount +1
	// count word, in string
	for _, v := range arrRunes {
		if v == ' ' {
			if letterConcat != "" {
				wordCount++
				letterConcat = ""
			}
		} else {
			letterConcat = letterConcat + string(v)
		}

	}
	// wordCount +1, and nil value letterConcat
	if letterConcat != "" {
		wordCount++
		letterConcat = ""
		// fmt.Println(wordCount)
	}

	//array string - create - size,
	arrSentence = make([]string, wordCount) // 4 words
	//
	wordCount2 := 0
	for _, v := range arrRunes {
		if v == ' ' {
			if letterConcat != "" {
				//fmt.Println(wordCount2)
				wordCount2++
				arrSentence[wordCount2-1] = letterConcat // 1-1 [0: Hello, 1: How, 2: Are] write 0 index, word - inside arrSentece
				letterConcat = ""
			}
		} else {
			letterConcat = letterConcat + string(v)
		}
	}
	if letterConcat != "" {
		wordCount2++
		arrSentence[wordCount2-1] = letterConcat // 4-1 = 3 [3: you?]
		letterConcat = ""
	}

	return arrSentence
}

// if letter -
// count + 1
// if space, count old,
