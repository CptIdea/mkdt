package generator

import (
	"os"
	"path/filepath"

	"github.com/CptIdea/mkdt/internal/cli"
	"github.com/CptIdea/mkdt/internal/parser"
)

type Options struct {
	DryRun bool
	Debug  bool
}

func Generate(root *parser.Node, opts Options) error {
	return generateNode(root, "", opts)
}

func generateNode(node *parser.Node, basePath string, opts Options) error {
	for _, child := range node.Children {
		fullPath := filepath.Join(basePath, child.Name)

		if opts.Debug {
			cli.PrintDebug("Processing: %s (dir: %v)", fullPath, child.IsDir)
		}

		if child.IsDir {
			if opts.DryRun {
				cli.PrintDryRun("Create directory: " + fullPath)
			} else {
				if err := os.MkdirAll(fullPath, 0755); err != nil {
					return err
				}
			}

			if err := generateNode(child, fullPath, opts); err != nil {
				return err
			}
		} else {
			if opts.DryRun {
				cli.PrintDryRun("Create file: " + fullPath)
			} else {
				if err := os.WriteFile(fullPath, []byte{}, 0644); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
