package rst

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func execFromString(rst string) (buf bytes.Buffer) {
	p := NewParser(nil)
	html := ToHTML(&buf)

	p.ReStructuredText(strings.NewReader(rst), html)

	return buf
}

func execFromFile(filename string) (buf bytes.Buffer) {
	input, err := os.Open(filename)
	if err != nil {
		log.Printf("%v", err)
	}
	defer input.Close()

	p := NewParser(nil)
	html := ToHTML(&buf)
	p.ReStructuredText(input, html)

	return buf
}

func TestExampleOfSimpleText(t *testing.T) {
	buf := execFromFile("data/simple.rst")
	ioutil.WriteFile("simple.html", buf.Bytes(), 0644)
}

func TestExampleOfPydebsignReadme(t *testing.T) {
	buf := execFromFile("data/pydebsign.readme.rst")
	ioutil.WriteFile("pydebsign.html", buf.Bytes(), 0644)

	output := buf.String()
	if !strings.Contains(output, "debsign is a command of devscripts that sign a Debian") {
		t.Errorf("invalid convert")
	}
}

func TestExampleOfMeowReadme(t *testing.T) {
	buf := execFromFile("data/meow.readme.rst")
	ioutil.WriteFile("meow.html", buf.Bytes(), 0644)
}

func TestExampleOfAutopep8Readme(t *testing.T) {
	buf := execFromFile("data/autopep8.readme.rst")
	ioutil.WriteFile("autopep8.html", buf.Bytes(), 0644)

	output := buf.String()
	if !strings.Contains(output, "<h1>Requirements</h1>") {
		t.Errorf("contain heading in autopep8")
	}
	if !strings.Contains(output, "be reported by pep8.</p>") {
		t.Errorf("para in autopep8")
	}
}

func TestExampleOfHeadingTitle(t *testing.T) {
	const input = `
=============
heading title
=============

title
-----
hoge1
		`
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<h1 class=\"title\">") {
		t.Errorf("not find '<h1 class=title>' string")
	}
}

func TestExampleOfHeading(t *testing.T) {
	const input = `
title
-----
hoge1

title2
======
hoge2
		`
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<h1>") {
		t.Errorf("not find '<h1>' string")
	}
	if !strings.Contains(output, "</h1>") {
		t.Errorf("not find '</h1>' string")
	}

	if !strings.Contains(output, "<h2>") {
		t.Errorf("not find '<h2>' string")
	}
	if !strings.Contains(output, "</h2>") {
		t.Errorf("not find '</h2>' string")
	}

	if !strings.Contains(output, "<p>hoge2</p>") {
		t.Errorf("paragraph error")
	}
}

func TestAutoLink(t *testing.T) {
	const input = "http://example.com/."
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href=\"http://example.com/\"") {
		t.Errorf("not find '<a>' string")
	}
	if !strings.Contains(output, "</a>") {
		t.Errorf("not find '</a>' string")
	}

	if !strings.Contains(output, ">http://example.com/<") {
		t.Errorf("paragraph error")
	}
}

func TestLinkContainsUnderbar(t *testing.T) {
	const input = "`test - _ -`_\n\n.. _`test - _ -`: http://example.com"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href") {
		t.Errorf("not find '<a>' string")
	}
	if !strings.Contains(output, "</a>") {
		t.Errorf("not find '</a>' string")
	}

	if !strings.Contains(output, "test - _ -") {
		t.Errorf("paragraph error")
	}
}

func TestUnquotedRefLinkUnderbarWithDot(t *testing.T) {
	const input = "this is LINK_.\n\n.. _LINK: http://this.is.link.com/\n\n"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href") {
		t.Errorf("not find '<a>' tag: %v", output)
	}
	if !strings.Contains(output, ">LINK</a>") {
		t.Errorf("invalid string: %v", output)
	}
}

func TestUnquotedRefLinkUnderbarWithDotAndList(t *testing.T) {
	const input = `
#. this is LINK_.
#. list 2

.. _LINK: http://this.is.link.com/
`
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href") {
		t.Errorf("not find '<a>' tag: %v", output)
	}
	if !strings.Contains(output, ">LINK</a>") {
		t.Errorf("invalid string: %v", output)
	}
}

