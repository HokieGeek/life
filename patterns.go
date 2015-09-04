package life

import (
	// "fmt"
	"math/rand"
	"time"
)

/////////////////////////// COMMON ///////////////////////////

func getCountsForDimensions(dimensions Dimensions, width int, height int) (int, int) {
	numPerRow := dimensions.Width / height
	numPerCol := dimensions.Height / width

	// Special case for when the spacer is not needed
	if numPerRow == 0 && dimensions.Height == height-1 {
		numPerRow = 1
	}
	if numPerCol == 0 && dimensions.Width == width-1 {
		numPerCol = 1
	}

	// fmt.Printf(">>>> numPerRow = %d\n", numPerRow)
	// fmt.Printf(">>>> numPerCol = %d\n", numPerCol)
	return numPerRow, numPerCol
}

// , startingLocation Location,
func getRepeatingPattern(dimensions Dimensions, height int, width int,
	pattern func(*[]Location, int, int)) []Location {

	numPerRow, numPerCol := getCountsForDimensions(dimensions, width, height)

	seed := make([]Location, 0)
	for row := 0; row < numPerCol; row++ {
		currentY := (row * height)
		// currentY := (row * height) + startingLocation.Y

		for col := 0; col < numPerRow; col++ {
			currentX := (col * width)
			// currentX := (col * width) + startingLocation.X
			pattern(&seed, currentX, currentY)
		}
	}

	return seed
}

/////////////////// RANDOM ///////////////////

func Random(dimensions Dimensions, percent int) []Location {
	seed := make([]Location, 0)

	for i := 0; i < dimensions.Height; i++ {
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < dimensions.Width; j++ {
			if rand.Intn(100) > percent {
				seed = append(seed, Location{X: i, Y: j})
			}
		}
	}

	return seed
}

/////////////////// OSCILLATORS ///////////////////

// func Blinkers(dimensions Dimensions, startingLocation Location) []Location {
func Blinkers(dimensions Dimensions) []Location {
	// put in as many lengthx1 vertical lines as you can fit
	// Period 1   Period 2
	// -0-		  ---
	// -0-        000
	// -0-        ---

	const HEIGHT = 4 // 3 for the line itself and 1 for the spacer
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			for i := 0; i < 3; i++ {
				*seed = append(*seed, Location{X: currentX + i, Y: currentY + 1})
			}
		})

}

func Toads(dimensions Dimensions) []Location {
	// Period 1	  Period 2
	// ----       --0-
	// -000       0--0
	// 000-       0--0
	// ----       -0--

	const HEIGHT = 5
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			// ROW 1
			for i := 1; i < 4; i++ {
				*seed = append(*seed, Location{X: currentX + i, Y: currentY + 1})
			}
			// ROW 2
			for i := 0; i < 3; i++ {
				*seed = append(*seed, Location{X: currentX + i, Y: currentY + 2})
			}
		})
}

func Beacons(dimensions Dimensions) []Location {
	// Period 1   Period 2
	// 00--       00--
	// 0---       00--
	// ---0       --00
	// --00       --00

	const HEIGHT = 5
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			// ROW 1
			*seed = append(*seed, Location{X: currentX, Y: currentY})
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY})
			// ROW 2
			*seed = append(*seed, Location{X: currentX, Y: currentY + 1})
			// ROW 3
			*seed = append(*seed, Location{X: currentX + 3, Y: currentY + 2})
			// ROW 4
			*seed = append(*seed, Location{X: currentX + 2, Y: currentY + 3})
			*seed = append(*seed, Location{X: currentX + 3, Y: currentY + 3})
		})
}

