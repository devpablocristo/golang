# stygis - golang Hexagonal Architecture 

Just another sample of hexagonal architecture for golang
focus on simplifed code, orgenized structure and better naming for functions and packages name
**No duplicate naming for package**

### old version:
https://github.com/iDevoid/stygis/tree/master

## Project structure

```sh
stygis/
├── bin                                 
├── cmd                                 # Packages that provide support for a specific program that is being built, this contains cmd for REST, gRPC, etc
├── initiator                           # initialize any domain here that will be called inside cmd, so, no spaghetti init on cmd main package
├── internal                            # Contains all application packages (all functions must have return)
      ├── constants                     # For declaring all constants variable based on it's entity, declaring with full name descripting it
            ├── model                   # model is where the declaration of structs are being written
            ├── query                   # where query is being declared for storage sql, and it's better using sqlx naming system for passing the parameters
            ├── state                   # contains all constants variable (string, int, etc..) like status, or constant strings
      ├── glue                          # is just another name of middleware
            ├── routing                 # where the routing for handlers are assigned based on method and url
      ├── handler                       # any protocol interface that needs handler is being declare here, whether for cleaning input data, error payload handling, etc..
            ├── rest                    # handler for rest API technology
      ├── module                        # Domain packages are here, contains business logic and interfaces that belong to each domain, and no usecase calling other usecase
            ├── user                    # only sample domain, user package which handler for user business logic
                ├── initiator.go        # this is the only file to declare interface methods from storage and repository. also where to put func init the package.
                ....                    # name the file based on the usecase that may contains multiple acticity; example: login usecase for only one activity, profile contains edit profile and show profile
      ├── repository                    # this may a bit weird for you, this package uses for data storing logic, it is an optional if your domain only saves data to one db, but it's different when a domain uses multiple storage, for example caching and multiple persistences
      ├── storage                       # this is where you put the data storing code. whether persistence like postgresql, monggodb, etc. and caching like redis, etc. 
            ├── cache                   # package where caching storing code is written based on its domain.
            ├── persistence             # like its name, this package contains storing code for SQL or noSQL db
├── mocks                               # this contains all mock from any source. it contains some rules. it must be a generate files and the package name must be mock_(source)
      ├── psql                          # this is mock for psql contains package mock_psql, even tho it is written manually, it must have // Code generated manually. DO NOT EDIT.
      ├── redis                         # this is mock for redis contains package mock_redis, even tho it is written manually, it must have // Code generated manually. DO NOT EDIT.
      ├── user                          # this is generated by golang mock official generator
├── platform                            # external app for all uses (in here you can make function without return)
      ├── encryption                    # you can have any encryption method here
      ├── postgres                      # contains functions to open database postgres connections, with mutiple servers can be added, like db master and/or slave
      ├── redis                         # contains functions to open database redis connection, currently it uses only one connection, but this can be adjust just like the postgres connections
      ├── routers                       # contains functions to serve the HTTP Listener using all registered URLs with the handlers
├── static                              # to store any static file like pic, html, etc 
```

## The Flow

```
constants (optional) -> glue (optional) -> handler -> module -> repository and/or storage


1. contants : depends on the development you may need to add some structs for data structuring that later can be used to the end of the flow
2. glue : depends on the technology you use
3. handler : requests are being orgenized, whether convertions, checking bad payloads, etc..
4. module : the business logic for processing the requests, everything that domain business logic has is centralized here, and write the log inside this package only
5. repository : use this if the business logic need the data logic that combine multiple data storage
6. storage : write the template of CRUD operation whatever the technology is
```

## How TO Mock
Before you start to generate the mocks for any domain. Make sure that all interfaces are in initiator.go under your domain package.

1. inside every initiator.go, put `//go:generate mockgen -destination=../../../mocks/(source)/(file name destination).go -source=initiator.go` under your `package yourdomain` line or under the first line
example : `//go:generate mockgen -destination=../../../mocks/user/user_mock.go -source=initiator.go`

2. run `go generate ./...` on your terminal (your directory position is the root of your project)
this will runs the command you put inside the initiator.go
the `-destination=../../../mocks/(source)` is to set the generated files to be all under mocks folder in your root project, so all mocks will be centralized under mocks package

How To:
1. the package name (in file) must not be the same as the source
2. to call the mock, write the package name (in file). for example: `mock_user.NewMockCaching(ctrl)`
at least on VSCODE it automatically generate import `mock_user "github.com/stygis/mocks/user`
why it must be `mock_(source)`. because this is not a package of something to implement on the app. you cannot call the mock just using `user.NewMockCaching(ctrl)` or `mock.NewMockCaching(ctrl)`. this is because unit test should be written manually, and the programmer who's writing it understands that unit test is important as the main code. 
also in technical reason, if we put the mock inside one package `mocks` as I did before. the other domain cannot make its mock because the same name of interfaces. this has a good side because you don't get many mocks function if the mocks are centralized inside one package, and the code and structure stay clean and organized.

## Notes

Be aware that this may not be the better structure for your application, but i'm hopping to write into more general structure and stay simple

There is no duplication of package naming here, you can see that the there is `rest` folder inside `cmd` and `handler`. But the `cmd` doesn't use the `rest` package name, it uses `main` package name.

# Medium
[stygis — Golang Hexagonal Architecture](https://medium.com/p/a2d89d01f84b)

## References
[Ashley McNamara + Brian Ketelsen. Go best practices.](https://youtu.be/MzTcsI6tn-0)

## golang Hexagonal Arch Repositories
1. [Gira (A Hex Example)](https://github.com/Holmes89/hex-example)
2. [Kat Zień - Achieving maintainability with hexagonal architecture](https://github.com/katzien/go-structure-examples/tree/master/domain-hex)