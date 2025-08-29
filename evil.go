package main

import (
    "fmt"
    "os"
)

func init() {
    data, err := os.ReadFile("/flag.txt")
    if err != nil {
        fmt.Println("[!] Error reading flag:", err)
        return
    }
    fmt.Println("FLAG:", string(data))
}

func main() {}
