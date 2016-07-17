import sys
import traceback
import markdown2
from path import Path
from .utils import fmt


EXTRAS = [
    'footnotes',
    'metadata',
    'tables',
    'smarty-pants',
    'markdown-in-html',
]


def render_directory(src, dst):
    src = Path(src)
    dst = Path(dst)
    dst.mkdir_p()

    assert src.isdir()
    assert dst.isdir()

    for file in src.files('*.md'):
        name = file.basename()[:-3] + '.html'
        out = Path(dst.joinpath(name))
        out.write_text(try_render(file.text()))


def try_render(text):
    try:
        return render_text(text)
    except:
        etype, value, tb = sys.exc_info()
        msg = ''.join(traceback.format_exception(etype, value, tb))
        return render_error(msg)


def render_error(msg):
    return fmt('error.html', content=msg)


def render_text(text):
    md = markdown2.markdown(text, extras=EXTRAS)
    return fmt(
            'template.html',
            title=md.metadata['title'],
            content=md,
            )
