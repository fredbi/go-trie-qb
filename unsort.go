package main

import (
    "./fileutil"
    "container/list"
    "flag"
    "fmt"
)

func Output(queue *list.List) {
    for ; queue.Len() != 0; {
        target := (queue.Front().Value).([]string)
        queue.Remove(queue.Front())

        switch (len(target)) {
        case 0:
            continue
        case 1:
            fmt.Print(target[0])
            continue
        case 2:
            fmt.Print(target[0])
            fmt.Print(target[1])
            continue
        }

        mid := len(target) / 2
        fmt.Print(target[mid])
        queue.PushBack(target[0:mid])
        queue.PushBack(target[mid + 1:])
    }
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

    queue := list.New()
    queue.PushBack(lines)
    Output(queue)
}

func main() {
    Unsort()
}
