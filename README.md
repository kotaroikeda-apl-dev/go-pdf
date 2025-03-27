## **概要**

## **実行方法**

```sh
go run cmd/basic/main.go # データベースからPDFを出力
go run cmd/layout/main.go # テーブル形式のデータを出力
go run cmd/parallel/main.go # 並列処理でPDFを出力
```

## **学習ポイント**

1. **`gofpdf.New()`** で PDF の用紙サイズ・単位・フォントフォルダを設定できる。
2. **`Cell()`** で列幅・高さを指定しながらテーブル形式のデータを出力できる。
3. 日本語対応フォントは **`AddUTF8Font()`** で登録し、**`SetFont()`** で使用する。
4. ページ番号は **`SetFooterFunc()`** で設定する。
5. テーブルのヘッダーは **`SetFillColor()`** で背景色を設定し、**`CellFormat()`** でテーブル形式のデータを出力する。
6. 並列処理は **`sync.WaitGroup`** で待ち合わせを行う。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
