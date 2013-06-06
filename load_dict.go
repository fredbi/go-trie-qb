package main

import (
    "./fileutil"
    "./trie"
    "flag"
    "fmt"
    "strings"
)

func LoadDict() {
    flag.Parse()
    if flag.NArg() < 1 {
        fmt.Println("Too few arguments");
        return
    }
    path := flag.Arg(0)

    trie := trie.NewTernaryTrie()
    count := 0
    for line := range fileutil.NewLineReadChannel(path) {
        trie.Add(strings.TrimSpace(line), count)
        count++
    }
    fmt.Printf("trie.Count=%d\n", trie.Count())
    fmt.Printf("trie.Find(\"あ\")=%s\n", trie.Find("あ"))
}

func Print(array []string) {
    switch (len(array)) {
    case 0:
        return
    case 1:
        fmt.Print(array[0])
        return
    case 2:
        fmt.Print(array[0])
        fmt.Print(array[1])
        return
    }
    mid := len(array) / 2

    fmt.Print(array[mid])
    Print(array[0:mid])
    Print(array[mid + 1:])
}

func Unsort() {
    flag.Parse()
    if flag.NArg() < 1 {
        fmt.Println("Too few arguments");
        return
    }
    path := flag.Arg(0)

    lines := []string{}
    for line := range fileutil.NewLineReadChannel(path) {
        lines = append(lines, line)
    }
    Print(lines[:])
}

func main() {
    LoadDict()
}
