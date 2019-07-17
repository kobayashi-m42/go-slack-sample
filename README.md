# go-slack-sample
Slackにメッセージを送信するサンプルプロジェクト

## Incoming Webhookの追加
SlackのIncoming Webhookを利用して、Slackにメッセージを送信します。

下記よりIncoming Webhookの設定を追加し、webhook用のURLを取得します。

https://slack.com/services/new/incoming-webhook

## Dockerでの実行

### 環境変数の設定
`.env`を作成します。

```
WEBHOOK_URL=https://hooks.slack.com/services/… // 上記で取得したWebhook URL
```

### 初回起動
```
docker-compose up --build -d
```

### 2回目以降の起動
```
docker-compose up
```

`docker-compose.yml`の`command`で指定している`from, docker`がSlackに送信されます。

## Mac上での実行

### 環境変数の設定
`.envrc`を作成します。

```
export WEBHOOK_URL=https://hooks.slack.com/services/… // 上記で取得したWebhook URL
```

## 実行
`go run`で実行します。引数で渡している`message`がSlackに送信されます。


```
$ go run main.go message
```
