package piscine

func ConcatParams(args []string) string {
	//args  count 3
	//[0HOw 1are 2 you?]
	count := 0
	str := "" // How are you? friend
	nl := "\n"
	//4
	for range args {
		count++
	}

	for i := 0; i < count; i++ { //0<3, 1<3,2<3,
		str = str + args[i] // "" = ""+ How| How = How + are + you?
		if count-1 > i {    // 3>0, newlnine,  3>1, newline 3>2, 3>3 false
			str = str + nl
		}
	}
	return str // how are you? friend
}
