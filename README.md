# mark

## Installation

There are three ways to install `mark`:
- using `go install`
- use a pre-built binary
- build `mark` from scratch


### Using `go install`

If you [already have Go installed], then installing `mark` is as simple as:

```bash
go install github.com/commondatageek/mark
```

You'll then need to make sure that your `GOPATH` (typically `$HOME/go` by default) is on your `PATH` variable.


### Use a Pre-built Binary

1. Download the appropriate pre-built binary from the [Releases](https://github.com/commondatageek/mark/releases/) page.
2. Extract the `mark` file from the tar archive.  For example:

```bash
tar -xvf mark-v0.6.1-darwin-amd64.tar.gz
```

3. Move the `mark` binary into your preferred location on your path.  For example:

```bash
sudo mv ./mark /usr/local/bin
```


### Build `mark` from Scratch

1. [Install](https://go.dev/doc/install) Go
2. Clone the `mark` repository to your local machine
3. From the `mark` repository directory, run `go build .`
    - (Alternatively, if you have `just` task runner [installed](https://just.systems/man/en/chapter_5.html), you can run `just build`)
4. Move the newly built `mark` binary to some directory on your `PATH`. For example:

```bash
sudo mv ./mark /usr/local/bin
```

In the future, we'll have a [devcontainer](https://code.visualstudio.com/docs/devcontainers/containers) that will make this easier.


## Usage

### The Database File

You must have a file called `.bookmarks.jsonl` in your `$HOME` directory.

It should have the following format, one item per line:

```json
{"names": ["Hacker News", "hn"], "tags": ["news", "technology"], "url": "https://news.ycombinator.com/news"}
{"names": ["InfoQ News", "infoq"], "tags": ["news", "technology"], "url": "https://infoq.com"}
{"names": ["Mark Github Repository", "gh/mark"], "tags": ["personal", "github", "mark", "cli", "golang"], "url": "https://github.com/commondatageek/mark"}
```

You'll want to add good tags and names, as it will help immensely when you are searching for something 6 months from now.


### Search your item database

The search functionality will perform indexing on your database and match all
the substrings of the words you type in against the substrings in the words in your items
--including names, tags, and URL.

When it gives you search results, you can `Command`+Click on the URL and open it:
```bash
mark
```

### Open a specific item 

Open a specific item by giving one of its names:
```bash
mark my/cool/item/name
```

