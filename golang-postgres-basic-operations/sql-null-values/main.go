package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {
	var b bool
	var s string
	var f float64
	var i int
	var t time.Time

	// Default Values
	fmt.Println(b) // false
	fmt.Println(s) // blank
	fmt.Println(f) // 0
	fmt.Println(i) // 0
	fmt.Println(t) // 0001-01-01 00:00:00 +0000 UTC

	// Null Values
	fmt.Println(newNullBool(nil))    // {false false}
	fmt.Println(newNullString(""))   // { false}
	fmt.Println(newNullFloat64(nil)) // {0 false}
	fmt.Println(newNullInt32(nil))   // {0 false}
	fmt.Println(newNullInt64(nil))   // {0 false}
	fmt.Println(newNullTime(t))      // {0001-01-01 00:00:00 +0000 UTC false}

	// Actual Values
	b = true
	s = "teststring"
	f = 1.2345
	i = 12345
	t = time.Now()
	fmt.Println(newNullBool(&b))    // {true true}
	fmt.Println(newNullString(s))   // {teststring true}
	fmt.Println(newNullFloat64(&f)) // {1.2345 true}
	fmt.Println(newNullInt32(&i))   // {12345 true}
	fmt.Println(newNullInt64(&i))   // {12345 true}
	fmt.Println(newNullTime(t))     // {2009-11-10 23:00:00 +0000 UTC m=+0.000000001 true}
}

// Function to handle NULL Value for int32
func newNullInt32(s *int) sql.NullInt32 {
	if s == nil {
		return sql.NullInt32{}
	}
	i := int32(*s)
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// Function to handle NULL Value for int64
func newNullInt64(s *int) sql.NullInt64 {
	if s == nil {
		return sql.NullInt64{}
	}
	i := int64(*s)
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

// Function to handle NULL value for Timestamp
func newNullTime(t time.Time) sql.NullTime {
	var tempTime time.Time
	if t.Equal(tempTime) {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// Function to handle NULL value for Float64
func newNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: *f,
		Valid:   true,
	}
}

// Function to handle NULL value for String
func newNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// Function to handle NULL value for Boolean
func newNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *b,
		Valid: true,
	}
}
