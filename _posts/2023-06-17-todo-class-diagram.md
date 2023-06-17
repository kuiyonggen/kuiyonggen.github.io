---
layout: post
title: "Todo class diagram"
categories: tech
---

I make the below todo class diagram with plantuml in order to think about good design of a todo app.

@startuml
class Todo {
    - title: string
    - content: string
    - done: false
    - createTime: datetime
    - updateTime: datetime
    - userId: uint
    + add(): void
    + detail(): void
    + done(): void
    + remove(): void
    + subscribe(): void
    # inform(): void
    - post(): void 
}
@enduml



