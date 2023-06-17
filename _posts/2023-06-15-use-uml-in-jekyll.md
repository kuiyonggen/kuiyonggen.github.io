---
layout: post
title: "Use UML in Jekyll"
categories: tech
---
## Update 2


I works! That's what I want. Thanks to Github action and this post([Deploying Jekyll sites to GitHub Pages using GitHub Actions](https://milanaryal.com.np/deploying-jekyll-sites-to-github-pages-using-github-actions/#:~:text=Deploying%20Jekyll%20sites%20to%20GitHub%20Pages%20using%20GitHub,Pages%20site.%20...%204%20Build%20and%20deploy%20){:target="blank"}).

There may be two problems as follow.

1. build error


```bash
GitHub Metadata: Error processing value 'baseurl':
  Liquid Exception: No repo name found. Specify using PAGES_REPO_NWO environment variables, 'repository' in your configuration, or set up an 'origin' git remote pointing to your github.com repository. in /_layouts/post.html
             ERROR: YOUR SITE COULD NOT BE BUILT:
                    ------------------------------------
                    No repo name found. Specify using PAGES_REPO_NWO environment variables, 'repository' in your configuration, or set up an 'origin' git remote pointing to your github.com repository.
Error: Process completed with exit code 1.
```

The solution is adding the repository setting in _config.yml between 'title' and 'author'. Then commit and push to trigger the Github action again.

```
title: Kuiyonggen
repository: kuiyonggen/kuiyonggen.github.io
author: Kuiyonggen
```


2. The url of the deployed page contained "pages/kuiyonggen", which caused that static files were not found.

The solution is change the bundle execute command in .github/workflow/pages.yml.

```bash 
# original
# I don't know how to display double curly baces in the code block. 
# It wasted time to fix it. I noted here for others and me.
# Thanks to the post: link How to escape double curly braces in Jekyll.
# https://devcoops.com/escape-double-curly-braces-inside-a-markdown/
# As it mentioned:
# Jekyll uses liquid tags, so if you are searching 
# how to escape double curly braces in Markdown 
# it will not solve the issue.
{% raw %}
run: bundle exec jekyll build --baseurl ${{ steps.pages.outputs.base_path }} --profile --trace
{% endraw %}
# changed
run: bundle exec jekyll build --baseurl '' --profile --trace
```

"Keep looking, don't settle!" - by Steve Jobs.

I think that's the only way to find what you want.

I will keep looking too and make the website better and better.

***************************************
## Update 1


