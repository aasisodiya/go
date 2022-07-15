# Using MySQL with Go Lang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-mysql-basic-operations&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## MySQL Windows Setup

> **Download Link:** [Download](https://dev.mysql.com/downloads/installer/)

Go with mysql-installer-web-community-X.X.X.X.msi and download the installer and install it before proceeding else you won't be able to test the code on local setup. We won't be covering the setup procedure in this document, in case of any issues refer this [video](https://www.youtube.com/watch?v=UgHRay7gN1g). You can even use your own MySQL server settings instead of local server, just keep those server details handy.

Note: *Before Proceeding ahead ensure you have an active MySQL Server running and accessible from your machine*

### Required MySQL Server Details

- Username
- Password
- IP Address (127.0.0.1 for local)
- Post (3306 is default)
- Schema Name
- Table Name

Note down above details as we will be needing it ahead

## Go Code

> **Important Note:** Below code isn't meant for Copy Paste Use, you will have to write your own correct code in order to execute the code. Below I have just explained important parts of the code

### Import Required Libraries

For MySQL operations we will be using following libraries

```golang
import (
"database/sql"
_ "github.com/go-sql-driver/mysql"
)
```

> `"github.com/go-sql-driver/mysql"` should be imported with _ (underscore)

### Establish Connection with MySQL Server

Here below we use function: func Open(driverName, dataSourceName string), The Open function should be called just once.

```go
    // Open our database connection.
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testschema")
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    // Open may just validate its arguments without creating a connection
    // to the database. To verify that the data source name is valid, call
    // Ping.
    err = db.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    // defer the close till after the main function has finished
    defer db.Close()

```

Repeating this because it can be important, `sql.Open` may just validate its arguments without creating a connection to the database. So to verify that the data source name is valid, call function `Ping` on `db` object.

```go
    err = db.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
```

### Performing required operations

- #### Perform Insert Operation

    First prepare a query statement

    ```go
        statement := "INSERT INTO " + tableName + " VALUES( ?, ?, ? )"
        // Prepare statement for inserting data, this will depend on your table
        stmtIns, err := db.Prepare(statement) // ? = placeholder
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        defer stmtIns.Close()  // Close the statement when we leave main() / the program terminates
    ```

    Now we have created a ready to use statement that we can reuse

    ```go
        _, err = stmtIns.Exec(ID, lname, fname) // Insert the record using statement
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    ```

- #### Perform Select Operation

    Here also you can prepare a statement if required, but for now we are going with direct approach

    ```golang
        rows, err := db.Query("SELECT * FROM Persons")
        // above statement will fetch records in rows variable
        defer rows.Close()
        // if there is an error here, handle it
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    ```

    Here `rows` variable will contain all our query data, now processing/extraction of that data will totally vary from use case to use case. Example for is given below

    ```golang
        var PersonID int
        var LastName string
        var FirstName string
        for rows.Next() {
            // for each row, scan the result into defined variables
            err = rows.Scan(&PersonID, &LastName, &FirstName)
            if err != nil {
                panic(err.Error()) // proper error handling instead of panic in your app
            }
            fmt.Printf("PersonID: %d, FirstName: %s, LastName: %s \n", PersonID, FirstName, LastName)
        }
    ```

> For other operations you can refer to this [link](https://godoc.org/github.com/go-sql-driver/mysql)

## Troubleshooting

- **Error:** `imported and not used: "github.com/go-sql-driver/mysql"`

    **Solution:** `"github.com/go-sql-driver/mysql"` should be imported with _ (underscore) so try using `_ "github.com/go-sql-driver/mysql"`

## Reference

- [Database in Go Lang](https://golang.org/pkg/database/sql/)
- [How to Install MySQL Server on Windows](https://www.youtube.com/watch?v=UgHRay7gN1g)
- [MySQL Operations in Go Lang](https://godoc.org/github.com/go-sql-driver/mysql)
- [Error: imported and not used: "github.com/go-sql-driver/mysql"](https://stackoverflow.com/questions/36256230/connection-fails-with-mysql-using-golang)
