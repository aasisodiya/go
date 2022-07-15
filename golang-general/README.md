# General

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-general&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

- [General](#general)
  - [`os.Exit()` Notes](#osexit-notes)
  - [Links](#links)

## `os.Exit()` Notes

- Use `os.Exit` to immediately exit with a given status
- `defers` will not be run when using `os.Exit`
- Specifying an exit code above zero usually means an error has occurred.

| Exit Code Number | Meaning                                                    | Example                 | Comments                                                                                                     |
| ---------------- | ---------------------------------------------------------- | ----------------------- | ------------------------------------------------------------------------------------------------------------ |
| 1                | Catchall for general errors                                | let "var1 = 1/0"        | Miscellaneous errors, such as "divide by zero" and other impermissible operations                            |
| 2                | Misuse of shell builtins (according to Bash documentation) | empty_function() {}     | Missing keyword or command, or permission problem (and diff return code on a failed binary file comparison). |
| 126              | Command invoked cannot execute                             | /dev/null               | Permission problem or command is not an executable                                                           |
| 127              | "command not found"                                        | illegal_command         | Possible problem with $PATH or a typo                                                                        |
| 128              | Invalid argument to exit                                   | exit 3.14159            | exit takes only integer args in the range 0 - 255 (see first footnote)                                       |
| 128+n            | Fatal error signal "n"                                     | kill -9 $PPID of script | $? returns 137 (128 + 9)                                                                                     |
| 130              | Script terminated by Control-C                             | Ctl-C                   | Control-C is fatal error signal 2, (130 = 128 + 2, see above)                                                |
| 255\*            | Exit status out of range                                   | exit -1                 | exit takes only integer args in the range 0 - 255                                                            |

> [Exit Code Reference](https://tldp.org/LDP/abs/html/exitcodes.html)

## Links

- [Data Conversion in Go](https://aasisodiya.github.io/go/golang-general/golang-data-types/)
- [QnA](https://aasisodiya.github.io/go/golang-general/golang-qna/)