It works! Thanks to [PlantUML Github action](https://github.com/marketplace/actions/plantuml-generator){:target="_blank"}.

The only one problem encountered is no push permission. The solution is as follow.

1. Navigate to "Settings" of the github repo.
2. Click "Actions" on left side.
3. Choose the "Read and write permissions" item in the "Workflow permissions" section.
4. Click the "Save" button.
5. Trigger the action to run again.
6. Done.

And mermaid may work at the same way. But I don't need it right now. Plantuml is enough for me.

***************************************

It doesn't work on github page described as below. I continue to find the method to code to uml on github page.

****************************************

## Original

I was looking for an uml plugin for my jekyll website. I found some plugins worked at local environment and must install some dependencies. Those were not what I wanted. Today I found this plugin - [A Jekyll plugin to provide powerful supports for table, mathjax, plantuml, mermaid, emoji, video, audio, youtube, vimeo, dailymotion, soundcloud, spotify, etc.](https://github.com/jeffreytse/jekyll-spaceship)

I installed the plugin by following the steps in the link. Then I used some demos and found it very simple.

As a programmer, I love coding to uml. I think it wastes time to use a mouse. Time is limited. So It's important to focus on coding to make something valuable.


# plantuml

* [plantuml class diagram](https://plantuml.com/class-diagram)


## Class diagram

@startuml
class Object << general >>
Object <|--- ArrayList

note top of Object : In java, every class\nextends this one.

note "This is a floating note" as N1
note "This note is connected\nto several objects." as N2
Object .. N2
N2 .. ArrayList

class Foo
note left: On last defined class
@enduml

# mermaid.js

* [mermaid.js](https://mermaid.js.org/syntax/flowchart.html)

## Flowchart

@startmermaid
flowchart LR

A[Easy] -->|Text| B(Good)
B --> C{Decision}
C -->|One| D[Result 1]
C -->|Two| E[Result 2]
@endmermaid


## Sequence diagram

@startmermaid
sequenceDiagram
Alice->>John: Hello John, how are you?
loop Healthcheck
    John->>John: Fight against hypochondria
end
Note right of John: Rational thoughts!
John-->>Alice: Great!
John->>Bob: How about you?
Bob-->>John: Jolly good!
@endmermaid



## Gantt chart

@startmermaid
gantt
    section Section
    Completed :done,    des1, 2014-01-06,2014-01-08
    Active        :active,  des2, 2014-01-07, 3d
    Parallel 1   :         des3, after des1, 1d
    Parallel 2   :         des4, after des1, 1d
    Parallel 3   :         des5, after des3, 1d
    Parallel 4   :         des6, after des4, 1d
@endmermaid

## State diagram

I don't know why it happend. It's too big. I didn't find a solution to fix it. But it doesn't matter. Just use another processor. Plantuml works for me.

@startmermaid
stateDiagram-v2
[*] --> Still
Still --> [*]

Still --> Moving
Moving --> Still
Moving --> Crash
Crash --> [*]
@endmermaid

## Pie chart

@startmermaid
pie
"Dogs" : 386
"Cats" : 85.9
"Rats" : 15
@endmermaid



## User Journey diagram

@startmermaid
journey
  title My working day
  section Go to work
    Make tea: 5: Me
    Go upstairs: 3: Me
    Do work: 1: Me, Cat
  section Go home
    Go downstairs: 5: Me
    Sit down: 3: Me
@endmermaid


## C4 diagram

@startmermaid
C4Context
title System Context diagram for Internet Banking System

Person(customerA, "Banking Customer A", "A customer of the bank, with personal bank accounts.")
Person(customerB, "Banking Customer B")
Person_Ext(customerC, "Banking Customer C")
System(SystemAA, "Internet Banking System", "Allows customers to view information about their bank accounts, and make payments.")

Person(customerD, "Banking Customer D", "A customer of the bank, <br/> with personal bank accounts.")

Enterprise_Boundary(b1, "BankBoundary") {

  SystemDb_Ext(SystemE, "Mainframe Banking System", "Stores all of the core banking information about customers, accounts, transactions, etc.")

  System_Boundary(b2, "BankBoundary2") {
    System(SystemA, "Banking System A")
    System(SystemB, "Banking System B", "A system of the bank, with personal bank accounts.")
  }

  System_Ext(SystemC, "E-mail system", "The internal Microsoft Exchange e-mail system.")
  SystemDb(SystemD, "Banking System D Database", "A system of the bank, with personal bank accounts.")

  Boundary(b3, "BankBoundary3", "boundary") {
    SystemQueue(SystemF, "Banking System F Queue", "A system of the bank, with personal bank accounts.")
    SystemQueue_Ext(SystemG, "Banking System G Queue", "A system of the bank, with personal bank accounts.")
  }
}

BiRel(customerA, SystemAA, "Uses")
BiRel(SystemAA, SystemE, "Uses")
Rel(SystemAA, SystemC, "Sends e-mails", "SMTP")
Rel(SystemC, customerA, "Sends e-mails to")
@endmermaid

## references

* [Deploying Jekyll sites to GitHub Pages using GitHub Actions](https://milanaryal.com.np/deploying-jekyll-sites-to-github-pages-using-github-actions/#:~:text=Deploying%20Jekyll%20sites%20to%20GitHub%20Pages%20using%20GitHub,Pages%20site.%20...%204%20Build%20and%20deploy%20){:target="blank"}
* [How to use GitHub Actions environment variables](https://snyk.io/blog/how-to-use-github-actions-environment-variables/){:target="_blank"}
* [Permissions for the GITHUB_TOKEN](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token){:target="_blank"}
* [plantuml class diagram](https://plantuml.com/class-diagram){:target="_blank"}
* [PlantUML GitHub action](https://github.com/marketplace/actions/plantuml-generator){:target="_blank"}
