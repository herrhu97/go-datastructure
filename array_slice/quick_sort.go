package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 0)

	for i := 0; i < 10; i++ {
		arr = append(arr, int(rand.Int63n(100)))
	}

	fmt.Println(arr)
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quick_sort(q []int, l, r int) {
	if l >= r {
		return
	}

	i, j, x := l-1, r+1, q[(l+r)>>1]

	for i < j {
		i++
		for q[i] < x {
			i++
		}

		j--
		for q[j] > x {
			j--
		}

		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quick_sort(q, l, j)
	quick_sort(q, j+1, r)
}
