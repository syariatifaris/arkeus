## Arkeus


Its a simple web api generator including some basic framework for:

1. Handler
2. Module
3. Model, and
4. Repository

##
### Instalation

**A. Using Go Get**

1. `go get github.com/syariatifaris/arkeus` to obtain the latest project
2. `cd $GOPATH/src/github.com/syariatifaris/arkeus` to enter the project
3. `go get -d ./...` to obtain all framework dependencies

**B. Using Dep**

1. `dep ensure` or `dep ensure -v`

##
### How to use

Sample usage: 
```$xslt
./arkeus -gosrc=/your/go/src -project_path=github.com/syariatifaris/diklet -module="page"
```

| Param  | Description |
| ------------- | ------------- |
| gosrc  | Your **$GOPATH/src** filepath  |
| project_path  | Project directory inside $GOPATH/src  |
| module  | Sample module name  |
| port  | Web application port (9093 for default)  |


##
### Project Structure

After the project is generated successfully, you project app projects follows:
```$xslt
[app]
-----[core]
--------core.go
-----[handler]
--------modulename.handler.go
-----[module]
--------[modulename]
---------------[model]
------------------modulename.model.go
---------------[repo]
------------------modulename.repo.go
main.go

```

##
### Dependencies

List of external dependencies (libraries):


```$xslt
github.com/agtorre/gocolorize
github.com/otiai10/copy 
github.com/karlkfi/inject
github.com/dgrijalva/jwt-go
github.com/jmoiron/jsonq
github.com/rubyist/circuitbreaker
github.com/didip/tollbooth
github.com/gorilla/mux
github.com/facebookgo/grace/gracehttp
```

This libraries is being used on **core/** for helpers and framework purpose