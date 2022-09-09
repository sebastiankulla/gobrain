# gobrain
An Interpreter for the esoteric **Brainfuck** programming language written in golang.

## Usage
Write your **Brainfuck** code to a file and then have it interpreted in the following way:

```go
import (
    "https://github.com/sebastiankulla/gobrain"
    "fmt"
)

filename := "./hello_world.txt"
result := gobrain.interpret_file(filename)
fmt.Println(result)
```

