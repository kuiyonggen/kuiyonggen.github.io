---
layout: post
title: "fx - A dependency injection based application framework for go"
mermaid: true
categories: [ "tech", "golang" ]
src_file: block/code/go/fxdemo/main.go
md_readme_file: block/md/fx/get-started/README.md
md_minimal_file: block/md/fx/get-started/minimal.md 
md_http_server_file: block/md/fx/get-started/http-server.md
md_another_handler_file: block/md/fx/get-started/another-handler.md
md_conclusion_file: block/md/fx/get-started/conclusion.md
md_echo_handler_file: block/md/fx/get-started/echo-handler.md
md_logger_file: block/md/fx/get-started/logger.md
md_many_handlers_file: block/md/fx/get-started/many-handlers.md
md_registration_file: block/md/fx/get-started/registration.md
md_introduction_file: block/md/fx/intro.md
md_parameter_objects_file: block/md/fx/parameter-objects.md
md_result_objects_file: block/md/fx/result-objects.md
md_annotations_file: block/md/fx/annotate.md
md_value_groups_file: block/md/fx/value-groups/README.md
md_value_groups_feeding_file: block/md/fx/value-groups/feed.md
md_value_groups_consuming_file: block/md/fx/value-groups/consume.md
md_concepts_application_lifecycle_file: block/md/fx/lifecycle.md
md_concepts_modules_file: block/md/fx/modules.md
md_community_contributing_file: block/md/fx/contributing.md
---
A dependency injection based application framework for Go.

{% include {{ page.md_readme_file }} %}
{% include {{ page.md_minimal_file }} %}
{% include {{ page.md_http_server_file }} %}
{% include {{ page.md_echo_handler_file }} %}
{% include {{ page.md_logger_file }} %}
{% include {{ page.md_registration_file }} %}
{% include {{ page.md_another_handler_file }} %}
{% include {{ page.md_many_handlers_file }} %}
{% include {{ page.md_conclusion_file }} %}
{% include {{ page.md_introduction_file }} %}
{% include {{ page.md_parameter_objects_file }} %}
{% include {{ page.md_result_objects_file }} %}
{% include {{ page.md_annotations_file }} %}
{% include {{ page.md_value_groups_file }} %}
{% include {{ page.md_value_groups_feeding_file }} %}
{% include {{ page.md_value_groups_consuming_file }} %}
{% include {{ page.md_concepts_application_lifecycle_file }} %}
{% include {{ page.md_concepts_modules_file }} %}
{% include {{ page.md_community_contributing_file }} %}

# demo

## create the demo project

```bash
# cd your workspace directory
mkdir fxdemo
cd fxdemo
go mod init fxdemo
```

## code
```bash
cd fxdemo
vim main.go
```
```go
{% include {{ page.src_file }} %}
```

## build

```bash
go mod tidy
go build
./fxdemo
```

# reference

* [Introduction - Fx is a dependency injection system for Go.](){:target="_blank"}

* [github - A dependency injection based application framework for Go.](https://github.com/uber-go/fx){:target="_blank"}
