// cmd/blockchainnftregistryenginepro/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistryenginepro/internal/blockchainnftregistryenginepro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistryenginepro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
