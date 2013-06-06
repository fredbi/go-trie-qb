package trie

import (
    "fmt"
)

type Node struct {
    ch rune
    lo, eq, hi *Node
    Val interface {}
}

type TernaryTrie struct {
    root *Node
    count int
    countLo, countHi int
}

func NewTernaryTrie() (*TernaryTrie) {
    return &TernaryTrie { nil, 0, 0, 0 }
}

func (t *TernaryTrie) Dig(key string) (last *Node) {
    p := &t.root
    counter := &t.count
    for _, ch := range key {
        for {
            if *p == nil {
                *p = &Node { ch, nil, nil, nil, nil }
                (*counter)++
                //t.count++
            }

            diff := ch - (*p).ch
            if diff == 0 {
                last = *p
                p = &(*p).eq
                counter = &t.count
                break
            } else if diff < 0 {
                p = &(*p).lo
                counter = &t.countLo
            } else {
                p = &(*p).hi
                counter = &t.countHi
            }
        }
    }
    return last
}

func (t *TernaryTrie) Add(key string, val interface {}) {
    t.Dig(key).Val = val
}

func (t *TernaryTrie) Find(key string) (node *Node) {
    p := t.root
    for _, ch := range key {
        for {
            if p == nil {
                node = nil
                break
            }

            diff := ch - p.ch
            if diff == 0 {
                node = p
                p = p.eq
                break
            } else if diff < 0 {
                p = p.lo
            } else {
                p = p.hi
            }
        }
    }
    return
}

func (t *TernaryTrie) Count() int {
    fmt.Printf("Count: LO=%d EQ=%d HI=%d\n", t.countLo, t.count, t.countHi)
    return t.count
}
