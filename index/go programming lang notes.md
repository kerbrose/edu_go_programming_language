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
