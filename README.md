# Chronicler - A Storytelling web toolkit.

Chronicler is a toolkit for web request routing in Go, it's very heavily influenced by [Cuba](http://cuba.is/), a fantastic web microframework for Ruby. What separates Chronicler and Cuba from the bulk of other web toolkits (at least in Rubyland and Goland) is the notion that the ramifications of the routing tree can be determined by arbitrary code and conditions as opposed to being strictly bound to evaluating the path of an http request or it's verb. This is a very powerful idea.

This project started with the intention of being a simple port of Cuba but the differences in the languages they are implemented in forced me to adapt Cuba's ideas to a more idiomatic approach, Chronicler is the result of that effort.

**DISCLAIMER**: This is right now at a "proof of concept" stage, please don't go sending it around to people until I can make a proper release, but do open issues or ping me in irc for feedback, I hang out at #lesscode in freenode :).

## The Basics

The concept of Chronicler is simple: your application will be composed of **nodes** and **routes**, nodes group routes which represent both a set of matching conditions and the code to be executed if these conditions are met.

A very basic Chronicler application will look like this:

```go
package main

import(
  "github.com/pote/chronicler"
  "io"
)

func main() {
  mainNode := chronicler.NewNode()
  mainNode.Register(&home{})

  mainNode.Perform(":8080")
}

type home struct { }

func (r *home) Match(req *http.Request) bool {
  return true
}

func (r *home) Perform(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "It was night again. The Waystone Inn lay in silence, and it was a silence of three parts.\n")
}
```

Hitting `http://localhost:8080/` should then yield the first sentence of [The Name of the Wind](http://www.amazon.com/Name-Wind-Kingkiller-Chronicle-Day-ebook/dp/B0010SKUYM/ref=sr_1_1?ie=UTF8&qid=1413554542&sr=8-1&keywords=The+Name+of+the+Wind), easy as pie.
