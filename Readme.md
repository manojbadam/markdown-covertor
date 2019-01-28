## Markdown convertor

Often times, we need to convert our docs on github readme to Confluence Wiki markup. This repo helps in converting it. 
This is still work in progress. Currently capable of converting headers, code blocks, bold, blockquotes, hyperlink messages. This needs to extend futher for Table, horizontal rule, images and Lists.

### How to run this locally
```
go get git@github.com:manojbadam/markdown-covertor.git
cd $GOPATH/src/github.com/manojbadam/markdown-convertor
export FILE_PATH="<path/to/readme.md>"
go build
./markdown-covertor wiki -i /Users/badam/test/clone/gist/cilium-podcidr.md -o converted.md
```

Once the markdown is converted into `output.wiki` file,  refer this [page](https://confluence.atlassian.com/doc/confluence-wiki-markup-251003035.html#ConfluenceWikiMarkup-markdownCanIinsertmarkdown?) for how to insert the markdown 


##### References
* [Github Markdown CheatSheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)
* [Confluence wiki Markdown](https://confluence.atlassian.com/doc/confluence-wiki-markup-251003035.html)
