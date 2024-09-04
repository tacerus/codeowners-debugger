# CODEOWNERS debugger

Simple tool to check if CODEOWNERS parses as expected.
Mainly intended for [Pistis](https://github.com/SUSE/pistis), as it uses the same Go library.

Install it as a package from [isv:SUSEInfra:Tools](https://build.opensuse.org/package/show/isv:SUSEInfra:Tools/codeowners-debugger) using `zypper in codeowners-debugger` or build it from source using `go build`.

Example usage:

```
codeowners-debugger -directory /path/to/git/repository -file relative/path/to/file/inside/repository
```
