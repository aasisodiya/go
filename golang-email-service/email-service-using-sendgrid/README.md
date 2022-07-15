# Using SendGrid with GoLang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-email-service.sendgrid&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## Getting Started

**Important:** Make sure to change the key value in sendgrid.go

``` golang
const (
    // Key for SendGrid
    Key = "PUT_YOUR_SENDGRID_KEY_HERE"
)
```

Import Required Packages

## Step 1: Create Sender Object

Here sender is the source that is sending the email

``` golang
var s models.Sender
s.SenderEmail = "example@test.com"
s.SenderName = "SendGrid"
```

## Step 2: Create Recipient Object/s

Here Recipient is the target that will be receiving the email

**Note:** If you want to keep the Recipient in CC then set RecipientCCBCC value to "CC" and if you want to keep them in BCC set it to "BCC", for all other values the user is kept in "TO" target by default.

``` golang
var r1 models.Recipient
r1.RecipientName = "Akash"
r1.RecipientEmail = "akash_sisodiya@test.com"
r1.RecipientCCBCC = "none"
// r1.RecipientCCBCC = "CC"
// r1.RecipientCCBCC = "BCC"
```

## Step 3: Create Recipients Array

Create an array of Recipients, and add all intended recipient to the array. 

(*Example: Consider there are 3 users and all are part of one common email, then only append the users, but if you want to send them individually then you will have to create an array of single user only*)

```golang
var r []models.Recipient
r = append(r, r1)
```

## Step 4: Send Email

Now use method `SendEmail(recipients, sender, subject, message)` to send an email

```golang
subject := "Sample Subject"
plainTextMessage := "This is Test Message!"
htmlTextMessage := "<h1>This is Test Message but in HTML!</h1>"

// Sending one sample Plain Text Email
services.SendEmail(r, s, subject, plainTextMessage)
// Sending one sample HTML body Email
services.SendEmail(r, s, subject, htmlTextMessage)
```

## Final Step: Execute & Check Response

You will get `202` status code if everything goes well.

## Troubleshooting

* If you get 401 status code, it can mean that you have provided wrong SendGrid Key, so check it once and try again
* If you see multiple users in "TO" but wanted independent email for every single user then you might have made mistakes in step 3. If you want all users to receive independent email then you will have to run the code in for loop with Recipients array being created with only one Recipient object
* User don't come in CC/ BCC, then you have made mistake in step 2. RecipientCCBCC flag needs to be set only to CC / BCC - ALL CAPS and for all other values (even cc/bcc) it will neglect the flag

## Reference
[Source](https://github.com/sendgrid/sendgrid-go, "Github")
