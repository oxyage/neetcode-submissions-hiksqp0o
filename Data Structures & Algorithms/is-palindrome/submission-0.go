func isPalindrome(s string) bool {

	// подготовить входные данные:
	// очистить от знаков препинания, привести к нижнему регистру, удалить пробелы
	
	// из строки "Pal il ap?"
	// получить palilap
	
	// 1. вариант
	// если четное число символов, просто разделить пополам
	// если нечетное, разделить по границе  mid = floor(len / 2): s[0:mid] "_" s(mid:] 
	// цикл для l = 0; до l / 2; l++ - идем слева
	// второй цикл идем справа r = len(); до r / 2; r--
	s = keepAlphabet(strings.ToLower(s))
	var s_len = len(s)
	var s_len_split = int(len(s) / 2)
	var s_char = []byte(s)

	
	
	for l := 0; l < s_len_split; l += 1 {
		var r = s_len - 1 - l

		if s_char[l] != s_char[r] {
			return false
		}
	}

	return true
}

func keepAlphabet(s string) string {

	rg := regexp.MustCompile(`[^a-z0-9]+`)
	return rg.ReplaceAllString(s, "")
}