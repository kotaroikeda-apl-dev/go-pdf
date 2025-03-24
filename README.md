## **概要**

## **実行方法**

```sh
go run cmd/basic/main.go # データベースからExcelを出力
```

## **学習ポイント**

1. **`gofpdf.New()`** で PDF の用紙サイズ・単位・フォントフォルダを設定できる。
2. **`Cell()`** で列幅・高さを指定しながらテーブル形式のデータを出力できる。
3. 日本語対応フォントは **`AddUTF8Font()`** で登録し、**`SetFont()`** で使用する。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
