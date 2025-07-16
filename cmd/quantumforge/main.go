// cmd/quantumforge/main.go
package main

import (
"flag"
"log"
"os"

"quantumforge/internal/quantumforge"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumforge.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
