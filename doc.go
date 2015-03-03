/*
Go implementation of reStructuredText

Usage example:

	package main

	import (
		"github.com/hhatto/gorst"
		"os"
		"bufio"
	)

	func main() {
		p := rst.NewParser(&rst.Extensions{Smart: true})

		w := bufio.NewWriter(os.Stdout)
		p.ReStructuredText(os.Stdin, rst.ToHTML(w))
		w.Flush()
	}
*/
package rst
