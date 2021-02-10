# eagain.net/go/entropy -- Output random bytes aka entropy

`entropy` is like a saner, quicker, easier, safer and more memorable `dd if=/dev/urandom bs=1 count=BYTES`

(Plus it'll even work even where `/dev/urandom` doesn't exist, by using the `getrandom` syscall.)

```
go install eagain.net/go/entropy
```

Examples:

```
$ entropy 32 >secret
$ entropy 32
entropy: stdout is a terminal, refusing to output binary
$ entropy 32 | zbase32-encode
k6dsmurfj3he3qrmdc99z1pzdrm4ipgqmzkk9cbp4w8zubfkg3ko
```
