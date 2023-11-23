package main

// Step 2: Import Package
import (
	"fmt"
	"main/aws"
)

// Step 3: Call Functions
func main() {
	email := "akash_sisodiya@test.com"
	target := "test@test.com"
	fmt.Println("Main")
	test, err := aws.ListValidEmails(aws.DefaultRegion)
	printIfError(err)
	fmt.Println(test)
	test, err = aws.ListAllEmails(aws.DefaultRegion)
	printIfError(err)
	fmt.Println(test)
	test2, err := aws.VerifyEmailID(target, aws.DefaultRegion)
	printIfError(err)
	fmt.Println(test2)
	_, err = aws.SendEmail(target, email, "Subject", "Text", "<h1>HTML</h1>", aws.DefaultRegion) //won't show "Text" in email body
	printIfError(err)
	_, err = aws.SendEmail(target, email, "Subject", "", "<h1>HTML</h1>", aws.DefaultRegion) // Only HTML
	printIfError(err)
	test3, err := aws.SendEmail(target, email, "Subject", "Text", "", aws.DefaultRegion) // Only Text
	printIfError(err)
	fmt.Println(test3)
	test4, err := aws.RemoveEmail(target, aws.DefaultRegion)
	printIfError(err)
	fmt.Println(test4)
	err = aws.SESStatistics(aws.DefaultRegion)
	printIfError(err)
}

func printIfError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
