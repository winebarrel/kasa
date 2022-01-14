# kasa ☂️

[![test](https://github.com/winebarrel/kasa/actions/workflows/test.yml/badge.svg)](https://github.com/winebarrel/kasa/actions/workflows/test.yml)

CLI for [esa](https://esa.io/).

## Usage

```
{{ .usage }}
```

### Example

```sh
$ export ESA_TEAM=winebarrel
$ export ESA_TOKEN=...

$ kasa ls
2021-09-07 11:07:44  -    https://winebarrel.esa.io/posts/1        README

$ echo hello | kasa post -b - -n title -c foo/bar
https://winebarrel.esa.io/posts/38

$ kasa cat foo/bar/title
hello

$ echo world | kasa post 38 -b -
https://winebarrel.esa.io/posts/38

$ kasa cat foo/bar/title
world

$ kasa mv foo/bar/ zoo/
mv 'foo/bar/title' 'zoo/'
Do you want to move posts? (y/n) [n]: y

$ kasa ls
2021-09-07 11:07:44  -    https://winebarrel.esa.io/posts/1        README
2022-01-09 09:47:24  WIP  https://winebarrel.esa.io/posts/38       zoo/title

$ kasa ls zoo/
2022-01-09 09:47:24  WIP  https://winebarrel.esa.io/posts/38       zoo/title

$ kasa post 38 -t tagA
https://winebarrel.esa.io/posts/38

$ kasa ls '#tagA'
2022-01-09 09:47:24  WIP  https://winebarrel.esa.io/posts/38       zoo/title  [#tagA]

$ kasa mv '#tagA' baz/
mv 'zoo/title' 'baz/'
Do you want to move posts? (y/n) [n]: y

$ kasa ls
2021-09-07 11:07:44  -    https://winebarrel.esa.io/posts/1        README
2022-01-09 09:47:24  WIP  https://winebarrel.esa.io/posts/38       baz/title  [#tagA]

$ kasa rm baz/
rm 'baz/title'
Do you want to delete posts? (y/n) [n]: y

$ kasa ls
2021-09-07 11:07:44  -    https://winebarrel.esa.io/posts/1        README
```

## Installation

```sh
# OSX or Linux
brew tap winebarrel/kasa
brew install kasa
```

## Install shell completions

```sh
kasa install-completions >> ~/.zshrc
```
