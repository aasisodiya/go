# Regexp

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-general.regexp&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

- [Regexp](#regexp)
  - [Remove all numbers from a string](#remove-all-numbers-from-a-string)

## Remove all numbers from a string

In order to remove all number from a string we could have used replaceAll in loop while replacing int i from 0 to 9. And for some reason it's faster But one better way is to use regex. Sample of which is given below

```go
package main

import (
    "log"
    "regexp"
)

func main() {
    log.Println("Using regexp to remove numbers from string")
    sample := "ThisIsASampleStringWithNumbersFrom0To9InOrder0123456789AndThat'sIt"
    numbersRegExp := regexp.MustCompile("[0-9]+")
    op := numbersRegExp.ReplaceAllString(sample, "")
    log.Println(op)
}
```

[Click here](https://go.dev/play/p/IMvOb0A2gSQ) to run this code

Now what we have done here is:

1. We have created a regexp that identifies all the numbers using `regexp.MustCompile("[0-9]+")`
2. Then we are using it as replacement in `sample` string with blanks using `numbersRegExp.ReplaceAllString(sample, "")`

That's it! in simple 2 steps you get your work done.

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&label=aasisodiya/go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
