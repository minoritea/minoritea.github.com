
+++
date = "2010-11-05"
draft = true
title = "SafariのLocationエンコードについて2"
tags  = ['javascript','disqus','safari']
+++

前回、SafariでエスケープされたURLのサイトを表示したときに、location.hrefでアドレスを取得するとデコードした文字列が取得されてしまう問題について、webkitのブラウザを判別するスクリプトを書いて対応した。
その後、ChromeだとURLがきちんとエスケープされて取得されることが分かった。僕はchromeは普段使わないのでこの仕様の違いに気付かなかったのである。

どうやって判別するか悩んでいたのだが、結局decodeURLでエスケープされていないURLをデコードしても結果が変わらないことに着目して以下の対応とした。


  var disqus_url = encodeURI(decodeURI(window.location.href));


エンコードされたURL（Firefox,Chrome,etc）だけがdecodeURIされ、その後再びencodeURIされるので正しくエスケープされたURLが取得出来る。
URLの下位パスをタイトルから生成していたのでこのような問題となったのだから、そこを改めた方がよいのかもしれない。	

	
