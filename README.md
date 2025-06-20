# changecase
* convert command line arguments to upper, lower or title case
* return the combined length of all command-line arguments
* check for string equality with optional case-insensitive matching
* outputs all content except a trailing newline, mimicking Perl's chomp functionality

## Synopsis

These programs convert all command line arguments to their respective case and
are also Unicode aware.

Provided programs:
* lower
* upper
* titlecase
* len
* eq
* chomp

## Usage

```shell
lower [arguments]
upper [arguments]
titlecase [arguments]
len [arguments]
eq [arguments]
chomp
(consider surrounding command-line arguments in double-quotes to preserve spacing)
```

## Installation

* macOS: `brew update; brew install jftuga/tap/changecase`
* Binaries for Linux, macOS and Windows are provided in the [releases](https://github.com/jftuga/changecase/releases) section.
