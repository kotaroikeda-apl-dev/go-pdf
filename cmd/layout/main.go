package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "font")
	pdf.AddUTF8Font("IPA", "", "ipaexg.ttf")
	pdf.AddUTF8Font("SHS", "", "SourceHanSansJP-VF.ttf")
	pdf.AliasNbPages("")

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("IPA", "", 10)
		pdf.CellFormat(0, 10,
			fmt.Sprintf("ページ %d / {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AddPage()

	pdf.SetFont("IPA", "", 16)
	pdf.Cell(190, 10, "ユーザーリスト")
	pdf.Ln(15)

	// 表のヘッダー
	pdf.SetFillColor(220, 220, 220)
	pdf.SetFont("IPA", "", 12)
	headers := []string{"ID", "名前", "年齢"}
	widths := []float64{30, 100, 30}

	for i, header := range headers {
		pdf.CellFormat(widths[i], 10, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// 表のデータ
	data := [][]string{
		{"1", "Alice", "30"},
		{"2", "Bob", "25"},
	}

	pdf.SetFont("SHS", "", 12)
	for _, row := range data {
		for i, item := range row {
			pdf.CellFormat(widths[i], 10, item, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
	}

	err := pdf.OutputFileAndClose("users.pdf")
	if err != nil {
		fmt.Println("エラー:", err)
	}
}
