# Newly Acquired Knowledge

## Linting and Correction Tooling
The gometalinter contains a suite of linters and tools to run Go code. it is available at https://github.com/alecthomas/gometalinter.
Gometalinter runs a variety of tools and combines their output in one standard format. Most important tools
in the suite include:
- deadcode
- ineffassign
- misspell
To install the suite, use the command:

```bash
$ go get -u github.com/alecthomas/gometalinter
$ gometalinter --install

# to use 
$ gometalinter --deadline=180s --exclude=vendor --disable=all --enable=misspell ./..
```

