+++
date = "2016-03-22T09:23:52+09:00"
title = "slide15"

+++
## 作り方04

1. SYSROOTに対象サーバのライブラリとヘッダをごっそりコピーする（rsync）

```bash
$ sudo cd $SYSROOT
$ sudo rsync -a root@solaris_sparc_machine:/lib/ $SYSROOT/lib/
$ sudo rsync -a root@solaris_sparc_machine:/usr/lib $SYSROOT/usr/
$ sudo rsync -a root@solaris_sparc_machine:/usr/include $SYSROOT/usr/
$ sudo rsync -a root@solaris_sparc_machine:/usr/platform $SYSROOT/usr/
$ sudo rsync -a root@solaris_sparc_machine:/usr/local/lib $SYSROOT/usr/local/
$ sudo rsync -a root@solaris_sparc_machine:/usr/local/include $SYSROOT/usr/local/
$ sudo rsync -a root@solaris_sparc_machine:/usr/sfw/lib $SYSROOT/usr/sfw/
$ sudo rsync -a root@solaris_sparc_machine:/usr/sfw/include $SYSROOT/usr/sfw/
```
