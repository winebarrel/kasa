---
name: kasa
description: CLI tool for esa.io - list, create, edit, move, delete posts and manage tags/categories
allowed-tools: Read, Grep, Glob, Bash(kasa *)
---

# kasa

CLI for [esa.io](https://esa.io/).

## Global Flags

```
--team=STRING     esa team ($ESA_TEAM)
--token=STRING    esa access token ($ESA_TOKEN)
--debug           Debug flag
```

`<path>` には記事名のほか、Post URL (`https://<TEAM>.esa.io/posts/<NUM>` or `//<NUM>`) も指定可能。

## Commands

### ls - List posts

```
kasa ls [<path>] [--json] [-p PAGE] [--recursive/--no-recursive]
```

- `<path>` - 記事名/カテゴリ/タグ (省略可)
- `--json` - JSON出力
- `-p` - ページ番号 (default: 1)
- `--recursive/--no-recursive` - 再帰的にリスト (default: true)

### cat - Print a post

```
kasa cat <path>
```

### info - Show a post info

```
kasa info <path>
```

### open - Open a post in the browser

```
kasa open <path>
```

### search - Search posts

```
kasa search <query> [--json] [-p PAGE]
```

- `<query>` - 検索クエリ ([esa検索構文](https://docs.esa.io/posts/104))

### post - Create/Update a post

```
kasa post [<path>] [-n NAME] [-b BODY_FILE] [-t TAG ...] [-c CATEGORY] [--wip/--no-wip] [-m MESSAGE] [--notice/--no-notice]
```

- `<path>` - 更新する記事番号 (省略時は新規作成)
- `-b` - 本文ファイル (`-` でstdin)
- `-n` - 記事タイトル
- `-t` - タグ (複数指定可)
- `-c` - カテゴリ
- `--wip/--no-wip` - WIP指定
- `-m` - メッセージ
- `--notice/--no-notice` - 通知

### touch - Create an empty post

```
kasa touch <path> [--notice/--no-notice]
```

### edit - Edit a post

```
kasa edit <path> --editor=STRING [--notice/--no-notice]
```

- `--editor` - エディタ ($EDITOR)

### append - Append text to a post

```
kasa append <path> -b BODY_FILE [--prefix TEXT] [--notice/--no-notice]
```

- `-b` - 追加する本文ファイル
- `--prefix` - 先頭に付与するテキスト
- `--notice/--no-notice` - 通知 (default: true)

### comment - Comment on a post

```
kasa comment <path> -b BODY_FILE
```

### import - Import a file or directory

```
kasa import <src> <path> [--notice/--no-notice] [--wip/--no-wip]
```

- `<src>` - ソースファイルまたはディレクトリ (`-` でstdin)
- `<path>` - インポート先の記事名

### mv - Move posts

```
kasa mv <source> <target> [-s] [-f] [--notice/--no-notice] [-p PAGE] [--recursive/--no-recursive]
```

- `<source>` - 移動元 (記事名/カテゴリ/タグ)
- `<target>` - 移動先 (記事名/カテゴリ)
- `-s` - 検索モード ([esa検索構文](https://docs.esa.io/posts/104))
- `-f` - 確認スキップ

### mvcat - Rename a category

```
kasa mvcat <from> <to>
```

### cp - Copy posts

```
kasa cp <source> <target> [-f] [--notice/--no-notice] [-p PAGE] [--recursive/--no-recursive]
```

### rm - Delete posts

```
kasa rm <path> [-s] [-f] [-p PAGE] [--recursive/--no-recursive]
```

- `-s` - 検索モード
- `-f` - 確認スキップ

### rmi - Delete a post by number

```
kasa rmi <path> [-f]
```

### tag - Tagging posts

```
kasa tag <path> [-t TAG ...] [-o] [-d] [-s] [-f] [--notice/--no-notice] [-p PAGE] [--recursive/--no-recursive]
```

- `-t` - タグ (複数指定可)
- `-o` - タグを上書き
- `-d` - タグを削除
- `-s` - 検索モード

### tags - Print tags

```
kasa tags [-p PAGE]
```

### wip - Make posts WIP

```
kasa wip <path> [-s] [-f] [--notice/--no-notice] [-p PAGE] [--recursive/--no-recursive]
```

### unwip - Unwip posts

```
kasa unwip <path> [-s] [-f] [--notice/--no-notice] [-p PAGE] [--recursive/--no-recursive]
```

### stats - Print team statistics

```
kasa stats
```

## Usage Examples

```sh
# 環境変数を設定
export ESA_TEAM=myteam
export ESA_TOKEN=...

# 記事一覧
kasa ls
kasa ls foo/bar/

# タグで絞り込み
kasa ls '#tagA'

# 記事を表示
kasa cat foo/bar/title

# 記事を作成
echo hello | kasa post -b - -n title -c foo/bar

# 記事を更新 (記事番号指定)
echo world | kasa post 38 -b -

# 記事を移動
kasa mv foo/bar/ zoo/

# カテゴリをリネーム
kasa mvcat old/category new/category

# 記事を削除
kasa rm baz/

# エディタで編集
kasa edit any/post

# ファイルをインポート
echo hello | kasa import - any/import/post
kasa import ./local_dir/ any/import/

# タグ付け
kasa tag foo/bar/ -t tagA -t tagB

# 検索
kasa search "keyword"
```
