# scramble

`scramble` is a small utility that scrambles all the words in a text file.

## Usage

`scramble` reads from standard input:

    $ cat lipsum.txt
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    $ scramble < lipsum.txt
    Roelm miups dlroo sti tmae, cctrtensoeu niagdcpisi eitl.

or from a list of files:

    $ scramble foo bar baz

Specify `-r`/`--random` to completely randomize the character replacements. This will render the document illegible, but will maintain word and character counts.

    $ scramble -r lipsum.txt
    Iirka krfqq zeubz kgq khqz, dtovqdtdhue chyylzokja ulci.
