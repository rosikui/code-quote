# code-quote

> Your terminal deserves more than just errors and logs — a fresh code quote for every session.

A CLI tool that displays inspiring programming quotes with beautiful terminal formatting.

## Features

- **Random Quotes**: Get a random programming quote every time
- **Daily Quotes**: Same quote for the same day (date-based seeding)
- **Tag Filtering**: Filter quotes by tags like `programming`, `clean-code`, `humor`
- **Language Filtering**: Filter by language (e.g., `en`, `zh`)
- **Beautiful Output**: ANSI colored output with bright purple formatting
- **Markdown Support**: Output in Markdown blockquote format
- **Custom Files**: Load additional quotes from JSON/YAML files
- **Flexible**: Support for multiple file formats and configurations

## Installation

### From Source

```bash
git clone https://github.com/rosikui/code-quote.git
cd code-quote
go build -o code-quote .
```

### Binary Download

Download the latest release from [GitHub Releases](https://github.com/rosikui/code-quote/releases).

### Auto-start Setup

Add `code-quote` to your terminal startup for daily inspiration:

#### Bash/Zsh (Linux/macOS)

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
# Add code-quote to your PATH (if not already installed globally)
export PATH="$HOME/path/to/code-quote:$PATH"

# Display a quote on terminal startup
code-quote --daily
```

#### PowerShell (Windows)

Add to your PowerShell profile (`$PROFILE`):

```powershell
# Add code-quote to your PATH (if not already installed globally)
$env:PATH += ";C:\path\to\code-quote"

# Display a quote on terminal startup
code-quote --daily
```

#### Fish Shell

Add to your `~/.config/fish/config.fish`:

```fish
# Add code-quote to your PATH (if not already installed globally)
set -gx PATH $HOME/path/to/code-quote $PATH

# Display a quote on terminal startup
code-quote --daily
```

#### VS Code Terminal

Add to your VS Code settings (`settings.json`):

```json
{
    "terminal.integrated.shellArgs.linux": ["-c", "code-quote --daily; exec bash"],
    "terminal.integrated.shellArgs.osx": ["-c", "code-quote --daily; exec zsh"],
    "terminal.integrated.shellArgs.windows": ["-c", "code-quote --daily; cmd"]
}
```

#### iTerm2 (macOS)

1. Go to **Preferences** → **Profiles** → **General**
2. In **Send text at start**, add:
   ```
   code-quote --daily
   ```

#### Windows Terminal

Add to your Windows Terminal settings (`settings.json`):

```json
{
    "profiles": {
        "list": [
            {
                "name": "PowerShell",
                "commandline": "powershell.exe",
                "startingDirectory": "%USERPROFILE%",
                "launchActions": [
                    {
                        "command": "code-quote --daily"
                    }
                ]
            }
        ]
    }
}
```

## Quick Start

```bash
# Build and run
make run

# Or build and run manually
go build -o code-quote .
./code-quote

# Set up auto-start (add to your shell profile)
echo 'code-quote --daily' >> ~/.zshrc  # or ~/.bashrc
```

## Usage

### Basic Usage

```bash
# Display a random quote
./code-quote

# Display daily quote (same quote for the same day)
./code-quote --daily

# Show tags with the quote
./code-quote --show-tags

# Output in Markdown format
./code-quote --markdown

# Disable colors
./code-quote --no-color
```

### Filtering

```bash
# Filter by tags
./code-quote --tag programming,clean-code

# Filter by language
./code-quote --lang en

# Combine filters
./code-quote --tag programming --lang en --daily
```

### Custom Quote Files

```bash
# Load additional quotes from a single file
./code-quote --file my-quotes.json

# Load from multiple files
./code-quote --file quotes1.json,quotes2.yaml

# Combine with built-in quotes
./code-quote --file custom.json --tag motivation
```

## Quote File Format

### JSON Format

```json
[
  {
    "text": "The best error message is the one that never shows up.",
    "author": "Thomas Fuchs",
    "tags": ["error-handling", "ux", "programming"],
    "lang": "en"
  }
]
```

### YAML Format

```yaml
- text: "Any fool can write code that a computer can understand. Good programmers write code that humans can understand."
  author: "Martin Fowler"
  tags: ["clean-code", "readability", "programming"]
  lang: "en"
```

### Required Fields

- `text`: The quote content (required)
- `author`: The quote author (required)
- `tags`: Array of tags (optional)
- `lang`: Language code (optional)

## CLI Options

| Flag | Description |
|------|-------------|
| `--file` | Additional quotes file(s) (JSON/YAML, comma-separated) |
| `--tag` | Filter by tags (comma-separated) |
| `--lang` | Filter by language |
| `--daily` | Daily quote (same quote for the same day) |
| `--no-color` | Disable ANSI colors |
| `--markdown` | Output in Markdown format |
| `--show-tags` | Show tags after author |

## Examples

```bash
# Get a random programming quote with tags
./code-quote --tag programming --show-tags

# Get today's quote in Markdown format
./code-quote --daily --markdown

# Load custom quotes and filter by clean-code tag
./code-quote --file my-quotes.json --tag clean-code

# Get a quote without colors for scripting
./code-quote --no-color --tag motivation
```

## Project Structure

```
code-quote/
├── main.go                 # Main entry point with embed
├── cmd/
│   └── code-quote/
│       └── main.go         # CLI logic with Cobra
├── internal/
│   ├── quote/
│   │   ├── loader.go       # JSON/YAML file loading
│   │   └── picker.go       # Quote selection and filtering
│   └── term/
│       └── render.go       # Terminal output formatting
├── quotes/
│   └── quotes.en.json      # Built-in quotes
└── README.md
```

## Development

### Prerequisites

- Go 1.24+

### Building

```bash
go build -o code-quote .
```

### Testing

```bash
go test ./...
```

## License

- Code: MIT License
- Quotes: CC-BY-4.0 (see `quotes/LICENSE`)

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for detailed information.

For questions and discussions, please use [Issue](https://github.com/rosikui/code-quote/issues).

## Built-in Quotes

The tool comes with a curated collection of programming quotes covering topics like:
- Clean code principles
- Problem solving
- Learning and practice
- Debugging
- Innovation
- And more!

Feel free to contribute new quotes by submitting a pull request or by using the `--file` flag with your own quote files.
