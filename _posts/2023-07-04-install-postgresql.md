---
layout: post
title: "Install Postgresql at one click"
categories: [ tech, pgsql ]
include_file: block/txt/include.md
src_file: block/script/installpgsql.sh
---

I wanted to use postgresql in my systems. So I decided to install it from source code. When I referred to the official webiste, I installed it successfully both on Ubuntu22.04 server and Mac m1. Then I wrote a shell script to make the installation of postgresql at one click.

The script are not really one click right now, as the steps of the after installation shoud be done by hand which are showed as comments. But I will continue to make it work with the target.

## Installation script

```shell
{% include {{ page.src_file }} %}
```

## Download

 <a href="/assets/block/script/installpgsql.sh" download>download installation script</a>

{% include {{ page.include_file }}  %}

