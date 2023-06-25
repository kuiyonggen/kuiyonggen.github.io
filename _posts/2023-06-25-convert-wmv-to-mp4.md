---
layout: post
title: "Convert WMV To MP4"
categories: shell
script_file: script/convertmp4.sh
---
I wanted to convert all my wmv files into mp4 files at one click. So I wrote a shell script to run it. 

## script

This code block is included using the below Jekyll include tag expression.
```
{% raw %} { % include {{ page.script_file }} % } {% endraw %}
```
I'm happy it worked on Ubuntu 22.04, because Markdown doesn't support including other file content.
I think as a programmer it's important to keep the single responsibility principle.
Keep it clean and simple. That's very important for maintainability.

```
{% include {{ page.script_file }} %}
```

## usage

```bash
# Enter the folder containing the source wmv files
$ ls | xargs /<repo_path>/_include/script/convertmp4.sh 
# Let the conversion job be running and have a cup of coffee.
# Have fun@!@
```

## todo

* Make most of cpu to run with multiple processes.

## references

* [Jekyll - DOCS - Includes](https://jekyllrb.com/docs/includes/){:target="_blank"}
