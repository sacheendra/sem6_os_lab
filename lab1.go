package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(simple_queue(0.5, 0.3, 0, 10, 10000))
}

func simple_queue(arrival_rate float64, service_rate float64, balk_len float64, pre_count float64, iterations int) float64 {
	N := iterations
	var prev_q_len float64 = 0
	q_len_log := make([]float64, N)
	var no_depart_count float64 = 0

	for k := 0; k < N; k++ {
		var arrive float64
		if balk_len != 0 && prev_q_len >= balk_len {
			arrive = 0
		} else {
			arrive = eprob(arrival_rate)
		}
		depart := eprob(service_rate)
		if pre_count != 0 && no_depart_count > pre_count {
			depart = 1
		}
		if depart == 0 {
			no_depart_count++
		} else {
			no_depart_count = 0
		}
		next_q_len := math.Max(0, prev_q_len+arrive-depart)
		
		q_len_log[k] = next_q_len
		prev_q_len = next_q_len
	}
	mean_len := average(q_len_log)

	return mean_len
}

func eprob(x float64) float64 {
	t := rand.Float64()
	if t <= x {
		return 1
	} else {
		return 0
	}
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
