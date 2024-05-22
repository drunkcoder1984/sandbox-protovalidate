# 01_buf_confv1
protovalidateのREADME.mdがv1のconfig使ってるのでv1を使った実装
- [protovalidate](https://github.com/bufbuild/protovalidate/blob/v0.6.4/README.md)
- [protovalidate-go](https://github.com/bufbuild/protovalidate-go/blob/v0.6.2/README.md)

main.goでは簡単なvalidateを行いUser.Nameが５文字以上かどうか検証しています

## ディレクトリ構成
```
.
├── buf.yaml        #　bufのmoduleを定義するコンフィグ
├── buf.lock        # buf dep updateで生成される
├── buf.gen.yaml    #　　各言語のコードを生成するコンフィグ
├── proto           # protoファイル置き場
├── gen/go          # ここにprotoファイルを生成する
├── go.mod
├── go.sum
└── main.go
```
## protovalidateを使う手順
### protovalidate-goを追加
```
go get github.com/bufbuild/protovalidate-go
```
### buf.yamlを追加・修正したら依存関係を更新
```
buf dep update
```
公式の手順だと`buf mod update`を使用してるが非推奨とのことなので`buf dep update`を使用する
### protoのgoファイルの生成
```
buf generate
```
