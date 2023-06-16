---
layout: post
title: "Use UML in Jekyll"
categories: tech
---

It works! Thanks to [PlantUML Github action](https://github.com/marketplace/actions/plantuml-generator){:target="_blank"}.

And mermaid may work at the same way. But I don't need it right now. Plantuml is enough for me.

***************************************

It doesn't work on github page described as below. I continue to find the method to code to uml on github page.

****************************************

I was looking for an uml plugin for my jekyll website. I found some plugins worked at local environment and must install some dependencies. Those were not what I wanted. Today I found this plugin - [A Jekyll plugin to provide powerful supports for table, mathjax, plantuml, mermaid, emoji, video, audio, youtube, vimeo, dailymotion, soundcloud, spotify, etc.](https://github.com/jeffreytse/jekyll-spaceship)

I installed the plugin by following the steps in the link. Then I used some demos and found it very simple.

As a programmer, I love coding to uml. I think it wastes time to use a mouse. Time is limited. So It's important to focus on coding to make something valuable.


# plantuml

* [plantuml class diagram](https://plantuml.com/class-diagram)


## Class diagram

![](/assets/dia/cls.png)

# mermaid.js

* [mermaid.js](https://mermaid.js.org/syntax/flowchart.html)

## Flowchart

@startmermaid
flowchart LR

A[Hard] -->|Text| B(Round)
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

* [How to use GitHub Actions environment variables](https://snyk.io/blog/how-to-use-github-actions-environment-variables/){:target="_blank"}
* [Permissions for the GITHUB_TOKEN](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token){:target="_blank"}
* [plantuml class diagram](https://plantuml.com/class-diagram){:target="_blank"}
* [PlantUML GitHub action](https://github.com/marketplace/actions/plantuml-generator){:target="_blank"}
