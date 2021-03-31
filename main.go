package main

import (
    "github.com/ksmt88/search-cloudwatchlogs/cmd"
)

func main() {
    input := cmd.Input{
        Profile: "",
    }
    cmd.Search(&input)
}
