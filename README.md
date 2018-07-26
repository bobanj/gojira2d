# gojira2d

[![Build Status](https://travis-ci.org/markov/gojira2d.svg?branch=master)](https://travis-ci.org/markov/gojira2d)
[![Go Report Card](https://goreportcard.com/badge/github.com/markov/gojira2d)](https://goreportcard.com/report/github.com/markov/gojira2d)
[![Join the chat at https://gitter.im/gojira2d/Lobby](https://badges.gitter.im/gojira2d/Lobby.svg)](https://gitter.im/gojira2d/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Simple 2D game library written in Go and based on modern OpenGL

## Installation

Install Golang and GLFW:

    $ brew install go glfw

Setup your [`$GOPATH`](https://golang.org/doc/code.html#GOPATH) and clone the
repository into `$GOPATH/src` folder:

    $ go get -u markov/gojira2d
    $ cd $GOPATH/src/markov/gojira2d
    $ git remote set-url origin git@github.com:markov/gojira2d.git

Install Golang dependencies:

    $ go get \
        github.com/go-gl/mathgl/mgl32 \
        github.com/go-gl/gl/v4.1-core/gl \
        github.com/go-gl/glfw/v3.2/glfw

Try running some examples:

    $ go run examples/quad/main.go
    ...
