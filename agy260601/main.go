package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "xiaoqinli/codegen"
)

type AST interface{}

func main() {
    var inputPath string
    var target string
    flag.StringVar(&inputPath, "in", "", "Path to .xql.json input file (required)")
    flag.StringVar(&target, "target", "go", "Target language: go|rust|ts|py")
    flag.Parse()

    if inputPath == "" {
        fmt.Fprintln(os.Stderr, "error: -in flag is required")
        os.Exit(1)
    }

    data, err := ioutil.ReadFile(inputPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
        os.Exit(1)
    }
    var ast AST
    if err := json.Unmarshal(data, &ast); err != nil {
        fmt.Fprintf(os.Stderr, "invalid JSON AST: %v\n", err)
        os.Exit(1)
    }

    switch target {
    case "go":
        out, err := codegen.GenerateGo(ast)
        if err != nil {
            fmt.Fprintf(os.Stderr, "codegen error: %v\n", err)
            os.Exit(1)
        }
        fmt.Println(string(out))
    default:
        fmt.Fprintf(os.Stderr, "unsupported target: %s\n", target)
        os.Exit(1)
    }
}
