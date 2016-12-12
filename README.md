# こんな構成の場合のアプリケーションサンプル

・キャッシュサーバとしてdockerの上にredisを配置
・WEBサーバとして、dockerの上にgo+beegoのアプリケーションを配置

## いじっているファイル

以下ファイルをいじってます

```
 conf/app.conf // cacheへの接続情報（cache_）

 cache/* // 接続処理用ライブラリ
```

## 必要なライブラリ

beegoで作っているので以下go getする必要があります。

```
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
$ go get github.com/astaxie/beego/cache
$ go get github.com/astaxie/beego/cache/redis
```
