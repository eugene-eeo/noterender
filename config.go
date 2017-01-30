package main

import "fmt"
import "path"
import "io/ioutil"
import "encoding/json"
import "strings"

type config struct {
	Template string            `json:"template"`
	Src      string            `json:"src"`
	Dst      string            `json:"dst"`
	Params   map[string]string `json:"params"`
}

type pair struct {
	Dir  string
	Base string
}

func newConfig(data []byte) (*config, error) {
	c := new(config)
	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("cannot read json: %s", err)
	}
	if c.Src == "" {
		return nil, fmt.Errorf("src field not specified")
	}
	if c.Dst == "" {
		return nil, fmt.Errorf("dst field not specified")
	}
	return c, nil
}

func (c *config) Files() []pair {
	files, err := ioutil.ReadDir(c.Src)
	if err != nil {
		return []pair{}
	}
	b := []pair{}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			b = append(b, pair{c.Src, f.Name()})
		}
	}
	return b
}

func (p pair) ChangeDir(d string) pair {
	return pair{d, p.Base}
}

func (p pair) ChangeExt(ext string) pair {
	l := strings.Split(p.Base, ".")
	last := l[len(l)-1]
	return pair{
		p.Dir,
		strings.TrimSuffix(p.Base, last) + ext,
	}
}

func (p pair) Path() string {
	return path.Join(p.Dir, p.Base)
}
