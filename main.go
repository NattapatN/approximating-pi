package main

import (
	"fmt"
	"math"
	"math/rand"
)

var total int
var circle int
var r int = 1000

func main() {
	fmt.Println("Approximating the value PI")
	var pi float64
	round := 0
	for {
		for i := 0; i < 100000; i++ {
			round++
			x := -1*float64(r) + rand.Float64()*(2*float64(r))
			y := -1*float64(r) + rand.Float64()*(2*float64(r))
			total++
			distance := math.Sqrt(x*x + y*y)
			if distance < float64(r) {
				circle++
			}
		}
		pi = 4 * (float64(circle) / float64(total))
		if math.Abs(pi-math.Pi) < 0.000001 {
			break
		}
	}
	fmt.Println("Trying :", round, "roudns.")
	fmt.Println("PI value is :", math.Pi)
	fmt.Println("Approximating the value PI is", pi)
	fmt.Println("Abslute different between 'PI' and 'Approximatingthe value PI' :", math.Abs(pi-math.Pi))
}
