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

func TestExampleOfMeowReadme(t *testing.T) {
	buf := execFromFile("data/meow.readme.rst")
	ioutil.WriteFile("meow.html", buf.Bytes(), 0644)
}

func TestExampleOfAutopep8Readme(t *testing.T) {
	buf := execFromFile("data/autopep8.readme.rst")
	ioutil.WriteFile("autopep8.html", buf.Bytes(), 0644)

	output := buf.String()
	if !strings.Contains(output, "<h2>Requirements</h2>") {
		t.Errorf("contain heading in autopep8")
	}
	if !strings.Contains(output, "be reported by pep8.</p>") {
		t.Errorf("para in autopep8")
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
	t.Skip("still not support :target:")
	const input = `test

.. image:: http://example.com/example.png
    :target: http://example.com/

`
	const input2result = `<p>test</p><a href="http://example.com"><img src="http://example.com/example.png" alt="test text" /></a>`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid image tag. '%v' ... '%v", output, input2result)
	}
}

func TestImageWithAltAndTarget(t *testing.T) {
	t.Skip("still not support :target:")
	const input = `test

.. image:: http://example.com/example.png
    :alt: test text
    :target: http://example.com/

`
	const input2result = `<p>test</p><a href="http://example.com"><img src="http://example.com/example.png" alt="test text" /></a>`
	buf := execFromString(input)
	output := strings.TrimSpace(buf.String())
	if !strings.Contains(output, input2result) {
		t.Errorf("invalid image tag. '%v' ... '%v", output, input2result)
	}
}
