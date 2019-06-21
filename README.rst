***********************************
DRYcop Command-Line Interface (CLI)
***********************************

.. image:: https://img.shields.io/badge/license-Public%20Domain-blue.svg
   :alt: Project license
   :target: https://unlicense.org

.. image:: https://img.shields.io/badge/godoc-reference-blue.svg
   :alt: GoDoc reference
   :target: https://godoc.org/github.com/dryproject/drycop

.. image:: https://goreportcard.com/badge/github.com/dryproject/drycop
   :alt: Go Report Card score
   :target: https://goreportcard.com/report/github.com/dryproject/drycop

.. image:: https://img.shields.io/travis/dryproject/drycop/master.svg
   :alt: Travis CI build status
   :target: https://travis-ci.org/dryproject/drycop

|

Caveats
=======

*This is a semi-public, pre-alpha, work-in-progress project.*

**Here be dragons.**

**Caveat utilitor:** assume nothing works, and you may be pleasantly
surprised; and when it breaks, you get to keep both pieces.

Prerequisites
=============

To build the DRYcop CLI, you will need the following software:

- `GNU Make <https://www.gnu.org/software/make/>`__ 3.81+

- `Go <https://golang.org/>`__ 1.12+

Installation
============

::

   $ go get -u github.com/dryproject/drycop/drycop

Usage
=====

Usage on Travis CI
------------------

::

   $ cat .travis.yml

   # See: https://drycop.org
   language: go
   go:
     - 1.12
   git:
     depth: 1
   install:
     - go get -u github.com/dryproject/drycop/drycop
   script:
     - drycop init
     - drycop check -v
