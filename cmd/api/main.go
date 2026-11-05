package main

import (
	"context"
	"fmt"
	"os"

	"github.com/MikelGV/Personal-Website/internal/server"
)

func main() {
	if err := server.Run(context.Background(), os.Stdout, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