func Pulsar(dimensions Dimensions) []Location {
	// Period 1          Period 2           Period 3
	//                                        012345678901234
	// ---------------   ----0-----0----    0 ---------------
	// ---000---000---   ----0-----0----    1 ---00-----00---
	// ---------------   ----00---00----    2 ----00---00----
	// -0----0-0----0-   ---------------    3 -0--0-0-0-0--0-
	// -0----0-0----0-   000--00-00--000    4 -000-00-00-000-
	// -0----0-0----0-   --0-0-0-0-0-0--    5 --0-0-0-0-0-0--
	// ---000---000---   ----00---00----    6 ---000---000---
	// ---------------   ---------------    7 ---------------
	// ---000---000---   ----00---00----    8 ---000---000---
	// -0----0-0----0-   --0-0-0-0-0-0--    9 --0-0-0-0-0-0--
	// -0----0-0----0-   000--00-00--000    0 -000-00-00-000-
	// -0----0-0----0-   ---------------    1 -0--0-0-0-0--0-
	// ---------------   ----00---00----    2 ----00---00----
	// ---000---000---   ----0-----0----    3 ---00-----00---
	// ---------------   ----0-----0----    4 ---------------

	const HEIGHT = 16
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			for i := 0; i < HEIGHT-1; i++ {
				switch i {
				case 1, 6, 8, 13:
					// Repeating pattern 1: --000---000--, rows: 0,5,7,12
					for j := 0; j < 2; j++ {
						for k := 3; k <= 5; k++ {
							*seed = append(*seed,
								Location{X: currentX + (k + (6 * j)), Y: currentY + i})
						}
					}
				case 3, 4, 5, 9, 10, 11:
					// Repeating pattern 2: 0----0-0----0, rows: 2,3,4,8,9,10
					for j := 0; j < 2; j++ {
						*seed = append(*seed, Location{X: currentX + (1 + (7 * j)), Y: currentY + i})
						*seed = append(*seed, Location{X: currentX + (6 + (7 * j)), Y: currentY + i})
					}
				}
			}
		})
}

/////////////////// GLIDERS ///////////////////

func Gliders(dimensions Dimensions) []Location {
	// Period 1   Period 2   Period 3   Period 4
	// -0--       ----       ----       ----
	// --0-       0-0-       --0-       -0--
	// 000-       -00-       0-0-       --00
	// ----       -0--       -00-       -00-

	const (
		HEIGHT = 3
		WIDTH  = 4
	)
	return getRepeatingPattern(dimensions, HEIGHT, WIDTH,
		func(seed *[]Location, currentX int, currentY int) {
			// ROW 1
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY})
			// ROW 2
			*seed = append(*seed, Location{X: currentX + 2, Y: currentY + 1})
			// ROW 3
			for i := 0; i < 3; i++ {
				*seed = append(*seed, Location{X: currentX + i, Y: currentY + 2})
			}
		})
}

/////////////////// STILLS ///////////////////

func Blocks(dimensions Dimensions) []Location {
	// 00
	// 00

	const HEIGHT = 5
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			// ROW 1
			*seed = append(*seed, Location{X: currentX, Y: currentY})
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY})
			// ROW 2
			*seed = append(*seed, Location{X: currentX, Y: currentY + 1})
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY + 1})
		})
}

// Beehive
func Beehive(dimensions Dimensions) []Location {
	// -00-
	// 0--0
	// -00-

	const (
		HEIGHT = 3
		WIDTH  = 4
	)
	return getRepeatingPattern(dimensions, HEIGHT, WIDTH,
		func(seed *[]Location, currentX int, currentY int) {
			for row := 0; row < 3; row++ {
				switch row {
				case 0, 2:
					*seed = append(*seed, Location{X: currentX + 1, Y: currentY + row})
					*seed = append(*seed, Location{X: currentX + 2, Y: currentY + row})
				case 1:
					*seed = append(*seed, Location{X: currentX, Y: currentY + row})
					*seed = append(*seed, Location{X: currentX + 3, Y: currentY + row})
				}
			}
		})
}

func Loaf(dimensions Dimensions) []Location {
	// -00-
	// 0--0
	// -0-0
	// --0-

	const HEIGHT = 4
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			// ROW 1
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY})
			*seed = append(*seed, Location{X: currentX + 2, Y: currentY})
			// ROW 2
			*seed = append(*seed, Location{X: currentX, Y: currentY + 1})
			*seed = append(*seed, Location{X: currentX + 3, Y: currentY + 1})
			// ROW 3
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY + 2})
			*seed = append(*seed, Location{X: currentX + 3, Y: currentY + 2})
			// ROW 4
			*seed = append(*seed, Location{X: currentX + 2, Y: currentY + 3})
		})
}

func Boat(dimensions Dimensions) []Location {
	// 00-
	// 0-0
	// -0-

	const HEIGHT = 3
	return getRepeatingPattern(dimensions, HEIGHT, HEIGHT,
		func(seed *[]Location, currentX int, currentY int) {
			// ROW 1
			*seed = append(*seed, Location{X: currentX, Y: currentY})
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY})
			// ROW 2
			*seed = append(*seed, Location{X: currentX, Y: currentY + 1})
			*seed = append(*seed, Location{X: currentX + 2, Y: currentY + 1})
			// ROW 3
			*seed = append(*seed, Location{X: currentX + 1, Y: currentY + 2})
		})
}
