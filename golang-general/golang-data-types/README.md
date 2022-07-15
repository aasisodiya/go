# Data Conversion in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-general.data-types&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

Go's basic types are

* bool
* string
* int  int8  int16  int32  int64
* uint uint8 uint16 uint32 uint64 uintptr
* byte // alias for uint8
* rune // alias for int32 and also represents a Unicode code point
* float32 float64
* complex64 complex128

The example below shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

```go
var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

> Reference: [A Tour of Go](https://tour.golang.org/basics/11)

|Conversion  |int|float|bool|string|
|------------|-----------------------------------------------------------------|------------------------------------|--------------------|---------------------------------------------------------------------|
|var i int   |i                                                                |float64(i) <br> float32(in)         |                    |strconv.Itoa(i)<br>strconv.FormatInt(i, [10](base))                  |
|var f float |int(f)                                                           |f                                   |                    |strconv.FormatFloat(f, ['E'](format), [-1](precision), [64](bitsize))|
|var b bool  |                                                                 |                                    |b                   |strconv.FormatBool(b)                                                |
|var s string|strconv.Atoi(s)<br>strconv.ParseInt(s, [10](base), [64](bitsize))|strconv.ParseFloat(s, [32](bitsize))|strconv.ParseBool(s)|s                                                                    |

> Reference: [strconv - Go Doc](https://golang.org/pkg/strconv/)
