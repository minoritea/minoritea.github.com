+++
date = "2017-01-25T20:50:09+09:00"
title = "slide004"

+++




`layouts/_default/list.html`

- デフォルトのテンプレートでは見出しを表示するが、各ページの内容を全部表示するように変更
- Hugo のヘルパーメソッドでsection内の各ページの内容を取り出す
- 取り出した各コンテンツはsectionタグで囲む（後述）

```
      {{ range .Data.Pages.ByTitle }}
        <section>
          {{ .Content }}
        </section>
      {{ end }}
```

記法はgo-template


