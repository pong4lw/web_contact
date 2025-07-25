# Contact Page Scraper

企業サイト一覧から **「お問い合わせページ」へのリンク** を自動的に抽出し、結果をCSVファイルに保存します。

## 📝 概要

- 指定されたURL一覧（コード内にURLの一覧を記載）からWebページを取得
- 各HTML内の `<a>` タグをスキャンして、「contact」または「お問い合わせ」を含むリンクを検索
- 検出された問い合わせページのURLを `results.csv` に出力

## 📦 出力形式

出力ファイル: `results.csv`

| Company URL                        | Contact Page URL                 |
|-----------------------------------|----------------------------------|
| https://example-company.com       | https://example-company.com/contact |
| https://company2.co.jp            |                                  |
| ...                               | ...                              |
---

※ 見つからなかった場合は空欄になります。

## 🔧 必要条件

### 1. Go 環境

- Go 1.18 以上推奨  
- Mac / Linux / Windows 対応

▶️ インストール & 実行方法
1. 必要ライブラリのインストール
ターミナルで以下のコマンドを実行して、依存パッケージをインストールしてください：

```bash
go get github.com/PuerkitoBio/goquery
go get github.com/corpix/uarand
```
または go mod を使っている場合は、プロジェクトフォルダで以下を実行：

```bash
go mod init contact-scraper
go get github.com/PuerkitoBio/goquery
go get github.com/corpix/uarand
```

### 実行時
main.goの下記関数にURLの一覧を記入ください

```bash
（記入例）
var companyList = []string{
	"https://example-company1.com",
	"https://example-company2.co.jp",
	"https://example-company3.jp",
}
```

```bash
go run main.go
```
