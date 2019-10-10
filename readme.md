#
This go repo folder conatins all the R&D + it is in gopath so that all packages are available to other projects.
Note: go module (`go mod init go-game`), can not be in go path

## Go dependency management tool - go module
### Starting a new project
https://github.com/golang/go/wiki/Modules


## Go Build & packages

packages are name of folders and this is what go looks to import in import statement.
Normally the directory name is same as the package declared in the code file in that directory.

Go build looks in path `go env GOPATH`, and normally all packages are kept in this 1 directory
Though multiple such diretories can also be defined:
```
export GOPATH=$HOME/go:/Users/d069900/rnd/go/go-school/hello-world
export PATH="${PATH}:$(go env GOPATH)/bin"
```


http://todsul.com/tech/setup-golang-on-mac-os-x/

https://github.com/golang/go/wiki/GOPATH


building from a path
```
go build -o hello src/app/hello.go
```

building and providing location of binary
```
export outputBinary=build/hello
export locationOfMain=/src/app/hello.go
go build -o  $outputBinary $locationOfMain
```

## Go install
??

## Go run 
go run path/xx.go build and run

## go declare variables
- var a,b string = greet.GetGreeings()
- short notation: a,b := greet.GetGreeings()

  