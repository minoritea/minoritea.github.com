+++
title = "slide006"
date = "2017-01-25T20:50:09+09:00"

+++




### reveal.jsのHTML

sectionタグで囲まれた部分が一枚のスライドに相当する

```html
<div class="reveal">
  <div class="slides">
    <section>
      ここにスライド一枚分の内容をHTMLで入れると
      いい感じにスライド表示に変換してくれる
    </section>
    <section>二枚目</section>
    <section>...</section>
  </div>
</div>
<script>Reveal.initialize();</script>
```


