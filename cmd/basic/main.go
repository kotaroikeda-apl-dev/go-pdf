package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "font") // フォントのあるフォルダを指定
	pdf.AddUTF8Font("IPA", "", "ipaexg.ttf")   // フォント名・スタイル・ファイル名
	pdf.AddPage()
	pdf.SetFont("IPA", "", 16)

	// ヘッダー行
	pdf.Cell(60, 10, "商品名")
	pdf.Cell(30, 10, "単価")
	pdf.Cell(30, 10, "数量")
	pdf.Ln(-1)

	// データ行
	pdf.Cell(60, 10, "りんご")
	pdf.Cell(30, 10, "100円")
	pdf.Cell(30, 10, "3個")
	pdf.Ln(-1)

	pdf.Cell(60, 10, "みかん")
	pdf.Cell(30, 10, "80円")
	pdf.Cell(30, 10, "5個")
	pdf.Ln(-1)

	err := pdf.OutputFileAndClose("report.pdf")
	if err != nil {
		fmt.Println("エラー:", err)
	}

	fmt.Println("PDF ファイル作成完了")
}
