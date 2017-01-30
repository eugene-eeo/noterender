package main

import "fmt"
import "io/ioutil"
import "github.com/ericaro/frontmatter"
import bf "github.com/russross/blackfriday"

const htmlFlags int = 0 |
	bf.HTML_USE_XHTML |
	bf.HTML_USE_SMARTYPANTS |
	bf.HTML_SMARTYPANTS_LATEX_DASHES |
	bf.HTML_FOOTNOTE_RETURN_LINKS

const mdExtensions int = 0 |
	bf.EXTENSION_TABLES |
	bf.EXTENSION_FENCED_CODE |
	bf.EXTENSION_FOOTNOTES |
	bf.EXTENSION_STRIKETHROUGH |
	bf.EXTENSION_HEADER_IDS |
	bf.EXTENSION_DEFINITION_LISTS |
	bf.EXTENSION_AUTOLINK

type page struct {
	Title   string            `yaml:"title"`
	Params  map[string]string `yaml:"params,inline"`
	Content string            `fm:"content" yaml:"-"`
}

func newPageFrom(filename string) (*page, error) {
	p := new(page)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error when reading file \"%s\": %s", filename, err)
	}
	err = frontmatter.Unmarshal(data, p)
	if err != nil {
		return nil, fmt.Errorf("cannot parse frontmatter: %s", err)
	}
	return p, nil
}

func (p *page) Render() string {
	return string(bf.Markdown(
		[]byte(p.Content),
		bf.HtmlRenderer(htmlFlags, "", ""),
		mdExtensions,
	))
}
