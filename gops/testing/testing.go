package main

import (
    "log"
    "time"

    "github.com/google/gops/agent"
)

func main() {
    if err := agent.Listen(agent.Options{}); err != nil {
        log.Fatal(err)
    }
    time.Sleep(time.Hour)
}
