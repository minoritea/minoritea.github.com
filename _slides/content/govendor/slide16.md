+++
date = "2016-12-08T19:49:09+09:00"
title = "slide16"

+++



### ステータスの種類

```
    +local    (l) プロジェクト内にあるパッケージ
    +external (e) importされている、かつGOPATHにあるがvendoringされていないパッケージ
    +vendor   (v) vendoringされているパッケージ
    +std      (s) 標準ライブラリ

    +excluded (x) vendoringを禁止しているパッケージ
    +unused   (u) 使われていないパッケージ
    +missing  (m) importされているが、存在しないパッケージ

    +program  (p) mainパッケージ

    +outside  +external +missing
    +all      +all packages
```
