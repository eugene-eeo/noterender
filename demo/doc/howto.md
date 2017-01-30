---
title: Using noterender for fun and profit
version: 0.1-beta
---

`noterender`
: Cross platform program that takes a bunch of Markdown files in a
: directory and then spits out a bunch of HTML files into another
: directory.

## Up and running

Set up a new folder with the following structure:

    nr-config.json
    template.html
    doc/
        hello-world.md

In `nr-config.json` – this is the file where `noterender` will read
and parse when you run the program in a folder. It is where you can
configure the tool. The contents of the file are parsed as [JSON](http://www.json.org).

    {
        "src": "doc/",
        "dst": "out/",
        "template": "template.html"
    }

In `template.html` – this is a file which will be used as a template
for your HTML output. It will be treated as a [mustache](https://mustache.github.io/mustache.5.html)
template. The important things to include in your template are:


 - `{{title}}` will be replaced with the title of the document
 - `{{{content}}}` will be replaced with the rendered HTML Markdown content
 (note the triple braces, these are important and necessary).
 - `{{{katex_bundle}}}` needed for rendering math equations

```
<html>
<head>
    <title>{{title}}</title>
</head>
<body>
    {{{content}}}
    {{{katex_bundle}}}
</body>
</html>
```

In `hello-world.md` – this file will be read by `noterender` when
you run the program. It is important that it lives in the `doc/`
directory because it is where we specified our source files (see
the `src` option in `nr-config.json`). It will then be rendered
into `out/hello-world.html` when we execute the program. For a
very quick tutorial to Markdown you should read `adam-p`'s
excellent [Markdown Cheatsheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet).

```
---
title: My First Document
---

Some **Markdown** content. Some `$ \KaTeX $` rendered _inline_
math content: `$ \frac{n}{2} $`; and then some display-style
math content:

`$$
1 + 1 = 2
$$`
```

Now that we have everything in place, run the `noterender`
program at the shell. Assuming you have the `noterender.exe`
binary in your current path and your current path is the
folder that you've created at the start (don't type the `> `
part):

```
> noterender.exe
```

You should see some informational messages and a green "OK"
at the very end. Now open `out/hello-world.html` in a browser
and you should see something like (but not exactly, because
our template does not contain any CSS styling):

----------

# My First Document

Some **Markdown** content. Some `$ \KaTeX $` rendered _inline_
math content: `$ \frac{n}{2} $`; and then some display-style
math content:

`$$
1 + 1 = 2
$$`

----------

Now just write your custom stylesheet in the `template.html` file
and you'll be good to go.

## More Examples

### Footnotes

```
Lorem ipsum[^ref].

[^ref]: http://www.lipsum.com
```

Scroll to the bottom of the page to see the footnote being rendered.
Footnotes automatically come with return links. Note that the footnote
need not contain only links and can be placed anywhere in the document,
not just at the bottom.

--------

Lorem ipsum[^ref].

[^ref]: http://www.lipsum.com

--------

### Math Equations

```
Inline math: `$ \frac{n}{2} $` (use it in paragraphs), for example when
explaining a proof. Use display style (center aligned, looks bigger) for
longer or important equations:

`$$
\lim_{n \to \infty}{\sum_{i=1}^{n}{\frac{1}{2^i}}} = 1
$$`
```

------------

Inline math: `$ \frac{n}{2} $` (use it in paragraphs), for example when
explaining a proof. Display style (center aligned, looks bigger).

`$$
\lim_{n \to \infty}{\sum_{i=1}^{n}{\frac{1}{2^i}}} = 1
$$`

------------

### Custom Parameters

Templates can contain custom parameters. For instance if you want to
include a chapter number with custom styling at the bottom of the page,
you need copy and paste the HTML tags in every Markdown document. Say
we have the following template:

```
<html>
<head>
    <title>{{title}}</title>
</head>
<body>
    {{{content}}}
    <span class='chapterno'>{{chapterno}}</span>
</body>
</html>
```

And the following Markdown document:

```
---
title: Introduction to Harry Potter
chapterno: 1
---

Harry Potter is...
```

The rendered file would look like:

```
<html>
<head>
    <title>Introduction to Harry Potter</title>
</head>
<body>
    <p>Harry Potter is...
    <span class='chapterno'>1</span>
</body>
</html>
```

What if, for instance, most of the pages have the same value for
a parameter but some of them don't? We don't want to repeat the
parameters again and again so we can define a params "mapping"
in our configuration file:

```
{
    ...,
    "params": {
        "chapterno": "1"
    }
}
```

So all pages will now have the `{{chapterno}}` as `1` by default,
but can be changed in the metadata on top of the page (the block
of text between the two `---` bars). It is important to note that
**while custom parameters are not required and can be omitted, the
`title` parameter is required.**

### KaTeX Bundle

`katex_bundle` is a special parameter that defaults to some HTML tags
that will include the KaTeX JavaScript and CSS files in the rendered
HTML, and a custom script that will help find the math expressions for
KaTeX. It needs to be at the _bottom_ of the page, right before the
closing `</body>` tag.
