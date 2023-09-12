# Unused Code in Golang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-unused-code&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

In order to better optimize the code its always better to remove/get rid of the unused code. Many IDEs and code editors have plugins or built-in features that can highlight unused code in your project as you work on it. For example, Visual Studio Code with the Go extension has this feature.

But sometimes, specially in case of VSCode the extension don't work as expected. So you can try below methods to find the unused code.

## Install `golangci-lint`

Use below command to install `golangci-lint`, it can take some time to complete.

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

To check and verify the installation run command `golangci-lint --version`

## Use `golangci-lint run --enable unused` command to identify unused code

Sample output looks like below:

```bash
C:\Users\akash\Documents\GitHub\go\golang-general\unused-code>golangci-lint run --enable unused
main.go:15:6: func `test1` is unused (unused)
func test1(){
     ^
```

## TODO

- Check for Ways to find unused code that is present in nested packages
