package main

import (
    "os"
    "strconv"
    "fmt"
    "./tempconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "tempconv-test: %v\n", err)
            os.Exit(1)
        }

        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)

        fmt.Printf("%s = %s, %s = %s\n", f, tempconv.F2C(f), c, tempconv.C2F(c))
    }
}
