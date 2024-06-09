---
layout: post
title: "Awesome Shell Commands"
categories: ["shell", "tech"]
---
These are the increasing collection of awesome shell commands, as they help a lot with my computer work. I hope to use these as soon as I need them, and if they can also help you, I will feel happy. 


# Awesome Shell Commands

## Filesystem

### Search

#### Find all files recursively by file types


If you want to move all the files, you can use "-exec" and "-o". Here is an example of how you can perform the task.

```shell

find /data/ -type f \( -name "*.pdf" -o -name "*.epub" -o  -name "*.mobi" -o -name ".azw3" -o -name "*.txt" \) -exec mv {} books/ \;

```

