# 02_buf_confv2
01_buf_confv1の`buf.yaml`と`buf.gen.yaml`をv2したもの

bufの公式に[移行ガイド](https://buf.build/docs/migration-guides/migrate-v2-config-files)があるのでそこに乗ってる移行コマンドを実行

## 移行コマンド
簡単なコンフィグなのでこのコマンド実行して修正しないで`buf generate`が実行できた
```
buf config migrate
```
