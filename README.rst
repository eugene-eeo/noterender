noterender
==========

.. image:: https://img.shields.io/badge/powered--by-oxygen-blue.svg?style=flat-square
.. image:: https://img.shields.io/badge/tests-none-green.svg?style=flat-square


Made for compiling my Markdown notes. Notes can be written in a
richer superset of markdown with footnotes_, tables_, metadata_,
and SmartyPants_. Also you can include math in the form of LaTeX.
Quick and dirty::

    $ noterender --src='notes/' --dst='build/'

Because this is largely made for my own use, it requires that you
already have Open Sans on your machine for the rendered notes to
look acceptable. Internally this project hitchhikes KaTeX_, markdown2_,
docopt_, chevron_, and path.py_.


.. _footnotes:   https://github.com/trentm/python-markdown2/wiki/footnotes
.. _metadata:    https://github.com/trentm/python-markdown2/wiki/metadata
.. _tables:      https://github.com/trentm/python-markdown2/wiki/tables
.. _SmartyPants: http://daringfireball.net/projects/smartypants/

.. _KaTeX:     https://github.com/Khan/KaTeX
.. _markdown2: https://github.com/trentm/python-markdown2
.. _docopt:    https://github.com/docopt/docopt
.. _chevron:   https://github.com/noahmorrison/chevron
.. _path.py:   https://github.com/jaraco/path.py
