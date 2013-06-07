package main

import (
    "./fileutil"
    "./trie"
    "flag"
    "fmt"
    "strings"
    "time"
)

func LoadDict(path string) {
    trie := trie.NewTernaryTrie()
    count := 0
    start := time.Now()
    for line := range fileutil.NewLineReadChannel(path) {
        trie.Add(strings.TrimSpace(line), count)
        count++
    }
    elapsed := time.Since(start)
    fmt.Printf("elapsed=%f\n", elapsed.Seconds())
    //fmt.Printf("trie.Count=%d\n", trie.Count())
    //fmt.Printf("trie.Find(\"あ\")=%s\n", trie.Find("あ"))
}

func main() {
    flag.Parse()
    if flag.NArg() < 1 {
        fmt.Println("Too few arguments");
        return
    }
    path := flag.Arg(0)
    LoadDict(path)
}
