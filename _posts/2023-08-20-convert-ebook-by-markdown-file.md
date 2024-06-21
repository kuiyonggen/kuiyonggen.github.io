---
layout: post
title: "Convert Ebook By Markdown File"
categories: ["tech", "shell"]
script_file: block/script/ebookbymd.sh
---
I love reading books with Kindle. There is a need to create ebooks. So I studied the tools and found a way suitable for the need.

Thanks for pandoc and calibre utility. They make the need come true. Then I wrote the below shell script for convenience.

Firstly, you must install pandoc and calibre utility.

Once you create ebooks, you can add them into you calibre library, with which you can maintain ebooks very easily.


```bash
{% include {{ page.script_file }} %}
```
