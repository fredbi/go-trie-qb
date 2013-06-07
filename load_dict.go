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

func main() {
    LoadDict()
}
