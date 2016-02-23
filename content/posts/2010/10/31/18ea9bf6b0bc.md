
+++
date = "2010-10-31"
draft = true
title = "SafariのLocationエンコードについて"
tags  = ['javascript','disqus','safari']
+++

やっぱりDisqusがSafariでうまく動作しない。どうも記事ごとにうまく読み込めるかエラーを起こすが決まっているようなので調べるとうまく動作しないのはタイトルがマルチバイトの記事だけだった。

このブログの各記事のリンクは全てURLをエスケープしてあるのだが、Safariでwindow.location.hrefを取得するとエスケープ前のURLになってしまうのがエラーの原因であるようだ。ちなみにFirefoxでは再現しない。何故このような仕様になっているのか理解に苦しむのだがご存知の方がいたら教えてもらいたい。（バグなんだろうか）。

ちなみに以下のコードを埋め込むことで一応Safariでのエラーは無くなった。あまりスマートなコードではないが仕方が無い。


if(!jQuery.support.checkOn){
  var disqus_url = encodeURI(window.location.href);
}


ちなみにjQueryでのブラウザ判別については「jQuery.supportでのブラウザ判別」を参照した。分かりやすいまとめで助かった。	

	
