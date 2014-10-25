# scramble

`scramble` is a small utility that scrambles words in place.

`scramble` preserves line, word, and character counts, as well as punctuation and the positions of uppercase letters.

## Usage

`scramble` reads from standard input:

    $ cat lipsum.txt
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    $ scramble < lipsum.txt
    Roelm miups dlroo sti tmae, cctrtensoeu niagdcpisi eitl.

or from a list of files:

    $ scramble foo bar baz

Specify `-r`/`--random` to randomize the character replacements. This will render the document completely illegible.

    $ scramble -r lipsum.txt
    Iirka krfqq zeubz kgq khqz, dtovqdtdhue chyylzokja ulci.
