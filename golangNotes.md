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