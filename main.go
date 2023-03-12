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
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    hasher := sha1.New()

    if _, err := io.Copy(hasher, file); err != nil {
        log.Fatal(err)
    }

    hash := fmt.Sprintf("%x", hasher.Sum(nil))

    return hash
}
