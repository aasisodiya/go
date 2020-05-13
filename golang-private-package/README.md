# Use/Import Private Repo in GoLang

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
$ go get -v some.private.repourl/folder
go: finding some.private.repourl/folder latest
go: downloading some.private.repourl/folder v0.0.0-20190921178888-98a48a9as456
verifying some.private.repourl/folder@v0.0.0-20190921178888-98a48a9as456: some.private.repourl/folder@v0.0.0-20190921178888-98a48a9as456: reading https://sum.golang.org/lookup/some.private.repourl/folder@v0.0.0-20190921178888-98a48a9as456: 410 Gone
```

If you get similar error `410 Gone` then it might be because of the URL being a private repo

### Solution

* in cmd.exe:

    ``` bat
    set GOPRIVATE=some.private.repourl/folder
    ```

* in Powershell:

    ``` posh
    $env:GOPRIVATE = "some.private.repourl/folder"
    ```

### Note

* `GOPRIVATE = "some.private.repourl/folder/"` will not work because of `/` at the end
* `GOPRIVATE = "some.private.repourl/folder,some.otherprivate.repourl/folder"` you can add more than one private repo by separating them with comma (,)
