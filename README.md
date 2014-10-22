# scramble

`scramble` is a small utility that replaces all letters in a text file with random ASCII letters of the same case.

It is not meant to obfuscate a document, but to render it completely illegible while preserving word and character counts.

## Usage

`scramble` reads from standard input:

    $ cat lipsum.txt
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    $ scramble < lipsum.txt
    Zvngz ptmlg zxxgq egr yhjh, uctcuqfoacm nkpstroajc eqvq.
