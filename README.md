# こんな構成の場合のアプリケーションサンプル

・キャッシュサーバとしてdockerの上にredisを配置
・WEBサーバとして、dockerの上にgo+beegoのアプリケーションを配置

## いじっているファイル

以下ファイルをいじってます

```
 conf/app.conf // cacheへの接続情報（cache_）

 cache/* // 接続処理用ライブラリ
```
