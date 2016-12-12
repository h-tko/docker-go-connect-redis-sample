<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
 <html>
    <head>
    </head>
    <body>
        <div class="container">
            <div class="card">
                <div class="card-header card-primary">
                    Redis cache read write サンプル.
                </div>
                <div class="card-block">
                    <form action="/" method="post">
                        <div class="alert alert-success"><strong>現在のcache: </strong> {{.NowCache}}</div>

                        <div class="input-group">
                            <input class="form-control" type="text" name="Item">

                            <div class="input-group-btn">
                                <button class="btn btn-default" type="submit">登録</button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>

        <script type="text/javascript" src="/static/dist/bundle.js"></script>
    </body>
 </html>
