// Programs entry point
package main

import (
    "log"
    "os"

    _ "github.com/Go_in_action_project/chapter2/matchers"
    "github.com/Go_in_action_project/chapter2/search"
)

func init() {
    // Change the device for logging to stdout
    log.SetOutput(os.Stdout)
}

func main() {
    // Perform the search for the specified term
    search.Run("president")
}