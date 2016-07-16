noterender
==========

.. image:: https://img.shields.io/badge/powered--by-oxygen-blue.svg?style=flat-square
.. image:: https://img.shields.io/badge/tests-none-green.svg?style=flat-square

Made for compiling my Markdown notes. Uses ``markdown2`` internally
and supports a richer superset of markdown with footnotes_, tables_,
metadata_, and SmartyPants_. Quick and dirty::

    $ mkdir build
    $ noterender --src='notes' --dst='build'

.. _footnotes:   https://github.com/trentm/python-markdown2/wiki/footnotes
.. _metadata:    https://github.com/trentm/python-markdown2/wiki/metadata
.. _tables:      https://github.com/trentm/python-markdown2/wiki/tables
.. _SmartyPants: http://daringfireball.net/projects/smartypants/
