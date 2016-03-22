+++
date = "2016-03-22T10:02:54+09:00"
title = "slide18"

+++
## コンパイル

```bash
$ go build --compiler gccgo \
  --gccgoflags "-v -static-libgo -static-libgcc -Wl,-dy -lnsl -lsocket -lrt"
```

**重要**
参照元のオプションに`-static-libgcc`を追加しないとSolaris環境では動作しませんでした。
