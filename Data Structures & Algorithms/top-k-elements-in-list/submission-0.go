func topKFrequent(nums []int, k int) []int {

	// как решим задачу? 
	// вернуть срез из k часто встречающихся элементов среза nums

	// - чтобы понять какие элементы часто встречаются, посчитаем их, составим гистограмму частот
	// составить гистограмму, где ключ = число, значение = счетчик
	// или
	// ключ = счётчик, значение = срез элементов, которые встречаются столько раз
	// фиксировать самый большой счетчик - поможет на следующем шаге

	// пройдем по срезу, составим гистограмму, проверяем счетчик maxFrequency - переписываем, если преодолели его
	// maxFrequency перейдет на следующий шаг

	// [ 1, 73, 73, 73, 34, 23, 34, 73 ]
	// histogram: {1 => 1; 73 => 4; 34 => 2; 23 => 1}

	// 
	// sorted histogram => []struct{key => Number, value => frequency}

	// 1 шагом, составим гистограмму обычную
	// параллельно ведем счетчик самого часто встречающегося числа
	var hist = make(map[int]int, len(nums))
	var most_freq = 0
	for _, n := range nums {
		if _, ok := hist[n]; ok {
			hist[n] += 1
		} else {
			hist[n] = 1
		}

		// если текущая выше, обновим 
		if hist[n] > most_freq {
			most_freq = hist[n]
		}
	}

	// 2 шагом составим новый срез структур и отсортируем
	type Hist struct {
		Value int
		Frequency int
	}

	var sorted_by_freq_desc = make([]Hist, len(nums))
	for val, freq := range hist {
		sorted_by_freq_desc = append(sorted_by_freq_desc, Hist{Value: val, Frequency: freq})
	}

	// desc sort
	sort.Slice(sorted_by_freq_desc, func (l, r int) bool {
		return sorted_by_freq_desc[l].Frequency > sorted_by_freq_desc[r].Frequency
	})

	// 3 шагом выбираем самые часто используемые
	var result = make([]int, 0, k)
	for _, s := range sorted_by_freq_desc {
		if len(result) >= k { 
			break
		}

		result = append(result, s.Value)
	}

	return result
}
