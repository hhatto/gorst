====
Meow
====

About
=====

Meow is an editor-agnostic Markdown, reStructuredText and Textile live previewer.
Once you save your file, the rendered HTML will automatically get reloaded
in your favourite browser.

The Original Version is moo_ . Meow fork from moo ( `commit history`_ ), And add function of
reStructuredText live previewer.

.. _moo: https://github.com/pyrocat101/moo
.. _`commit history`: https://github.com/pyrocat101/moo/commits/876de66f792ac42df7dba2e1416fb7aeae8feeb2

Installation
============

.. code-block:: bash

    pip install meow

Requirements
============

meow requires

* mistune_ (optional)
* `python-textile`_ (optional)
* `bottle.py`_
* pygments_
* docutils_
* CherryPy_
* docopt_

.. _mistune: https://github.com/lepture/mistune
.. _`python-textile`: https://github.com/sebix/python-textile
.. _`bottle.py`: http://bottlepy.org/
.. _pygments: http://pygments.org/
.. _docutils: https://pypi.python.org/pypi/docutils
.. _CherryPy: http://www.cherrypy.org/
.. _docopt: https://github.com/docopt/docopt

Usage
=====

Opens preview in browser with server listening on 3000::

    meow --port 3000 your-doc.markdown

Export to HTML only::

    meow -o exported.html your-doc.markdown

To specify file type::

    meow --filetype rst README

RESTful API
===========

 +-------------------------+-------------+----------------------------------+-------------------------+
 | Action                  | HTTP Method | Request URL                      | Response Body           |
 +=========================+=============+==================================+=========================+
 | Get preview             | GET         | http://localhost:\<port\>        | \<Preview content\>     |
 +-------------------------+-------------+----------------------------------+-------------------------+
 | Get updated content     | POST        | http://localhost:\<port\>/update | \<Rendered body\>       |
 +-------------------------+-------------+----------------------------------+-------------------------+
 | Close server            | DELETE      | http://localhost:\<port\>        |                         |
 +-------------------------+-------------+----------------------------------+-------------------------+

See `source files`_ for more details.

.. _`source files`: https://github.com/hhatto/meow

License
=======

(The MIT License)

Links
=====
* GitHub_
* PyPI_

.. _GitHub: https://github.com/hhatto/meow
.. _PyPI: https://pypi.python.org/pypi/meow/
