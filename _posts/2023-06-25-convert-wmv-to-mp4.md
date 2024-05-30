---
layout: post
title: "Convert WMV To MP4"
categories: shell
include_file: block/txt/include.md
script_file: block/script/convertmp4.sh
---
I wanted to convert all my wmv files into mp4 files at one click. So I wrote a shell script to run it. 

## script


```shell
{% include {{ page.script_file }} %}
```

## usage

```shell
# Enter the folder containing the source wmv files
$ ls | xargs /<repo_path>/_includes/block/script/convertmp4.sh 
# Let the conversion job be running and have a cup of coffee.
# Have fun@!@
```

## todo

* Make most of cpu to run with multiple processes.


## references

* [Jekyll - DOCS - Includes](https://jekyllrb.com/docs/includes/){:target="_blank"}


{% include {{ page.include_file }} %}
