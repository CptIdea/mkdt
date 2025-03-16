package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/CptIdea/mkdt/internal/cli"
	"github.com/CptIdea/mkdt/internal/generator"
	"github.com/CptIdea/mkdt/internal/normalizer"
	"github.com/CptIdea/mkdt/internal/parser"
)

func main() {
	// Flags
	dryRun := flag.Bool("dry-run", false, "Preview changes without execution")
	debug := flag.Bool("debug", false, "Show parsing details")
	readClipboard := flag.Bool("c", false, "Read from clipboard")
	filePath := flag.String("f", "", "Path to template file")
	flag.Parse()

	var input string
	var err error

	if *readClipboard {
		input, err = readFromClipboard()
		if err != nil {
			cli.ErrorExit("Clipboard error: %v", err)
		}
	} else {
		input, err = readInput(*filePath)
		if err != nil {
			cli.ErrorExit("Input error: %v", err)
		}
	}

	normalized := normalizer.Normalize(input)

	// Parsing
	tree, err := parser.Parse(normalized)
	if err != nil {
		cli.ErrorExit("Parsing failed: %v", err)
	}

	// Generation
	opts := generator.Options{
		DryRun: *dryRun,
		Debug:  *debug,
	}

	if err := generator.Generate(tree, opts); err != nil {
		cli.ErrorExit("Generation failed: %v", err)
	}
}

func readInput(filePath string) (string, error) {
	if filePath != "" {
		data, err := os.ReadFile(filePath)
		return string(data), err
	}

	// Read from stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("reading stdin: %w", err)
	}
	return string(data), nil
}

func readFromClipboard() (string, error) {
	switch runtime.GOOS {
	case "windows":
		out, err := exec.Command("powershell", "-command", "Get-Clipboard").Output()
		return string(out), err
	case "darwin":
		out, err := exec.Command("pbpaste").Output()
		return string(out), err
	case "linux":
		out, err := exec.Command("xclip", "-out", "-selection", "clipboard").Output()
		return string(out), err
	default:
		return "", fmt.Errorf("unsupported platform")
	}
}
