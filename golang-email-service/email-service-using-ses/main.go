package main

// Step 2: Import Package
import (
	"main/aws"
	"fmt"
)

// Step 3: Call Functions
func main() {
	email := "akash_sisodiya@test.com"
	target := "test@test.com"
	fmt.Println("Main")
	test, err := aws.ListValidEmails("us-west-2")
	printIfError(err)
	fmt.Println(test)
	test, err = aws.ListAllEmails("us-west-2")
	printIfError(err)
	fmt.Println(test)
	test2, err := aws.VerifyEmailID(target, "us-west-2")
	printIfError(err)
	fmt.Println(test2)
	test3, err := aws.SendEmail(target, email, "Subject", "Text", "<h1>HTML</h1>", "us-west-2") //won't show "Text" in email body
	printIfError(err)
	test3, err = aws.SendEmail(target, email, "Subject", "", "<h1>HTML</h1>", "us-west-2") // Only HTML
	printIfError(err)
	test3, err = aws.SendEmail(target, email, "Subject", "Text", "", "us-west-2") // Only Text
	printIfError(err)
	fmt.Println(test3)
	test4, err := aws.RemoveEmail(target, "us-west-2")
	printIfError(err)
	fmt.Println(test4)
	err = aws.SESStatistics("us-west-2")
	printIfError(err)
}

func printIfError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}