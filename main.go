package main

import "fmt"
import "os"
import "io/ioutil"
import "flag"
import "github.com/hoisie/mustache"
import "github.com/fatih/color"

const defaultTemplate string = `<html>
<head><title>{{title}}</title></head>
<body>
{{{content}}}
</body>
</html>`

func renderPage(c *config, p *page, t *mustache.Template) string {
	d := map[string]string{}
	for k, v := range c.Params {
		d[k] = v
	}
	d["title"] = p.Title
	for k, v := range p.Params {
		d[k] = v
	}
	d["content"] = p.Render()
	return t.Render(d)
}

func check(err error) {
	if err != nil {
		fmt.Println(color.RedString("ERROR:"), err)
		os.Exit(1)
	}
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "nr-config.json", "configuration file")
	flag.Parse()

	println("Build started.")
	println("Reading config from", color.GreenString(configFile))

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		check(fmt.Errorf("cannot read config: %s", err))
	}
	config, err := newConfig(data)
	check(err)

	t, _ := mustache.ParseString(defaultTemplate)
	if config.Template != "" {
		println("Using template", color.GreenString(config.Template))
		t, err = mustache.ParseFile(config.Template)
		check(err)
	} else {
		println("Using default template")
	}

	check(os.MkdirAll(config.Src, 0777))
	check(os.MkdirAll(config.Dst, 0777))

	for _, p := range config.Files() {
		fmt.Println("Compiling", p.Path())
		page, err := newPageFrom(p.Path())
		check(err)
		ioutil.WriteFile(
			p.ChangeDir(config.Dst).ChangeExt("html").Path(),
			[]byte(renderPage(config, page, t)),
			0644,
		)
	}
	color.Green("OK")
}
