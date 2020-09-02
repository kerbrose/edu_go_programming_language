### Chapter 01: Tutorial

- p.2, runnging script file using `run` vs. `build`

```bash
go run helloworld.go # or
go build helloworld.go
```

- p.6 different for loop options
```go
for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
}
// or
for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
}
// or
var i int = 1
for i < len(os.Args){
    s += sep + os.Args[i]
    sep = " "
    i++
}
// or
for i := 1; i < 5; i++ {
    sum += i
}
// or
for {
    // a traditional infinite loop
}
```

- p.9 datatype `map` is like dictionary in python
- p.12 a way to import sub-packages
```go
import (
    "fmt"
    "io/ioutil" // an example of this
    "os"
    "strings"
)
```

- p.14 intro to `composite literals`
```go
var palette = []color.Color{color.White, color.Black}
anim := gif.GIF{LoopCount: nframes}
```
- p.18 intro to `goroutine` & `channel`. you can consider a goroutine as thread

- p.19 intro to **fmt.Fprintf**, please note the difference between `Fprintf` vs `Sprintf` & their usage from [gopl.io/ch1/server1], [gopl.io/ch1/fetchall]

- p.19 note the difference between the argument of the `handler` & the usage of the *
```go
// because the first argument will always be 1 instance it doens'nt
// need *, however there would be multiple requests which needs *
func handler(w http.ResponseWriter, r *http.Request) {fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)}
```

- p.22 local variable declaration to precede the if condition
```go
if err := r.ParseForm(); err != nil {
    log.Print(err)
}
// is the same as, but the past reduces the scope of the variable err
err := r.ParseForm()
if err != nil {
    log.Print(err)
}
```

- p.24 there is no pointer arithmetic


### Chapter 02: PROGRAM STRUCTURE


- p.32 please note that `&x` return the address of x and `*p` returns the value of the pointer

- p.40 define a **method** for a `named type`. this like the `str()` in Python which calls `__str__()` of the object.
```go
type Celsius float64
type Fahrenheit float64
const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius=0
    BoilingC   Celsius=100
)
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) 
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// the following method would be related to type celsius
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// so we could run the following
c := FToC(212.0)
fmt.Println(c.String()) // "100°C"
fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
fmt.Printf("%s\n", c)   // "100°C"
fmt.Println(c)          // "100°C"
fmt.Printf("%g\n", c)   // "100"; does not call String, %g floating point number
fmt.Println(float64(c)) // "100"; does not call String
```

### Chapter 03: Basic Data Types

- p.51 computers operate fundamentally on fixed-size numbers called **words**, which are interpreted as integers, floating-point numbers,bit sets, or memory addresses, then combined into larger aggregates that represent packets, pixels, portfolios, poetry, and everything else.

- p.51 Golang’s types fall into four categories: _basic types_, _aggregate types_, _reference types_, and _interface types_.

- p.52 **Basic Type** _rune_ is synonum to `int32`, _byte_ is sysnonum to `uint8`

- p.52 intro to _uintptr_ data-type

- p.69 intro to `DecodeRuneInString`


### Chapter 04: Composite Types
- p.81 composite types are `arrays`, `slices`, `maps`, and `structs`. Arrays & structs are **aggregate** types. aggregate value cannot contain itself.
- p.82 array length determination
```go
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q) // "[3]int"
```

- p.86 the difference between `array` declaration & `slice` declaration
```go
months := [...]string{1: "January", 2:"February", 3:"March", 4:"April", 5:"May", 6:"June", 7:"July", 8:"August", 9:"September", 10:"October", 11:"November", 12: "December"} // array
a := [...]int{0, 1, 2, 3, 4, 5} // array
b := [6]int{0, 1, 2, 3, 4, 5} // array
c := []int{0, 1, 2, 3, 4, 5} //slice
```

- p.86 Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices contain the same elements.

- p.87 There are two reasons why deep equivalence is problematic. First,unlike array elements, the elements of a slice are indirect, making it possible for a slice to contain itself. Although there are ways to deal with such cases, none is simple, efficient, and most importantly, obvious.

- p.88 allocating new `array` for a `slice` or growth strateg y
```go
 z = make([]int, zlen, zcap)
```
- p.90 note that slices are not **pure** reference types but resemble an aggregate type such as the following struct:
```go
type IntSlice struct {
    ptr   *int
    len, cap int
}
```
- p.91 using built-in `append` 
```go
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // append the slice x
fmt.Println(x)   // "[1 2 3 4 5 6 1 2 3 4 5 6]"
```

- p.91 `variadic` function: is a function of indefinite arity, i.e., one which accepts a variable number of arguments.
```go
func appendInt(x []int, y ...int) []int { // here y ... could be any number. which means variadic
    var z []int
    zlen := len(x) + len(y)
    // ...expand z to at least zlen...
    copy(z[len(x):], y)
    return z
}
```

- p.91-92 In-Place Slice Techniques: the `slices` in the following code share the same underlying array, however the contents of array are partly overwritten
```go
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = si++
        }
    }
    return strings[:i]
}
```

- p.94 a map element is not a **variable**, and we cannot take its address because growing a map might cause rehashing of existing elements into new storage locations, thus potentially invalidating the address, unlike Python:
```go
_ = &ages["bob"] // compile error: cannot take address of map element
```

- p.96 checking if the key exist in a `map`
```go
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }
//or as 
if age, ok := ages["bob"]; !ok { /* ... */ }
```

- p.100 The last statement updates the **Employee** `struct` that is pointed to by the result of the call to **EmployeeByID**. If the result type of **EmployeeByID** were changed to **Employee** instead of ***Employee**, the assignment statement would not compile since its left-hand side would not identify a variable.
```go
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

func EmployeeByID(id int) *Employee { /* ... */ }

```
