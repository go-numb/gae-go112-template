# Static file & templates sample for Google GAE Go1.12
Google Cloud Platform GAE Go1.12にて
Static Dirを定義し、静的ファイルやHTMLを配信したり、Templates Dir以下でgo templatesを使ったりする雛形

- Static DirにHUGOやVueなどをいれて、Go APIなどと連携
- GoやStatic jsからDatastoreやFirestoreなどのGoogle servisesに接続

# Usage
```
$ gcloud config set project <project-id>
$ gcloud app deploy
```

# Errors
```
ERROR: (gcloud.app.deploy) Error Response: [13] An internal error occurred during deployment.
// 既存Projectが残っていると起こるエラーです。
// 新規Projectで実行してください。

```


## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bitbank/blob/master/LICENSE)