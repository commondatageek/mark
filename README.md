# mark

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


### Search your bookmark database

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
mark my/cool/bookmark/name
```
