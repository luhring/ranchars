# Ranchars

[![Build Status](https://circleci.com/gh/luhring/ranchars.svg?style=shield)](https://circleci.com/gh/luhring/ranchars)
[![Go Report Card](https://goreportcard.com/badge/github.com/luhring/ranchars)](https://goreportcard.com/report/github.com/luhring/ranchars)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Ran(dom) char(acter)s

Ranchars generates a string of [random](##randomness) characters that satisfies a set of specified criteria, of a specified length.

## Basic usage

The `ranchars` command always takes at least one argument to specify how many character long the random string should be.

```
$ ranchars 16
8rvSlZxZRhvImQGM
```

## Character types

Ranchars generates characters in terms of character types. There are five types of characters used:

1. Numeric digits (0-9)
1. Lowercase letters (a-z)
1. Uppercase letters (A-Z)
1. Special characters (set defined in [Special characters](###-special-characters))
1. The space character

By default, Ranchars uses numeric digits, lowercase letters, and uppercase letters. Strings composed of these types are easy to select in most text environments, since they aren't broken into chunks by special characters or spaces.

But, you can tell Ranchars to ensure that specified character types are each represented by at least one character in the generated string.

For example, to generate a random string that is guaranteed to contain a special character, a lowercase letter, and a number, you can run this command:

```
$ ranchars -cld 8
3r&n*kc0
```

### Usage

`ranchars [options] length`

**Notes:**
- `length` must be a positive integer.
- `length` must  be greater than or equal to the specified number of character types. For example, it is impossible to generate a 3-character string that includes at least one of each of the following: numeric digit, special character, lowercase letter, and uppercase letter.

### Options

`-d` Include numeric digits (0-9)

`-l` Include lowercase letters (a-z)

`-u` Include uppercase letters (A-Z)

`-L` Include both lowercase and uppercase letters (a shortcut for specifying both `-l` and `-u`)

`-c` Include special characters (see [Special characters](###-special-characters))

`-s` Include the space character

`-a` Include all kinds of characters (equivalent to specifying `-d`, `-l`, `-u`, `-c`, and `-s`)

`-n` Omit trailing newline character from output

### Special characters

The "special characters" type refers to this set of characters:

``!"#$%&'()*+,-./:;<=>?@[]^_`{\}~``

Backslashes (`\`) are omitted from the set to avoid accidental character escaping.

The space character is considered a distinct character type to make it easier to generate random strings with a large character space but without spaces.

### Examples

Generate a random 6-digit number:

```
$ ranchars -d 6
492029
```

Generate a random string of 20 characters that includes numbers, letters of both cases, and special characters.

```
$ ranchars -dLc 20
DATkFzLT)A=Ec1["wEFc
```

Generate a random string of 32 characters that includes numbers and letters of both cases. (This is the set of character types implicitly specified if no options for character types are used.)

```
$ ranchars 32
6wmK9s8xpZbGzJBIt9IcvNtejHJJvnPU
```

## Randomness

Ranchars uses Go's `crypto/rand` package. It's designed to be acceptable for contexts that require cryptographically secure randomness.

From Go's [documentation](https://golang.org/pkg/crypto/rand/) on the `crypto/rand` package:

> Package rand implements a cryptographically secure random number generator.

### The process of generating random characters

To generate a string of random characters, Ranchars follows the following process:

1. To produce a single character, a single byte is generated randomly.

1. Ranchars determines the character types that were specified when Ranchars was run, and tests the randomly generated byte to see if the byte represents an ASCII character of one of the specified character types.

1. If this test fails, the byte is discarded and a new byte is generated and tested. This process repeats until the count of characters that have passed this test matches the specified length given to Ranchars.

1. Now that the correct number of characters have been generated, Ranchars tests the string as a whole to see if all specified character types are represented in this string by at least on character each.

1. If this test passes, the string is returned to the user.

1. If this test fails, the string is discarded, and the process starts again from the first step.

