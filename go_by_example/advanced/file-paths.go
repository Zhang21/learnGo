package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// The `filepath` package provides functions to parse and construct file paths
// in a way that is portable between operating systems.
// Attention Unix-like(dir/file) and windows(dir\file)

func main() {
	// `Join` should used to construct paths in a portable way.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// In addition to providing portability, `Join` will also normalize paths by
	// removing superfluous seperators and directory changes.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` gets dir name, `Base` gets filename
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p:)", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// `Ext` gets extension out of such names.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))
	fmt.Println(ext)

	// `Rel` finds a relative path between a base and a target.
	// It returns an error if the target cannot be made relative to base.
	// Join(base, Rel(base, target)) eq target
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
