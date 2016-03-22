+++
date = "2016-03-22T09:56:31+09:00"
title = "slide16"

+++
## 作り方05

1. 対象アーキテクチャ向けにbinutilsとgccをビルドする

```bash
$ tar -zxf binutils-2.25.tar.gz
$ mkdir build-binutils
$ cd build-binutils
$ sudo ../binutils-2.25/configure  -target=$TARGET --prefix=$PREFIX -with-sysroot=$SYSROOT -v
$ sudo make -j4
$ sudo make install
```

```bash
$ untar gcc-5.1.0.tar.gz
$ mkdir gcc-build
$ cd gcc-build
$ sudo ../gcc-5.1.0/configure --target=$TARGET --with-gnu-as --with-gnu-ld  --prefix=$PREFIX -with-sysroot=$SYSROOT --disable-libgcj --enable-languages=c,c++,go -v
$ sudo make -j6
$ sudo make install
```

**完成**
