% SCRAMBLE(1) User Commands
%
% October 2014

# NAME

scramble - scramble words in place

# SYNOPSIS

**scramble** [*OPTIONS*] [*FILE*]...

# DESCRIPTION

Scramble words in each FILE in place and print the result to standard output. If FILE is not specified, or if FILE is -, read from standard input.

**scramble** preserves line, word, and character counts, as well as punctuation and the positions of uppercase letters.

# OPTIONS

-r, \--random
:   Replace letters with random ASCII letters, instead of scrambling words in place.

-h, \--help
:   Print usage message.

-V, \--version
:   Print the version.
