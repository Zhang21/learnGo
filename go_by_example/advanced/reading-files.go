package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Reading and writing files are basic tasks needed for many go programs.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Sluping a file's entire contents into memory.
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// You'll often want more control over how and what parts of a file are read.
	// Starting by `Open` a file to obtain an `os.File` value.
	f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// `Seek` to a known location in the file and `Read` from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	// The `bufio` implements a buffered reader that may be useful both for its efficiency
	// with small reads and beacuse of the additional reading methods it provides.
	// io操作本身的效率并不低，低的是频繁的访问本地磁盘的文件。所以bufio就提供了缓冲区(分配一块内存)，
	// 读和写都先在缓冲区中，最后再读写文件，来降低访问本地磁盘的次数，从而提高效率。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close  the file when you've done.
	// Usually this would be scheduled immediately after `Open` with `defer`.
	f.Close()
}
