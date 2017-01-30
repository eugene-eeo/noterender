---
title: Using noterender for fun and profit
version: 0.1-beta
---

`noterender` is a program that can be run to transform a directory of Markdown
files to HTML files in another directory. First create a new directory with
the following structure:

    noterender-config.json
    template.html
    doc/
        file1.md

Copy and paste the following into `nr-config.json`:

    {
        "src": "doc/",
        "dst": "out/"
    }

The `src` and `dst` fields tell `noterender` where to search for the Markdown
files (it only reads files ending with `.md`) and where to dump the rendered
HTML, respectively. In `file1.md`:

    ---
    title: My First File
    ---

    my first text

Run the following in your terminal:

    $ ./noterender

You should see a bunch of text and a green "OK" at the end if everything
goes well. You should see an `out/` directory created that contains
`file1.html`. When you open it, you should notice the following:

 - anything in the first few lines bounded by `---` are not rendered.
 - anything below the `---` is rendered and shown on the page.
 - the title of the page (the 'label' of the current tab) corresponds
 to the title specified in the frontmatter (the block of text between
 the `---` bars).

When you write markdown files for `noterender` you must follow following structure:

    ---
    title: its a required value
    key: value
    ---

    Markdown content.

You can also mix HTML tags with Markdown content. For instance, to create
a table in the page we can use either of the following:

    | Tables        | Are           | Cool  |
    | ------------- |:-------------:| -----:|
    | col 3 is      | right-aligned | $1600 |
    | col 2 is      | centered      |   $12 |
    | zebra stripes | are neat      |    $1 |

    or

    <table>
    <th>
        <td>Tables</td>
        <td>Are</td>
        <td>Cool</td>
    </th>
    <tr>
        ...
    </tr>
    </table>

Refer to https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet
for more information on how to write Markdown. For more control over the
HTML output, you can use templates in order to be able to use custom
stylesheets. In your `nr-config.json` add a `template` field:

    {
        ...,
        "template": "template.html"
    }

And then in `template.html`:

    <html>
    <head>
        <title>{{title}}</title>
    </head>
    <body>
        {{{content}}}
    </body>
    </html>

When `noterender` runs the file specified by `template` is used as a
mustache[^1] template. `{{title}}` and `{{{content}}}` will contain
the title of the page and the rendered markdown content, respectively.
**It is important that you use `{{{content}}}` instead of `{{content}}`**
because in the latter case, the HTML will be escaped.

Another gimmick: you can use custom parameters in your markdown files.
For example, say you have the following template:

    <html>
    <head>
        <title>{{title}}</title>
    </head>
    <body>
        {{{content}}}
        <hr/>
        Written by: {{author}}
    </body>
    </html>

In your markdown file you can write the following:

    ---
    title: My Title
    author: John Doe
    ---

    Something

To which you'll see the following output:

    <html>
    <head>
        <title>My Title</title>
    </head>
    <body>
        <p>Something
        <hr/>
        Written by: John Doe
    </body>
    </html>

## FAQ

### Displaying math?

Adapt the following template (uses KaTeX):

    <!doctype html>
    <html>
    <head>
      <meta charset='utf-8'/>
      <style>
          ...
      </style>
      <title>{{title}}</title>
    </head>
    <body>
      {{{content}}}
      {{{katex_bundle}}}
    </body>
    </html>

In your markdown files, when you want to write math wrap it in
code blocks, for instance:

| In                                      | Out           |
| --------------------------------------- |---------------|
| <code>&grave;\$ \KaTeX \$&grave;</code> | `$ \KaTeX $`  |
| <code>&grave;\$\$ \frac{n}{2} \$\$&grave;</code> | `$$ \frac{n}{2} $$`  |

The code only activates `$ \KaTeX $` when the code block
consists entirely of a math expression in the form of <code>&grave;$ ... $&grave;</code>
for inline math, or <code>&grave;$$ ... $$&grave;</code> for display
style (center-aligned) math. This is inline: `$ \frac{n}{2} $` and
this is display style:

`$$
\frac{n}{2}
$$`

### Footnotes

```
some text[^ref]

[^ref]: Lorem Ipsum...
```

[^1]: https://mustache.github.io/mustache.5.html