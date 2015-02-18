gorst
=====
This is a Go_ implementation of reStructuredText_.

Only Support for HTML output is implemented.

.. _reStructuredText: http://docutils.sourceforge.net/docs/ref/rst/restructuredtext.html
.. _Go: http://golang.org/

**This is experimental module. Highly under development.**


Installation
------------
.. code-block:: bash

    $ go get github.com/hhatto/gorst


Usage
-----
.. code-block:: go

    package main

    import (
        "bufio"
        "os"
        "github.com/hhatto/gorst"
    )

    func main() {
        p := rst.NewParser(nil)

        w := bufio.NewWriter(os.Stdout)
        p.ReStructuredText(os.Stdin, rst.ToHTML(w))
        w.Flush()
    }


TODO
----
* Simple Table
* Footnotes
* Citations
* Directives (figure, contents, ...)
* etc...
