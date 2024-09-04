# CODEOWNERS debugger

Simple tool to check if CODEOWNERS parses as expected.
Mainly intended for [Pistis](https://github.com/SUSE/pistis), as it uses the same Go library.

```
go build
./codeowners-debugger -directory /path/to/git/repository -file relative/path/to/file/inside/repository
```
