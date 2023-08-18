# kasa ☂️

[![test](https://github.com/kanmu/kasa/actions/workflows/test.yml/badge.svg)](https://github.com/kanmu/kasa/actions/workflows/test.yml)

CLI for [esa.io](https://esa.io/).

![](https://user-images.githubusercontent.com/117768/149801227-d64d4895-50a4-4dde-bc9e-b27950a99b81.gif)

## Usage

```
{{ .usage }}
```

### Example

```sh
$ export ESA_TEAM=kanmu
$ export ESA_TOKEN=...

$ kasa ls
2021-09-07 11:07:44  -    https://kanmu.esa.io/posts/1        README

$ echo hello | kasa post -b - -n title -c foo/bar
https://kanmu.esa.io/posts/38

$ kasa cat foo/bar/title
hello

$ echo world | kasa post 38 -b -
https://kanmu.esa.io/posts/38

$ kasa cat foo/bar/title
world

$ kasa mv foo/bar/ zoo/
mv 'foo/bar/title' 'zoo/'
Do you want to move posts? (y/n) [n]: y

$ kasa ls
2021-09-07 11:07:44  -    https://kanmu.esa.io/posts/1        README
2022-01-09 09:47:24  WIP  https://kanmu.esa.io/posts/38       zoo/title

$ kasa ls zoo/
2022-01-09 09:47:24  WIP  https://kanmu.esa.io/posts/38       zoo/title

$ kasa post foo/bar/title -t tagA
https://kanmu.esa.io/posts/38

$ kasa ls '#tagA'
2022-01-09 09:47:24  WIP  https://kanmu.esa.io/posts/38       zoo/title  [#tagA]

$ kasa mv '#tagA' baz/
mv 'zoo/title' 'baz/'
Do you want to move posts? (y/n) [n]: y

$ kasa ls
2021-09-07 11:07:44  -    https://kanmu.esa.io/posts/1        README
2022-01-09 09:47:24  WIP  https://kanmu.esa.io/posts/38       baz/title  [#tagA]

$ kasa rm baz/
rm 'baz/title'
Do you want to delete posts? (y/n) [n]: y

$ kasa ls
2021-09-07 11:07:44  -    https://kanmu.esa.io/posts/1        README

$ kasa edit any/new/post
https://kanmu.esa.io/posts/39

$ echo hello | kasa import - any/import/post
https://kanmu.esa.io/posts/40

$ date > hello
$ kasa import hello any/import/
https://kanmu.esa.io/posts/41

$ kasa ls any/import/
2023-08-18 10:12:00  -    https://kanmu.esa.io/posts/40       any/import/post
2023-08-18 10:12:30  -    https://kanmu.esa.io/posts/41       any/import/hello

$ mkdir foo/
$ touch foo/bar.txt
$ mkdir foo/zoo
$ touch foo/zoo/baz.txt

$ kasa import ./foo/ any/import2/
https://kanmu.esa.io/posts/42
https://kanmu.esa.io/posts/43

$ kasa ls any/import2/
2023-08-18 10:32:00  -    https://kanmu.esa.io/posts/42       any/import2/bar.txt
2023-08-18 10:22:30  -    https://kanmu.esa.io/posts/43       any/import2/zoo/baz.txt
```

## Installation

```sh
# OSX or Linux
brew tap kanmu/tools
brew install kasa
```

## Install shell completions

```sh
kasa install-completions >> ~/.zshrc
```
