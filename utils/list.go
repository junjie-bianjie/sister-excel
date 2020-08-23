package utils

func Different(args1, args2 []string) []string {
	var res []string
	intersection := Intersection(args1, args2)

	tempMap := make(map[string]int, len(args1))
	for _, i := range intersection {
		tempMap[i]++
	}

	for _, s1 := range args1 {
		if _, ok := tempMap[s1]; !ok {
			res = append(res, s1+"\t#1")
		}
	}

	for _, s2 := range args2 {
		if _, ok := tempMap[s2]; !ok {
			res = append(res, s2+"\t#2")
		}
	}
	return res
}

func Intersection(args1, args2 []string) []string {
	tempMap := make(map[string]int, len(args2))
	for _, s1 := range args1 {
		tempMap[s1]++
	}

	var res []string
	for _, s2 := range args2 {
		if _, ok := tempMap[s2]; ok {
			res = append(res, s2)
		}
	}
	return res
}
