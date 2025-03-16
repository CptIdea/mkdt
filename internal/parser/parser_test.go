package parser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    *Node
		wantErr bool
		errText string
	}{
		{
			name:  "Empty input",
			input: []string{},
			want:  &Node{Depth: -1, IsDir: true},
		},
		{
			name:  "Single directory",
			input: []string{"dir/"},
			want: &Node{
				Depth: -1,
				IsDir: true,
				Children: []*Node{
					{Name: "dir", IsDir: true, Depth: 0},
				},
			},
		},
		{
			name:  "Single file",
			input: []string{"file.txt"},
			want: &Node{
				Depth: -1,
				IsDir: true,
				Children: []*Node{
					{Name: "file.txt", IsDir: false, Depth: 0},
				},
			},
		},
		{
			name: "Nested structure",
			input: []string{
				"root_dir/",
				"\tsubdir/",
				"\t\tfile1.txt",
				"\tfile2.txt",
			},
			want: &Node{
				Depth: -1,
				IsDir: true,
				Children: []*Node{
					{
						Name:  "root_dir",
						IsDir: true,
						Depth: 0,
						Children: []*Node{
							{
								Name:  "subdir",
								IsDir: true,
								Depth: 1,
								Children: []*Node{
									{Name: "file1.txt", Depth: 2},
								},
							},
							{Name: "file2.txt", Depth: 1},
						},
					},
				},
			},
		},
		{
			name: "Mixed dirs/files",
			input: []string{
				"a/",
				"\tb/",
				"\t\tc.txt",
				"\td.txt",
				"e.txt",
			},
			want: &Node{
				Depth: -1,
				IsDir: true,
				Children: []*Node{
					{
						Name:  "a",
						IsDir: true,
						Depth: 0,
						Children: []*Node{
							{
								Name:  "b",
								IsDir: true,
								Depth: 1,
								Children: []*Node{
									{Name: "c.txt", Depth: 2},
								},
							},
							{Name: "d.txt", Depth: 1},
						},
					},
					{Name: "e.txt", Depth: 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)

			// Проверка ошибок
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if err != nil && tt.errText != "" && err.Error() != tt.errText {
				t.Fatalf("Parse() error = %v, want error text = %q", err, tt.errText)
			}

			if tt.wantErr {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Parse():\n%s", diffTrees(got, tt.want))
			}
		})
	}
}

func diffTrees(got, want *Node) string {
	return fmt.Sprintf(
		"Got:\n%s\nWant:\n%s\n",
		treeToString(got, 0),
		treeToString(want, 0),
	)
}

func treeToString(n *Node, level int) string {
	if n == nil {
		return ""
	}

	indent := strings.Repeat("  ", level)
	result := fmt.Sprintf("%s%s (d=%d, dir=%v)\n",
		indent,
		n.Name,
		n.Depth,
		n.IsDir,
	)

	for _, child := range n.Children {
		result += treeToString(child, level+1)
	}
	return result
}
