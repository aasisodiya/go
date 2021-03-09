# QnA in Go Lang

## Marshaling Empty Array Gives Null

**Problem**: While Marshaling an Empty Array of an Object while sending back the response in JSON in Lambda, It showed response as Null. (Sample to demosntrate the same is given below.) But I needed response as an empty array in JSON to represent proper data.

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
