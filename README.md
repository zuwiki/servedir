# servedir

Serve a directory over HTTP. I don't run a real server on my laptop, but often
find myself needing to serve some pages or files.

## Non-features

* Forking/daemonizing (just do `$ servedir &`)
* Authentication
* CGI
* Reverse proxying
* Everything really

## Quick Start

    $ go install github.com/zuwiki/servedir
    $ export PATH=$GOPATH/bin:$PATH
    $ cd my_html_stuff
    $ servedir
    2013/02/08 18:07:44 Listening on :8080
    ^C
    $ cd ..
    $ servedir -endpt=:9001 -dir=my_html_stuff
    2013/02/08 18:08:04 Listening on :9001
    ^C
    $ echo my god, why didn\'t i write this myself?
    my god, why didn't i write this myself?

## Options

    -endpt=:8080
        the interface and port to listen on
    -dir=./
        the directory to serve
    -spa
        usually serve "/index.html" instead of 404s; lazy man's single page app URL rewriting

