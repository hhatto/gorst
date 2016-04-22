===============================================
 pydebsign is a yet another library of debsign
===============================================

Status
------

.. image:: https://secure.travis-ci.org/mkouhei/pydebsign.png?branch=devel
   :target: http://travis-ci.org/mkouhei/pydebsign
.. image:: https://coveralls.io/repos/mkouhei/pydebsign/badge.png?branch=devel
   :target: https://coveralls.io/r/mkouhei/pydebsign?branch=devel
.. image:: https://pypip.in/v/pydebsign/badge.png
   :target: https://crate.io/packages/pydebsign
.. image:: https://readthedocs.org/projects/pydebsign/badge/?version=latest
   :target: https://readthedocs.org/projects/pydebsign/?badge=latest
   :alt: Documentation Status

Motivation
----------

`debsign` is a command of devscripts that sign a Debian .changes and .dsc file pare using GPG,
the command cannot use in environment witout TTY, for example invokeking by CI.

I had tried to use debsign from `subprocess` module of Python as follow,
but entering passphrase prompt is always returned.
It was the same in the case of using gnupg-agent and keyring.::

  >>> import subprocess
  >>> import shlex
  >>> command0 = 'echo -e "%s\n%s\n"' % (`passphrase`, `passphrase`)
  >>> command1 = '/usr/bin/debsign -k %s %s' % (`keyid`, `.changes`)
  >>> process0 = subprocess.Popen(shlex.split(command0),
  ...                             stdin=subprocess.PIPE,
  ...                             stdout=subprocess.PIPE,
  ...                             stderr=subprocess.PIPE)
  >>> process1 = subprocess.Popen(shlex.split(command1),
  ...                             stdin=process0.stdout,
  ...                             stdout=subprocess.PIPE,
  ...                             stderr=subprocess.PIPE)
  >>> stdout, stderr = process.communicate()

So, I decided to make a Python library to do the same behavior debsign.


Goal
----

* It is enable to sign `.changes` and `.dsc` files with GPG without the input of interactive passphrase.
* It can also be used by a user can not login shell on the CI, such as Jenkins.


Requires
--------

* Debian system, or the system derived from Debian.

  * Debian GNU/Linux Wheezy
  * Debian GNU/Linux Jessie/Sid
  * Ubuntu 14.04 LTS

* Debian package as follows;

  * gnupg
  * dput
  * lintian
  * python (= python2.7) or python3

* Python packages as follows;

  * python_gnupg (as debian package is python-gnupg or python3-gnupg)
  * python_debian (as debian package is python-debian or python3-debian)
  * chardet (as debian package is python-chardet or python3-chardet)


Usage
-----

Generic usage;::

  >>> from pydebsign import debsign
  >>> debsign.debsign_process('/path/to/some.changes', passphrase='secretkey')


When use another GPG Keyring instead of default GPG keyring;::

  >>> from pydebsign import debsign
  >>> debsign.debsign_process('/path/to/some.changes', passphrase='secretkey',
  ...                         keyid='keyid', gnupghome='/path/to/gpghome')

