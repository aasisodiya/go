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
