// Segitiga Terbalik
package main

import "fmt"
func cetakSegitiga(n int) {
    for i := n; i > 0; i-- {
        for j := 0; j < i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

func main() {
    var masukan int
    fmt.Scan(&masukan)
    cetakSegitiga(masukan)
}