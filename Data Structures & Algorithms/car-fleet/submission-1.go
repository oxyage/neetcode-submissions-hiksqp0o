import "slices"

type CarFleet struct {
	position int
	speed int
}

func carFleet(target int, positions []int, speeds []int) int {
	var fleet = make(map[int]CarFleet, len(positions))

	for i, p := range positions {
		fleet[p] = CarFleet{
			position: p,
			speed: speeds[i],
		}
	}

	slices.Sort(positions)
	slices.Reverse(positions)

	fmt.Printf("positions struct: %v\n", fleet)

	var stack = make([]float64, 0)

	for _, pos := range positions {

		t := float64(target - pos) / float64(fleet[pos].speed)
		stack = append(stack, t)

		last := len(stack) - 1
		preLast := len(stack) - 2

		if len(stack) >= 2 && stack[ last ] <= stack[ preLast ] {
			stack = stack[ : last]
		}
	}


	return len(stack)
}
