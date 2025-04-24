package main

import "fmt"

func main() {
	var image_width = 256
	var image_height = 256

  fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)

	for j := range image_height {
		for i := range image_width {
      var r = float64(i) / float64(image_width - 1)
      var g = float64(j) / float64(image_height - 1)
      var b = float64(0)

      var ir = int(255.999 * float64(r))
      var ig = int(255.999 * float64(g))
      var ib = int(255.999 * float64(b))

      fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
