# sourceview3

Go binding for GtkSourceview 3 on top of gotk3 (https://github.com/gotk3/gotk3).

## Install

It's tested on Ubuntu 18.04 and 20.04. First, install the libgtksourceview
dev package:

```bash
$ sudo apt install libgtksourceview-3.0-dev
```

Then go get the sourceview3 binding library:

```bash
$ go get github.com/linuxerwang/sourceview3
```

## Demo

A very simple demo:

```bash
$ go install github.com/linuxerwang/sourceview3/sourceview3-demo
$ $GOPATH/bin/sourceview3-demo
```

At present most of the GtkSourceBuffer functions do not have bindings, but with
glade and GtkBuilder it's useful enough. Read the demo code to see how to use
it.
