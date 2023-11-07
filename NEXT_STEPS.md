# Next Steps

- Keep track of how many times each one has been selected so that more frequently
  accessed items bubble to the top or are weighted more strongly
- implement soundex?
- add tags?
- add creation_date
- allow randomly opening something in your bookmarks
- do spaced repetition on things where we weight things we haven't seen in a while more highly
- create abstraction for multiple backends
    - S3
    - local
    - Dropbox
    - iCloud
- create a config file in ~/.mark
- allow the user to choose not just which browser to open something with, but which program
- allow to use my Obsidian notes as a database of sorts
    - the use cases are just so parallel and we should find a way to make it work better
- bring in the "welcome ot the bookmarks folder" cartoon from Dank Memes and make a justification for the program
- Use FSTs and do a single pass instead of 1) build index, 2) query index
    - We could totally do this in a single pass
    - Deconstruct the query
    - Deconstruct each item in a linear (or concurrent) sequence
    - Make the comparison for each item to the query
    - return the results as usual, count them, sort by count desc
    - This will really decrease the memory usage of the program as well
- Use levenshtein disntance with the names field to give "did you mean?" suggestions
- add UI for adding
- add UI for removing
- figure out how to store
- Create groups of URLs for a single entry
    - open a up a group of tabs for when you have a specific task that uses multiple sites
- add UI for updating
- have a hierarchy of names
    - can we have a group of names that are "work" names?
    - or "personal" names?
- choose which browser to open the URL with
    - Make a patch on the browser package that allows you to choose which browser to open
    - Which browsers?
        - Safari
        - Firefox
        - Chrome
- choose other launch arguments that we might use with a particular browser
    - Make a patch on the browser package that allows you to specify browser launch arguments
    - What arguments?
        - Let's see what options we have when we launch Chrome and work from there