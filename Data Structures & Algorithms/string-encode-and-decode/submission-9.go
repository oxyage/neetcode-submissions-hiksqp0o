type Solution struct{
	Delimiters []int
	Result string
}

// Hello, my, love, and, World
// 5,5,2,4,3,5#HellomyloveandWorld

const delimiter = "#*#*#*#"
const num_delimiter = ","

func (s *Solution) Encode(strs []string) string {
	
	var result string
	var delimiterResult string
	for _, str := range strs {

		if len(delimiterResult) > 0 {
			delimiterResult = delimiterResult + fmt.Sprintf(",%d",len(str))
		} else {
			delimiterResult = fmt.Sprintf("%d",len(str))
		}

		s.Result = s.Result + str
	}

	result = delimiterResult + delimiter + s.Result

	return result
}

func (s *Solution) Decode(encoded string) []string {
	splitted := strings.Split(encoded, delimiter) // 5,5,2,4,3,5	
	delimiters := strings.Split(splitted[0], num_delimiter) //[]string

	for _, v := range delimiters {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
	
		s.Delimiters = append(s.Delimiters, num)
	}


	payload := splitted[1]
	var result = make([]string, 0, len(s.Delimiters))
	var lastPos = 0
	for _, l := range s.Delimiters {
	
		var word string
		for i := lastPos; i < lastPos + l; i ++ {
			word += string(payload[i])
		}

		lastPos += l
		result = append(result, word)
	}

	return result
}
