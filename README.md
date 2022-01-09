# kasa ☂️

CLI for [esa](https://esa.io/).

## Usage

```
Usage: kasa --team=STRING --token=STRING <command>

Flags:
  -h, --help            Show context-sensitive help.
      --version
      --team=STRING     esa team ($ESA_TEAM)
      --token=STRING    esa access token ($ESA_TOKEN)

Commands:
  cat --team=STRING --token=STRING [<path>]
    Print post.

  info --team=STRING --token=STRING <post-num>
    Show post info.

  ls --team=STRING --token=STRING [<path>]
    List posts.

  mv --team=STRING --token=STRING <source> <target>
    Move posts.

  mvcat --team=STRING --token=STRING <from> <to>
    Move category.

  post --team=STRING --token=STRING [<post-num>]
    New/Update post.

  rm --team=STRING --token=STRING <post-num>
    Delete post.

  rmx --team=STRING --token=STRING <path>
    Delete multiple posts.

  search --team=STRING --token=STRING <query>
    Search posts.

Run "kasa <command> --help" for more information on a command.
```

### Example

```sh
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

$ kasa rmx baz/
rm 'baz/title'
Do you want to delete posts? (y/n) [n]: y

~% kasa ls
2021-09-07 11:07:44  -    https://winebarrel.esa.io/posts/1        README
```

## Installation

## OS X

```
brew tap winebarrel/kasa
brew install kasa
```