func TestSimpleLinkRef(t *testing.T) {
	const input = "BBB_\n\n.. _BBB: http://example.com"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href") {
		t.Errorf("not find '<a>' string")
	}
	if !strings.Contains(output, "</a>") {
		t.Errorf("not find '</a>' string")
	}

	if !strings.Contains(output, "BBB") {
		t.Errorf("paragraph error")
	}
	if strings.Contains(output, "BBB_") {
		t.Errorf("paragraph error")
	}
}

func TestNotLinkRef(t *testing.T) {
	const input = "BBB_\n\nCCC_DDD"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "BBB</a>") {
		t.Errorf("invalid content: %v", output)
	}
	if !strings.Contains(output, "CCC_DDD</p>") {
		t.Errorf("invalid content: %v", output)
	}
}

func TestEmbeddedURI(t *testing.T) {
	const input = "`hoge <http://example.com>`_\n"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href=\"http://example.com\">hoge </a>") {
		t.Errorf("not find '<a>' string")
	}
}

func TestEmbeddedURIwithNewline(t *testing.T) {
	const input = "`hoge\nhoge <http://example.com>`_\n"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href=\"http://example.com\">hoge\nhoge </a>") {
		t.Errorf("not find '<a>' string")
	}
}

func TestEmbeddedAnonymouseURI(t *testing.T) {
	const input = "`hoge example <http://example.com>`__\n"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<a href=\"http://example.com\">hoge example </a>") {
		t.Errorf("not find '<a>' string")
	}
}

func TestImage(t *testing.T) {
	const input = "test\n\n.. image:: http://example.com/example.png\n"
	const input2result = "<p>test</p><img src=\"http://example.com/example.png\" alt=\"http://example.com/example.png\" />"
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid image tag. '%v'", output)
	}
}

func TestImageWithAlt(t *testing.T) {
	const input = `test

.. image:: http://example.com/example.png
   :alt: test text
`
	const input2result = `<p>test</p><img src="http://example.com/example.png" alt="test text" />`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid image tag. '%v'", output)
	}
}

func TestImageWithTarget(t *testing.T) {
	const input = `test

.. image:: http://example.com/example.png
    :target: http://example.com

`
	const input2result = `<p>test</p><a href="http://example.com"><img src="http://example.com/example.png" alt="http://example.com/example.png" /></a>`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid image tag. '%v' ... '%v", output, input2result)
	}
}

func TestImageWithAltAndTarget(t *testing.T) {
	const input = `test

.. image:: http://example.com/example.png
    :alt: test text
    :target: http://example.com/

`
	const input2result = `<p>test</p><a href="http://example.com/"><img src="http://example.com/example.png" alt="test text" /></a>`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid image tag. '%v' ... '%v", output, input2result)
	}
}

func TestGridTable(t *testing.T) {
	const input = `test

+--------+--------+
| hd1    | hd2    |
+========+========+
| bd1    | bd2    |
+--------+--------+
`
	const input2result = `
<table>

<thead>
<tr><td>hd1 </td>
<td> hd2 </td></tr>
</thead>

<tbody>
<tr><td>bd1 </td>
<td> bd2 </td></tr>
</tbody>
</table>`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid table tag. '%v' ... '%v", output, input2result)
	}
}

func TestHeaderLessGridTable(t *testing.T) {
	const input = `
+--------+--------+
| bd1-1  | bd2-1  |
+--------+--------+
| bd1-2  | bd2-2  |
+--------+--------+
`
	const input2result = `<table>

<tbody>
<tr><td>bd1-1 </td>
<td> bd2-1 </td></tr>
<tr><td>bd1-2 </td>
<td> bd2-2 </td></tr>
</tbody>
</table>`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid table tag. '%v' ... '%v", output, input2result)
	}
}

func TestApplicationDepent(t *testing.T) {
	const input = "hello ``code`` `world`\n"
	buf := execFromString(input)
	output := buf.String()
	if !strings.Contains(output, "<p>hello <code>code</code> world</p>") {
		t.Errorf("invalid. output=[%v]", output)
	}
}
