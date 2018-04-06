// Stefan Nilsson 2013-02-27

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
	"time"
)

type ComplexFunc func(complex128) complex128

var limit chan struct{}

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z * z },
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
	before := time.Now()
	goroutines := 10                        // The number used to limit gogoutines
	pixels := 1024                          // The image will be pixels * pixels in size
	limit = make(chan struct{}, goroutines) // Antal tillåtna samtidiga goroutines
	imgwg := new(sync.WaitGroup)            // Waitgroup for letting all images get finished before ending main
	for n, fn := range Funcs {
		imgwg.Add(1)                                                       // Add one more picture to the image waitgroup
		go CreatePng("picture-"+strconv.Itoa(n)+".png", fn, pixels, imgwg) // Runs a goroutine for a new image
	}
	imgwg.Wait() // Wait until all images are done
	fmt.Println("time:", time.Now().Sub(before))
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int, imgwg *sync.WaitGroup) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		imgwg.Done()
		return
	}
	defer file.Close()
	err = png.Encode(file, Julia(f, n))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done with image", filename)
	imgwg.Done() // Says a new image is done
}

// Julia returns an image of size n x n of the Julia set for f.
func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	wg := new(sync.WaitGroup) // Waitgroup to make sure all pixels are done before returning image
	wg.Add(n * n)             // Adding ammount of required pixels
	img := image.NewRGBA(bounds)
	s := float64(n / 4)
	r := uint8(0) // Setting red color to zero
	g := uint8(0) // Setting green color to zero
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			select {
			case limit <- struct{}{}: // If there are below a certain amount of goroutines
				
				go func(i, j int) {
					n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
					b := uint8(n % 32 * 8)
					img.Set(i, j, color.RGBA{r, g, b, 255})
					<-limit   // Reduce "active goroutine counter" by one
					wg.Done() // Lets waitgroup know a new pixel is done
				}(i, j)

			default: // In case of the maximum number of goroutines are active
				n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
				b := uint8(n % 32 * 8)
				img.Set(i, j, color.RGBA{r, g, b, 255})
				wg.Done() // Lets waitgroup know a new pixel is done
			}
		}
	}
	wg.Wait() // Wait for all pixels before returning image
	return img
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}

// What has been done?
// There is a new goroutine started for each Image
// There are new goroutines launched in Julia for the Iterate function (with a goroutine limit since this will reach very high numbers if the image has a high number of pixels)
// The limit on goroutines are built using buffered channels

// The result went from around 17-18 seconds to around 5-6 seconds.
// Only creating goroutines for each Iterate launch brought the time to 7-8 sec, but adding goroutines for each image brought it to 5-6 sec.

// runtime.GOMAXPROCS() i main för att ställa in godkänt antal trådar
// This computer has 2 cores but 4 logical processors
