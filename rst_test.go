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
