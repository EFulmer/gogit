package main

import (
    "os/exec"
    "strings"
    "testing"
)

func TestHashObject(t *testing.T) {
    FILE := "README.md"
    ourHash := HashObject(FILE)
    gitHash := exec.Command("git", "hash-object", FILE)
    var out strings.Builder
    gitHash.Stdout = &out
    err := gitHash.Run()
    if err != nil {
        t.Errorf("git hash-object failed somehow") // TODO better error message.
    }
    if strings.TrimSpace(out.String()) != ourHash {
        t.Errorf("hashes did not match")
    }
}
