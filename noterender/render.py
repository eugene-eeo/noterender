import sys
import traceback
import markdown2
from concurrent.futures import ProcessPoolExecutor
from path import Path
from .utils import fmt


EXTRAS = [
    'footnotes',
    'metadata',
    'tables',
    'smarty-pants',
    'markdown-in-html',
    'fenced-code-blocks',
]


def unpack_render_file(f_dst):
    file, dst = f_dst
    return render_file(file, dst)


def render_directory(src, dst):
    with ProcessPoolExecutor() as pool:
        pool.map(unpack_render_file, [
            (file, dst) for file in src.files('*.md')
            ])
        pool.shutdown()


def render_file(file, dst):
    name = file.basename().rsplit('.', 1)[0] + '.html'
    out  = Path(dst.joinpath(name))
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
