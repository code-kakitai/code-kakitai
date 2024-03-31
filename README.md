# Go言語で構築するクリーンアーキテクチャ設計

このリポジトリは[『Go言語で構築するクリーンアーキテクチャ設計』](
https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn)に出てくるサンプルアプリケーションのリポジトリになります。
書籍では一部のコードしか記載できませんでしたが、こちらのリポジトリでより詳しくコードを確認できます。

## この書籍について
[『Go言語で構築するクリーンアーキテクチャ設計』](
https://techbookfest.org/product/9a3U54LBdKDE30ewPS6Ugn)は、Go言語を使用したアプリケーション開発においてクリーンアーキテクチャの原則をどのように適用するかを解説した書籍となります。

以下のような疑問や課題を1つでも持っている方、ぜひ読んでいただきたい本です。

 - クリーンアーキテクチャの概念がいまいち掴めない
 - レイヤーの役割はわかるが、具体的な実装方法が理解できない
 - 各レイヤーでの責務の明確な分担が難しい
 - ドメインやドメインサービスの実装の感覚を掴みたい
 - ユースケースレイヤーでのトランザクション制御に課題を感じている

各レイヤーの実装やそのポイントは書籍にて詳しく書いているため、こちらのリポジトリと合わせて読んでいただければと思います。

## 事前準備
ビルドタグを利用するために、vscodeのsetting.jsonに下記を追加してください。

```json
  "go.toolsEnvVars": {
    "GOFLAGS": "-tags=integration_read,integration_write"
  },
```

## 動作確認
以下の環境で動作確認を行うことができます。

- 初期コマンド
```bash
make init
```
こちらのコマンドで各種コンテナの起動やDBのマイグレーションが行われます。

- サーバー起動
```bash
make run
```
こちらのコマンドで、Goサーバーの起動が行われます。

- Swaggerを用いたAPIドキュメントの確認

```bash
make swagger-up
```
こちらのコマンドで、Swaggerのコンテナが起動します。
localhost:8002にて確認可能です。

## 著者
- [hiroaki](https://twitter.com/hiroaki_u329)
- [sugamaan](https://twitter.com/sugamaan)
- [ぎゅう](https://twitter.com/gyu_outputs)
- [釜地智也](https://twitter.com/tomoya_sakusaku)
