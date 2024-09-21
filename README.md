
# cmdify

`cmdify` is a simple Golang package that allows you to run shell commands, detect available shells, and execute common Linux commands programmatically in your Go applications. It provides utility functions for commonly used commands such as `ls`, `pwd`, `host`, `mkdir`, `touch`, and more.

## Installation

To install the package, use the following command:

```bash
go get github.com/cyinnove/cmdify
```

Then import it in your Go code:

```go
import "github.com/cyinnove/cmdify"
```

## Usage

### Running Custom Shell Commands

You can use the `RunCommand` or `RunCompinedCommand` functions to run shell commands and capture their output.

```go
package main

import (
    "fmt"
    "log"
    "github.com/cyinnove/cmdify"
)

func main() {
    output, err := cmdify.RunCompinedCommand("echo Hello World")
    if err != nil {
        log.Fatalf("Error running command: %v", err)
    }
    fmt.Println("Command Output:", output)
}
```

### Running Common Commands

#### List Files (`ls`)

```go
files, err := cmdify.Ls("-la")
if err != nil {
    log.Fatalf("Error running ls: %v", err)
}
fmt.Println("Files and directories:", files)
```

#### Get Current Directory (`pwd`)

```go
currentDir, err := cmdify.Pwd()
if err != nil {
    log.Fatalf("Error running pwd: %v", err)
}
fmt.Println("Current directory:", currentDir)
```

#### DNS Lookup (`host`)

```go
addresses, err := cmdify.Host("google.com")
if err != nil {
    log.Fatalf("Error running host: %v", err)
}
fmt.Println("Resolved addresses:", addresses)
```

#### Create a Directory (`mkdir`)

```go
err := cmdify.Mkdir("newdir")
if err != nil {
    log.Fatalf("Error creating directory: %v", err)
}
fmt.Println("Directory created successfully")
```

#### Create a File (`touch`)

```go
err := cmdify.Touch("newfile.txt")
if err != nil {
    log.Fatalf("Error creating file: %v", err)
}
fmt.Println("File created successfully")
```

### Detect Available Shells

The `DetectShells` function will return a list of available shells based on environment variables.

```go
shells := cmdify.DetectShells()
fmt.Println("Detected shells:", shells)
```



## Contributing

Feel free to open issues or submit pull requests to improve the package. Contributions are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
