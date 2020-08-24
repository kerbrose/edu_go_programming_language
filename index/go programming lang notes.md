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

