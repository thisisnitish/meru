# MERU

Meru is a programming language written in go

**I think the best way to demestify a topic is to just take the topic and build a toy
project in it. It doesn't have to perfect, it doesn't have to be explore every single aspect.
But building something can give you an insight of lot of things.**

## Motivation

I really wanted to know a how a programming language is constructed, how things like tokenization,
lexical analysis, and code generation happens. What is even AST and how does it get created throughout this
process. To learn everything around it, I tried to build a programming language.

## Examples

Few Examples of Code Snippets of Meru.

**Example 1**

```
let five = 5;
let ten = 10;

let add = fn(x, y){
    x+y;
};

let result = add(five, ten);
```
