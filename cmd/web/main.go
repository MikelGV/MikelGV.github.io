package web

import (
	"context"
	"fmt"
	"os"

	"github.com/MikelGV/Personal-Website/internal/server"
)

func main() {
	ctx := context.Background()

	if err := server.Run(ctx, os.Stdout, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
