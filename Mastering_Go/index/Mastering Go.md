# Mastering Go

### Chapter 01: Go and the Operating System

- p.9 big changes that are being considered for **Go 2** are `generics`, `package versioning`, and `improved error handling`.

- p.15 bypassing restriction of importing unused package
```go
import (
    "fmt"
    _ "os"    // add _ before the package
)
```