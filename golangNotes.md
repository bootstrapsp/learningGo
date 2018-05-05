# GoLang Notes

* Go is Strictly/Strong Typed
* Tpyes checked at compile time
* Types can't be mixed
* Introduced in Go, keeping in mind the problems encountered in C / C++

### Note : Strictly Typed e.g. if _i_ is defined as int, one can only interact with it as an integer type.

* Go has Typed and Untyped
    * Untyped e.g.  x := 33.40
    * Typed e.g. y:=int(42)

* Attempting to do z = x + y, as defined above, will throw exception "invalid operation: x + y (mismatched types float64 and int64)
"

* As long as something is untyped it can be combined with other untyped variable
* Go is **Pass by value**

## Control Flow

* Go does inlining while reading the code from top to bottom for better performance
* There's no while or do while loops in Go

## Rune

* Is a character
* In Go Rune is alias for int32 number
* Like 100 0001 in binary represents "A"
* Rune represents int32
* Rune is made of 32 0s & 1s or 4 bytes (4*8) = 32
* UTF-8 is a 4 byte coding scheme

## Types of Reference Types

* Map
* Slice: dynamica, doesn't need definition like []<dataTypes>
* 
* Array : non dynamic, need definition like [x]<dataType>
* Channel

## Struct

* Struct is a sequence of named elements called fields
    * Each of which has a name and a type
    * Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField)
    * Within a struct, non-blank field names must be unique
    * A field declared with a type but no explicit field name is an anonymous field, also called an embedded field or an embedding of the type in the struct. An embedded type must be specified as a type T or as a pointer to a non-interface type name *T and T itself may not be a pointer type.
    * The unqualified type name acts as the field name
    * It's a aggregate type, i.e. collecting bunch of other types in list of fields e.g.
    first string
    last string
    age int

- instead of creating a whole class just for this, a struct can be created for this

## Go is an Object Oriented programming language

* Encapsulation
    * state ("fields")
    * behaviour ("method")
    * exported / un-exported

* Reusability
    * inheritence ("embedded types")

* Polymorphism
    * interfaces

* Overriding
    * Promotion

## Variadic function

Type of Go function that takes multiple inputs using "..." , **in short** : variadic functions are the functions that can have n number of parameters of a type, this is done by putting "..." in front of the type

## Go feature list

### Formatting

* *gofmt* formats the go code
* Tabs used for the Indentation are ignored
* No line length limit

### Commenting

* can be done either by doing /* */ or // for single line comments
* pacakges must have comments detialing out the use case for the package

## Go's Memory Model

#### Advice

* Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access

Marshal vs Unmarshal vs Encode vs Decode

* Marshal and UnMarshal deals with the string or slice of byte within the application
* Have nothing to do with the external data

## Marshal, Unmarshal, Encoding & Decoding

* UnMarshal is used for parsing the JSON data and stores the value in a pointer

* Marshal encodes the data 

* Encode writes the value streaming out of the application

* Decode reads the value streaming into the application

## More tips

* avoid using types, e.g. var x int, best is to leave out the type that is more idiomatic

## Method sets

* Receivers in the Method are of two type

    * **value receiver**
        * Value Type
        * Pointer Type
    *  **Pointer receiver**
        *  Pointer Type

Another way to represent this

**Receivers**       **Values**
(t T)           T and *T
(t *T)           *T

## Wait Group

* Waits for a collection of goroutines to finish
* Main go routine calls Add to set the number of goroutines to wait for 
* Each go routines runs and calls Done when finished
* Wait can be used to block go routines have finished

## Concurrency

* Is composition of independently executing processes

## Parallelism

* Is the simultaneous execution of (possibly) related computation

## Init()

* It's the first program that executes
* Can be used for setting very first configuration for the page

## Pipeline Channels

* structure defining in/out for set of channels