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
	// TODO: sorted histogram
	// bucket_hist: {1 => [1, 23]; 2 => [34]; 4 => [73]}

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

	// 2 шагом составим гистограмму бакетов
	var bucket_hist = make(map[int][]int)
	for val, freq := range hist {
		if _, ok := bucket_hist[freq]; ok {
			bucket_hist[freq] = append(bucket_hist[freq], val)
		} else {
			bucket_hist[freq] = []int{val}
		}
	}
	// bucket_hist: {1 => [1, 23]; 2 => [34]; 4 => [73]}


	// 3 шагом выбираем самые часто используемые
	var result = make([]int, 0, k)
	for i := most_freq; i >= 1 && len(result) < k; i-- {

		if _, ok := bucket_hist[i]; ok {
			result = append(result, bucket_hist[i]...)
		} 		
	}

	return result
}
