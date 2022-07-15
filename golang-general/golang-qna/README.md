# QnA in Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-general.qna&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

- [QnA in Go Lang](#qna-in-go-lang)
  - [Marshaling Empty Array Gives Null](#marshaling-empty-array-gives-null)
  - [`go get -u` Not updating the go build](#go-get--u-not-updating-the-go-build)
  - [`invalid version control suffix in github.com/ path` Error message](#invalid-version-control-suffix-in-githubcom-path-error-message)

## Marshaling Empty Array Gives Null

**Problem**: While Marshaling an Empty Array of an Object while sending back the response in JSON in Lambda, It showed response as Null. (Sample to demonstrate the same is given below.) But I needed response as an empty array in JSON to represent proper data.

```go
var sampleArray []int
fmt.Println(sampleArray) //[]
jsonObject, _ := json.Marshal(sampleArray)
fmt.Println(string(jsonObject)) //  null

sampleArray2 := []int{}
fmt.Println(sampleArray2) //[]
jsonObject2, _ := json.Marshal(sampleArray2)
fmt.Println(string(jsonObject2)) // []
```

**Reason/Solution**: So the solution was to make sure that if the array was never initialized then simply return the initialized empty array `[]int{}`. This way when you Marshal it, it doesn't gives you null and instead gives an empty array as expected. (Code to demonstrate the same is in above sample)

---

## `go get -u` Not updating the go build

**Problem**: I had an external package on codecommit, which I was using in my code. Now there were some updates made to the external package (on codecommit) and I wanted to use the same updated code. So I tried running `go get -u git-codecommit.us-west-2-url` to update the version. But when I build the code the version resets in go.mod to the old version.

**Reason/Solution**: So the reason for this was that my internal shared packages were referring to the old version of the external package (on codecommit). So i had to update those referenced version. This in turn resolved the issue.

## `invalid version control suffix in github.com/ path` Error message

**Problem**: While building my code, which referenced private packages, I was getting below error

```bash
package github.com/aasisodiya/general.git: invalid version control suffix in github.com/ path
```

**Reason/Solution**: `.git` suffix at the end of the github url was thee root cause, which obviously is incorrect and a silly mistake. So simply changing the `github.com/aasisodiya/general.git` to `github.com/aasisodiya/general` in `go.mod` fixed the issue
