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

const katexBundle string = `
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.8.3/katex.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.8.3/katex.min.js"></script>
<script>
!function(){
	var set = document.getElementsByTagName('code');
	var pat = /(?:^\$\$\s*([\s\S]+)\s*\$\$$)|(?:^\$\s*([\s\S]+)\s*\$$)/;
	function extractTeX(text) {
		var match = text.match(pat);
		if (!match) return null;
		return {
			TeX: match[1] || match[2],
			displayMode: match[1] ? true : false,
		};
	}
	for (var i = set.length; i--;) {
		var code = set[i];
		var info = extractTeX(code.textContent);
		if (!info) continue;
		katex.render(info.TeX, code, {
			displayMode: info.displayMode,
			throwOnError: false,
		});
		code.classList.add('has-jax');
	}
}();
</script>
`

func renderPage(conf *config, pg *page, t *mustache.Template) string {
	d := map[string]string{}
	d["katex_bundle"] = katexBundle
	for k, v := range conf.Params {
		d[k] = v
	}
	d["title"] = pg.Title
	for k, v := range pg.Params {
		d[k] = v
	}
	d["content"] = pg.Render()
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
