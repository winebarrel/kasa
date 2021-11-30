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

  ls --team=STRING --token=STRING [<path>]
    List posts.

  mv --team=STRING --token=STRING <source> <target>
    Move posts.

  mvcat --team=STRING --token=STRING <from> <to>
    Move category.

  post --team=STRING --token=STRING --name=STRING <body-file> [<post-num>]
    New post.

  rm --team=STRING --token=STRING <post-num>
    Delete post.

  rmx --team=STRING --token=STRING <path>
    Delete multiple posts.

  search --team=STRING --token=STRING <query>
    Search posts.

Run "kasa <command> --help" for more information on a command.
```
