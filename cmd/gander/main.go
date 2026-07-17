package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const usage = `gander - take a gander at your clouds
	Usage:
		gander <command> [flags]

	Commands:
		scan Scan configured providers and list resources
		version Print version

	Run "gander <command> -h" for command flags.
`

// version is set at build time via:
// go build -ldflags "-X main.version=0.1.0"
var version = "dev"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(2)
	}
	// cancel on Ctrl-C / SIGTERM. Every subcommand recive this
	// ctx; a second signal kills the process for good via
	// signal.NotifyContext default behavior being restored.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var err error
	switch cmd := os.Args[1]; cmd {
	case "scan":
		err = runScan(ctx, os.Args[2:])
	case "version":
		fmt.Println(version)
	case "-h", "--help", "help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "gander: unknow command %q\n\n%s", cmd, usage)
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "gander:", err)
		os.Exit(1)
	}
}

func runScan(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "emit resources as JSON lines")
	fs.Parse(args)
	_ = jsonOut
	return fmt.Errorf("no providers config found")
}
