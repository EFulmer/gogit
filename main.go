package main

import (
    "crypto/sha1"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        log.Fatal("usage: main [file-name]")
        os.Exit(1)
    }
    fileHash := HashObject(args[0])

    fmt.Println(fileHash)
}

// HashObject hashes the file located at filename, replicating the functionality of `git hash-object`.
func HashObject(filename string) string {
    fileContents, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    header := "blob " + fmt.Sprintf("%d", len(fileContents))
    header = header + "\u0000"

    store := header + fmt.Sprintf("%s", fileContents)

    hasher := sha1.New()
    io.WriteString(hasher, store)
    return fmt.Sprintf("%x", hasher.Sum(nil))
}
