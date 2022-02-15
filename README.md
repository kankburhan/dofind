<h4 align="center">Dofind is fast domain related finder made with golang.</h4>

[![made-with-Go](https://img.shields.io/badge/made%20with-Go-green.svg)](http://golang.org)
[![issues](https://img.shields.io/github/issues/kankburhan/dofind?color=green)](https://github.com/kankburhan/dofind/issues)

```txt
	██████╗░░█████╗░███████╗██╗███╗░░██╗██████╗░
	██╔══██╗██╔══██╗██╔════╝██║████╗░██║██╔══██╗
	██║░░██║██║░░██║█████╗░░██║██╔██╗██║██║░░██║
	██║░░██║██║░░██║██╔══╝░░██║██║╚████║██║░░██║
	██████╔╝╚█████╔╝██║░░░░░██║██║░╚███║██████╔╝
	╚═════╝░░╚════╝░╚═╝░░░░░╚═╝╚═╝░░╚══╝╚═════╝░
	______________ by @kankburhan ______________

Usage:
	dofind [options]

Options:
	-t, --target <TARGET>		Target find domain
	-o, --output <FILE>		File to save results
	-c, --concurrent <i>		Set the concurrency level (default: 20)
	-s, --silent			Silent mode
	-v, --verbose			Verbose mode
	-u, --update-domain		Update domain list
	-h, --help			Display its help
```
---
## Resources

- [Installation](#installation)
  - [from Source](#from-source)
  - [from GitHub](#from-github)
- [Usage](#usage)
  - [Basic Usage](#basic-usage)
  - [Flags](#flags)
- [TODOs](#todos)
- [Help & Bugs](#help--bugs)
- [License](#license)
- [Version](#version)

## Installation

### from Source

go compiler:

```bash
▶ go install -v github.com/kankburhan/dofind/cmd/dofind@latest
```

### from GitHub

```bash
▶ git clone https://github.com/kankburhan/dofind
▶ cd dofind/cmd/dofind/
▶ go build .
▶ (sudo) mv dofind /usr/local/bin
```

## Usage

### Basic Usage

Run dofind with:

```bash
▶ dofind -t yahoo -o "/path/output"
```

### Flags

```bash
▶ dofind -h
```

This will display help for the tool.

```txt
	██████╗░░█████╗░███████╗██╗███╗░░██╗██████╗░
	██╔══██╗██╔══██╗██╔════╝██║████╗░██║██╔══██╗
	██║░░██║██║░░██║█████╗░░██║██╔██╗██║██║░░██║
	██║░░██║██║░░██║██╔══╝░░██║██║╚████║██║░░██║
	██████╔╝╚█████╔╝██║░░░░░██║██║░╚███║██████╔╝
	╚═════╝░░╚════╝░╚═╝░░░░░╚═╝╚═╝░░╚══╝╚═════╝░
	______________ by @kankburhan ______________

Usage:
	dofind [options]

Options:
	-t, --target <TARGET>		Target find domain
	-o, --output <FILE>		File to save results
	-c, --concurrent <i>		Set the concurrency level (default: 20)
	-s, --silent			Silent mode
	-v, --verbose			Verbose mode
	-u, --update-domain		Update domain list
	-h, --help			Display its help
```
## TODOs

- [ ] Validate parking domain
## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/kankburhan/dofind/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/kankburhan/dofind/issues).

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**dofind** released under MIT. See `LICENSE` for more details.

## Version

**Current version is 1.2.0** and still development.
