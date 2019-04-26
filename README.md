# @ritou11/hub-cup

[![GitHub last commit](https://img.shields.io/github/last-commit/ritou11/hub-cup.svg?style=flat-square)](https://github.com/ritou11/hub-cup)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ritou11/hub-cup.svg?style=flat-square)](https://github.com/ritou11/hub-cup)
[![license](https://img.shields.io/github/license/ritou11/hub-cup.svg?style=flat-square)](https://github.com/ritou11/hub-cup/blob/master/LICENSE.md)

> Sync your github forks without Git or Node.
>
> This is the Go version of [hub-sync](https://github.com/b1f6c1c4/hub-sync). You could use binary files rather than heavy envirnoments.

## Download Binary

Prebuilt binaries available at https://github.com/ritou11/hub-cup/releases

## TL;DR

```sh
# Generate a token **with public_repo scope** at https://github.com/settings/tokens
echo the-token > ~/.hub-cup
# Update your webpack fork default branch to the latest of upstream:
hub-cup webpack
# Update your material-ui fork default branch to the latest of upstream:
hub-cup material-ui
# Update your material-ui fork master branch to the latest of upstream:
hub-cup material-ui/master
# Update your antediluvian io.js fork to the latest nodejs:
hub-cup io.js # name doesn't need to match exactly
hub-cup io.js nodejs/node # but you MUST specify the repo if you want to sync to the upstream of upstream
```

## Why

Reference to [hub-sync](https://github.com/b1f6c1c4/hub-sync).

Since "hub-sync" is a part of the well-known tool [hub](https://hub.github.com/hub.1.html), this tool is named "hub-cup", which means "Make Git**Hub** forks **c**atch **up** with origins".

## Usage

```
NAME:
   hub-cup - Make your github forks catch up with origins

USAGE:
   hub-cup <what> [<from>]

VERSION:
   0.1.0

AUTHOR:
   Nogeek <ritou11@gmail.com>

GLOBAL OPTIONS:
   --token value, -t value  Github token, see https://github.com/settings/tokens
   --token-file path        path to your Github token file (default: "~/.hub-cup")
   --force, -f              As if {git push --force}
   --dry-run, -n            Don't actually update
   --debug                  print debug messages
   --help, -h               print the help
   --version, -v            print the version
```

## License

MIT
