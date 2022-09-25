// This line is package Synopsis, and will be shown under synopsis and also under overview
// This will also get inline, even though it's on new line in comment
//
// This is a overview only line as it is not the first line
package godoc

import "fmt"

const (
	// PI is the world wide known constant for pi
	PI = 3.14
)

// GetCircleArea returns you the area of the circle with given radius
func GetCircleArea(radius float64) (area float64) {
	area = PI * radius * radius
	fmt.Println(area)
	return
}

// CircleArea prints you the area of the circle with given radius
func CircleArea(radius float64) {
	area := PI * radius * radius
	fmt.Println(area)
}

// Command to open godoc : C:\Users\akash\go\bin\godoc.exe -http=:6060
