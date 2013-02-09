## servedir.go

    $ go install github.com/zuwiki/servedir
    $ export PATH=$GOPATH/bin:$PATH
    $ cd my_html_bullshit
    $ servedir
    2013/02/08 18:07:44 Listening on:8080
    ^C
    $ cd ..
    $ servedir -endpt=:9001 my_html_bullshit
    2013/02/08 18:08:04 Listening on:9001
    ^C
    $ echo my god, and it is under 20 lines of code! why didn't i write this myself?
    my god, and it is under 20 lines of code! why didn't i write this myself?

