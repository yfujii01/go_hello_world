package src

func Hello() string {
	hwStr := JoinMessage("Hello", "World")
	return hwStr
}

func JoinMessage(s1 string, s2 string) string {
	var hwStr string = s1 + " " + s2 + "!!"
	return hwStr
}

