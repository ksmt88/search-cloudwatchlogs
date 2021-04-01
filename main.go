package main

import (
    "flag"

    "github.com/ksmt88/search-cloudwatchlogs/cmd"
)

func main() {
    profile := flag.String("profile", "", "profile name")
    flag.Parse()

    input := cmd.Input{
        Profile: *profile,
    }
    cmd.Search(&input)
}
