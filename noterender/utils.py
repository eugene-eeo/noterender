import pkgutil
from pystache import render


def get_text(path):
    return pkgutil.get_data('noterender', 'data/' + path)\
                  .decode('utf-8')


def fmt(template, **kwargs):
    return render(template, kwargs)
