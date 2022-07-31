# Use/Import Private Repo in GoLang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-private-package&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## `go env` Command

`go env` command prints information about environment variables that affect the behavior of Go tools

```bat
go> go env
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\akash\AppData\Local\go-build
set GOENV=C:\Users\akash\AppData\Roaming\go\env  
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\akash\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=c:\go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=c:\go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\akash\AppData\Local\Temp\go-build859714054=/tmp/go-build -gno-record-gcc-switches
```

## Error with `go get private-repo`

```bat
$ go get -v some.private.repo.git
go: finding some.private.repo.git latest
go: downloading some.private.repo.git/folder v0.0.0-20190921178888-98a48a9as456
verifying some.private.repo.git/folder@v0.0.0-20190921178888-98a48a9as456: some.private.repo.git/folder@v0.0.0-20190921178888-98a48a9as456: reading https://sum.golang.org/lookup/some.private.repo.git/folder@v0.0.0-20190921178888-98a48a9as456: 410 Gone
```

If you get similar error `410 Gone` then it might be because of the URL being a private repo

### Solution

The new `GOPRIVATE` environment variable indicates module paths that are not publicly available. It serves as the default value for the lower-level `GONOPROXY` and `GONOSUMDB` variables, which provide finer-grained control over which modules are fetched via proxy and verified using the checksum database.

* in cmd.exe:

    ``` bat
    set GOPRIVATE=some.private.repo.git
    set GONOSUMDB=some.private.repo.git
    ```

* in Powershell:

    ``` posh
    $env:GOPRIVATE = "some.private.repo.git"
    $env:GONOSUMDB = "some.private.repo.git"
    ```

### Note

* `GOPRIVATE = "some.private.repo.git/"` will not work because of `/` at the end
* `GOPRIVATE = "some.private.repo.git,some.otherprivate.repo.git"` you can add more than one private repo by separating them with comma (,)
* Connection Time Out Error on EC2

  ```posh
  exit status 128:
  fatal: unable to connect to git-codecommit.us-west-2.amazonaws.com:
  git-codecommit.us-west-2.amazonaws.com[0: ww.xx.yy.zz]: errno=Connection timed out
  ```

  In this case, you will have to provide ec2 an access to Codecommit and then set this environment variable

  ```posh
  set GONOSUMDB=git-codecommit.us-west-2.amazonaws.com/v1/repos/some.private.repo.git
  ```

  Better way will be to simply store this variables in `/etc/environment` permanently. `/etc/environment` stores all key-value pairs of environment variables

## Reference

* [Private Repo in Go](https://stackoverflow.com/questions/57885949/private-repo-go-1-13-go-mod-failed-ping-sum-golang-org-lookup-ver/57887036#57887036)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
