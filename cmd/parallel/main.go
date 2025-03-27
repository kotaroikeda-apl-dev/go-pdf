package main

import (
	"fmt"
	"sync"

	"github.com/jung-kurt/gofpdf"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func generatePDFChunk(users []User, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	pdf := gofpdf.New("P", "mm", "A4", "font")
	pdf.AddUTF8Font("IPA", "", "ipaexg.ttf")
	pdf.AddUTF8Font("SHS", "", "SourceHanSansJP-VF.ttf")
	pdf.AddPage()

	pdf.SetFont("IPA", "", 16)
	pdf.Cell(190, 10, "ユーザーリスト")
	pdf.Ln(15)

	// 表のヘッダー
	pdf.SetFont("IPA", "", 12)
	headers := []string{"ID", "名前", "年齢"}
	widths := []float64{30, 100, 30}

	for i, header := range headers {
		pdf.CellFormat(widths[i], 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// 表のデータ
	pdf.SetFont("SHS", "", 12)
	for _, user := range users {
		row := []string{fmt.Sprintf("%d", user.ID), user.Name, fmt.Sprintf("%d", user.Age)}
		for i, item := range row {
			pdf.CellFormat(widths[i], 10, item, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
	}

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		fmt.Println("エラー:", err)
	}
}

func main() {
	users := make([]User, 10000)
	for i := 1; i <= 10000; i++ {
		users[i-1] = User{i, fmt.Sprintf("User%d", i), 20 + (i % 10)}
	}

	chunkSize := 5000 // 5000行ごとに分割
	var wg sync.WaitGroup

	for i := 0; i < len(users); i += chunkSize {
		end := i + chunkSize
		if end > len(users) {
			end = len(users)
		}
		filename := fmt.Sprintf("users_%d.pdf", i/chunkSize+1)
		wg.Add(1)
		go generatePDFChunk(users[i:end], filename, &wg)
	}

	wg.Wait()
	fmt.Println("PDF 出力完了")
}
