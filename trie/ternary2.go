package trie

type Node2 struct {
    ch rune
    lo, eq, hi int
    Val interface {}
}

type Trie2 struct {
    nodes []*Node2
    count int
    countLo, countHi int
}

func New2() (*Trie2) {
    return &Trie2 {
        []*Node2 { &Node2 { 0, 0, 0, 0, nil } },
        0, 0, 0,
    }
}

func (t *Trie2) Dig(key string) (last *Node2) {
    index := &t.nodes[0].eq
    counter := &t.count
    for _, ch := range key {
        for {
            if *index == 0 {
                *index = len(t.nodes)
                t.nodes = append(t.nodes, &Node2 { ch, 0, 0, 0, nil })
                (*counter)++
            }

            node := t.nodes[*index]
            diff := ch - node.ch
            if diff == 0 {
                index = &node.eq
                counter = &t.count
                last = node
                break
            } else if diff < 0 {
                index = &node.lo
                counter = &t.countLo
            } else {
                index = &node.hi
                counter = &t.countHi
            }
        }
    }
    return last
}

func (t *Trie2) Add(key string, val interface {}) {
    t.Dig(key).Val = val
}

/*
func (t *Trie2) Find(key string) (node *Node) {
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
*/

/*
func (t *TernaryTrie) Count() int {
    fmt.Printf("Count: LO=%d EQ=%d HI=%d\n", t.countLo, t.count, t.countHi)
    return t.count
}
*/
