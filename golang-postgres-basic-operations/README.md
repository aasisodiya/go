# Using Postgres with Go Lang

Most part of using Postgres with Go Lang is similar with using MySQL with Go Lang

## Aurora Postgres Connection

For setting up Postgres Connection you have to use following format

```go
db, err := sql.Open("postgres", "dbname=sample_db host=aurora-uw2-instance-1.awsid.us-west-2.rds.amazonaws.com user=username password=password")
```

---

## Q&A

### How to handle empty column data being returned during `row.Scan()`

**Solution:** Use `sql.NullString`, `mysql.NullTime`, `sql.NullInt64` as data type for variable that you are using for receiving value during scan.

Just for reference, if you see below snippet (taken from doc) you can see that SQL handles null a bit differently. The property `Valid` is `true` when String is not NULL else it is `false` for NULL value.

```go
type NullString struct {
    String string
    Valid  bool // Valid is true if String is not NULL
}
```

It can get complicated while creating JSON for the same object. You can refer to this [article I found for doing the same](https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267)

>**Reference:** <br> [Stackoverflow](https://stackoverflow.com/questions/44891030/scan-error-unsupported-scan-storing-driver-value-type-nil-into-type-string) <br> [Stackoverflow: Difference between *string and sql.NullString](https://stackoverflow.com/questions/40092155/difference-between-string-and-sql-nullstring) <br> [Stackoverflow: How do you marshal a sql.NullString such that the output is flattened to give just the value in go?](https://stackoverflow.com/questions/51961358/how-do-you-marshal-a-sql-nullstring-such-that-the-output-is-flattened-to-give-ju/51961903)

Note: *If You want to run your sample go code online you can use below mentioned link*

> **Go Online IDE:** [The Go Playground](https://play.golang.org/)

### How to handle NULL values entry in DB

Using below demonstrated sample function, we can insert NULL values in DB. The given below sample function simply checks if the given data is nil and then accordinly returns NULL equivalent for sql data insertion which in below cases are sql.NullInt64, sql.NullTime. Below are the types of Null in sql package.

- sql.NullBool
- sql.NullString
- sql.NullFloat64
- sql.NullInt32
- sql.NullInt64
- sql.NullTime

Below are the functions you can use to handle/convert respective NULL values

```go
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
```
