# mkdt  
**make directory tree**  

A minimalist CLI tool to generate directory structures from text input. Stop creating folders manually — describe your layout and let `mkdt` build it instantly.

## Installation

```shell
go install github.com/CptIdea/mkdt@latest
```

## Description
`mkdt` parses plain text descriptions of file structures and creates matching directories/files. Perfect for:
- Quickly scaffolding projects
- Replicating folder hierarchies from docs/chat
- Automating repetitive directory creation

**Supports:**  
✅ Nested folders/files  
✅ Comments in templates  
✅ "..."-like placeholders  
✅ Parsing decorators  
✅ Dry-run mode (preview changes)  
✅ Clipboard input

### Example Input Formats
```text
# Simple structure
project/
  src/
    main.go
  .gitignore
```

```text
# With ASCII art
app/
├── config/
│   └── settings.yaml
└── scripts/
    └── deploy.sh
```

## Usage

#### From stdin:
```shell
echo '
src/
  main.go
  internal/
    utils.go  # Helper functions
' | mkdt
```

#### From clipboard (*requires `xclip`/`pbpaste`*):
```shell
mkdt -c
```

#### From a file:
```shell
mkdt -f template.txt
```

#### Dry-run (preview):
```shell
cat structure.txt | mkdt generate --dry-run
```

#### Full options:
```shell
mkdt --help
```
