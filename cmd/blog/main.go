package main

import (
	"github.com/caijunduo/go-scaffold/internal/blog"
	"os"
)

func main() {
	if err := blog.New().Execute(); err != nil {
		os.Exit(1)
	}
}
