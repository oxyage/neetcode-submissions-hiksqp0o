import "slices"
func groupAnagrams(strs []string) [][]string {
	// todo: задачу можно решить через сортировку значений


	/*
	 чтобы построить конечный массив, нам нужна хеш таблица
	 она позволит строить группы внутри конеченого массива
	 что есть ключ в этой хеш таблице? = string (сортированный)

	map[string][]string {

// ключ должен быть сортированный 	
	"a1h1t1": ["hat"],
	"a1c1t1": ["act", "cat"],
	"s1t1o1p1": ["stop", "pots", "tops"],
	"x": ["x"]
	"": [""]
	
	}	

	*/



	var hist = make(map[string][]string)
	for _, word := range strs {

		
		histKeyRune := []rune(word)

		slices.Sort(histKeyRune)

		histKey := string(histKeyRune)

		if _, ok := hist[histKey]; ok {
			hist[histKey] = append(hist[histKey], word)
		} else {
			hist[histKey] = []string{word}
		}

		

		
	}


	// конечный массив данных содержит группы массивов
	var result [][]string
	for _, v := range hist {
		if len (v) > 0 {
			result = append(result, v)
		}
	}


	return result
}
