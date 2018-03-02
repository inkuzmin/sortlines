- `sortlines_linux_amd64` is a binary for Linux
- `sortlines_darwin_amd64` is a binary for MacOS

```
$ chmod +x  sortlines_linux_amd64
$ echo "a b c\nb c a\nc a b" | ./sortlines_linux_amd64
$ # or
$ ./sortlines_linux_amd64 -i test.txt
```
