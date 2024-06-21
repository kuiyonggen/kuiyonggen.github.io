---
layout: post
title: "惊艳的Shell命令"
categories: ["shell", "tech"]
---
这是我在平时工作中收集的令人惊叹的Shell命令，因为对我工作帮助很大。我希望随时可以查阅这些命令，如果正好也帮助到你，我将非常高兴。



## 文件系统

### 查找

#### 按文件扩展名递归地查找所有文件

如果你想移动所有同类型文件，你可以使用"-exec"和"-o"。这有一个例子可以执行这个任务。

```shell

find /data/ -type f \( -name "*.pdf" -o -name "*.epub" -o  -name "*.mobi" -o -name "*.azw3" -o -name "*.txt" \) -exec mv {} books/ \;

```

