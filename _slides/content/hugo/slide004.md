+++
date = "2017-01-25T20:50:09+09:00"
title = "slide004"

+++




`layouts/_default/list.html`

- 通常は短縮表示するが、コンテンツをループで全表示する
- sectionタグで囲む（後述）

```
      {{ range .Data.Pages.ByTitle }}
        <section>
          {{ .Content }}
        </section>
      {{ end }}
```

記法はgo-template


