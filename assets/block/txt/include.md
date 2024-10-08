This code block is included using the below Jekyll include tag expression.
The included file must be in the "_includes" folder of the Jekyll project.
```liquid
---
src_file: script/filename.sh
---
{% raw %} {% include {{ page.src_file }} %} {% endraw %}
```
I’m happy it worked on Ubuntu 22.04, because Markdown doesn’t support including other file content. I think as a programmer it’s important to keep the single responsibility principle. Keep it clean and simple. That’s very important for maintainability.
