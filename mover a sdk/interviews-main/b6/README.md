Receiver Type
Choosing whether to use a value or pointer receiver on methods can be difficult, especially to new Go programmers. If in doubt, use a pointer, but there are times when a value receiver makes sense, usually for reasons of efficiency, such as for small unchanging structs or values of basic type. Some useful guidelines:

If the receiver is a map, func or chan, don't use a pointer to them. If the receiver is a slice and the method doesn't reslice or reallocate the slice, don't use a pointer to it.
If the method needs to mutate the receiver, the receiver must be a pointer.
If the receiver is a struct that contains a sync.Mutex or similar synchronizing field, the receiver must be a pointer to avoid copying.
If the receiver is a large struct or array, a pointer receiver is more efficient. How large is large? Assume it's equivalent to passing all its elements as arguments to the method. If that feels too large, it's also too large for the receiver.
Can function or methods, either concurrently or when called from this method, be mutating the receiver? A value type creates a copy of the receiver when the method is invoked, so outside updates will not be applied to this receiver. If changes must be visible in the original receiver, the receiver must be a pointer.
If the receiver is a struct, array or slice and any of its elements is a pointer to something that might be mutating, prefer a pointer receiver, as it will make the intention clearer to the reader.
If the receiver is a small array or struct that is naturally a value type (for instance, something like the time.Time type), with no mutable fields and no pointers, or is just a simple basic type such as int or string, a value receiver makes sense. A value receiver can reduce the amount of garbage that can be generated; if a value is passed to a value method, an on-stack copy can be used instead of allocating on the heap. (The compiler tries to be smart about avoiding this allocation, but it can't always succeed.) Don't choose a value receiver type for this reason without profiling first.
Don't mix receiver types. Choose either pointers or struct types for all available methods.
Finally, when in doubt, use a pointer receiver.~~


So how to choose between Pointer vs Value receiver?

If you want to change the state of the receiver in a method, manipulating the value of it, use a pointer receiver. It’s not possible with a value receiver, which copies by value. Any modification to a value receiver is local to that copy. If you don’t need to manipulate the receiver value, use a value receiver.

The Pointer receiver avoids copying the value on each method call. This can be more efficient if the receiver is a large struct,

Value receivers are concurrency safe, while pointer receivers are not concurrency safe. Hence a programmer needs to take care of it.

Notes-

Try to use the same receiver type for all your methods as much as possible.
If state modification needed, use pointer receiver if not use value receiver.







# Directories

1. domain
2. usecases
3. interfaces
   1. controllers
   2. gateways
   3. presenters
4. infrastructure
   1. devices
   2. web
   3. db
   4. external
   5. ui

# Folder structure

- entity
  - domain
  - usecases  
  - interfaces
  - infrastructure

# Layer - Directory

- Frameworks & Drivers = infrastructure
- Interface = interfaces
- Usecases = usecases
- Entities = domain


# Golang file struct

1. interface
2. struct
3. func new
4. methods


# Name convencitions

- interface: handler
- struct: package name like

quizas en lugar de common deberi ir internal o pkg o src o utils u otra cosa?
para que pingo se usa internal?

To practice DIP in Golang, it's best to understand Golang's idea of "Accepts interfaces, return structs".

By defining an interface and returning a struct according to the idea of "Accept interfaces, return structs", the dependency will be directed to the interface.

```  go
package examples

// Logger is an interface which will be used for an argument of a function.
type Logger interface {
    Printf(string, ...interface{})
}

// FooController is a struct which will be returned by function.
type FooController struct {
    Logger Logger
}

// NewFooController is a function for an example, "Accept interfaces, return structs".
// Also, this style of a function take on a role of constructor for struct.
func NewFooController(logger Logger) *FooController {
    return &FooController{
        Logger: logger,
    }
}
```

context
goroutines
tests
grpc
graphql
rest
http/net
gin
tdd
microservice


Domain: Domain is the business logic that is independent of the application. In other words, the domain is common for the whole organization.

Use Case: Use Case is business logic that is particular to the application.

Interface: Interface consists of controllers, adapters, presenters. Interfaces allow us to connect our code to Infrastructure.

Infrastructure: Infrastructure is the layer where all the detail go. Web, Database, third-party libraries, file systems all belong to the interface layer.


Directories structure

cmd -> main.go
|
service 1
|
service 2
| 
service 3 -> entity 1
             entity 2
             entity 3 -> domain
                         usecases
                         interfaces
                         infrastructre
                         
