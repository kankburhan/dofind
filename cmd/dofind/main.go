package main

import "github.com/kankburhan/dofind/internal/runner"

func main() {
	option := runner.ParseOptions()
	runner.New(option)
}
