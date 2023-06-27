---
layout: post
title: "Convert WMV To MP4"
categories: shell
include_file: txt/include.md
script_file: script/convertmp4.sh
---
I wanted to convert all my wmv files into mp4 files at one click. So I wrote a shell script to run it. 

## script

{% include {{ page.include_file }} %}

```shell
{% include {{ page.script_file }} %}
```

## usage

```shell
# Enter the folder containing the source wmv files
$ ls | xargs /<repo_path>/_include/script/convertmp4.sh 
# Let the conversion job be running and have a cup of coffee.
# Have fun@!@
```

## todo

* Make most of cpu to run with multiple processes.

## references

* [Jekyll - DOCS - Includes](https://jekyllrb.com/docs/includes/){:target="_blank"}
