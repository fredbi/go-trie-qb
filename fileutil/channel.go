package fileutil

import (
    "bufio"
    "io"
    "log"
    "os"
)

func NewLineReadChannel(path string) (c chan string) {
    c = make(chan string)

    go func() {
        defer close(c)

        file, err := os.Open(path)
        if err != nil {
            log.Fatal(err)
            return
        }
        defer file.Close()

        r := bufio.NewReader(file)
        for {
            line, err := r.ReadString('\n')
            if err != nil {
                if err == io.EOF {
                    break
                } else {
                    log.Fatal(err)
                    return
                }
            }
            c <- line
        }
    }()
    return
}
