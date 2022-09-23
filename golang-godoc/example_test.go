package godoc_test

import "github.com/aasisodiya/go/godoc"

func Example() {
	godoc.CircleArea(10.0)
	// Output:
	// 314
}

// go test -timeout 30s -run ^Example$ github.com/aasisodiya/go/godoc

// Prints the area on terminal
func ExampleCircleArea() {
	godoc.CircleArea(10.0)
	// Output:
	// 314
}

// go test -timeout 30s -run ^ExampleCircleArea$ github.com/aasisodiya/go/godoc

// Prints the area on terminal
func ExampleGetCircleArea() {
	godoc.GetCircleArea(10.0)
	// Output:
	// 314
}

// go test -timeout 30s -run ^ExampleGetCircleArea$ github.com/aasisodiya/go/godoc

// ! This below example won't show up in doc as after _ you need lowercase
// Prints the area on terminal
func ExampleGetCircleArea_R100() {
	godoc.GetCircleArea(100.0)
	// Output:
	// 31400
}

// go test -timeout 30s -run ^ExampleGetCircleArea_R100$ github.com/aasisodiya/go/godoc

// Prints the area on terminal
func ExampleGetCircleArea_r100() {
	godoc.GetCircleArea(100.0)
	// Output:
	// 31400
}

// go test -timeout 30s -run ^ExampleGetCircleArea_r100$ github.com/aasisodiya/go/godoc
