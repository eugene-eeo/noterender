import os
import pkgutil
from chevron import render


def get_text(path):
    return pkgutil.get_data('noterender', os.path.join('data/', path))\
                  .decode('utf-8')


def fmt(filename, **kwargs):
    kwargs.update({
        'stylesheet': get_text('stylesheet.css'),
        'script':     get_text('render-tex.js'),
    })
    return render(
            get_text(filename),
            kwargs,
            )
