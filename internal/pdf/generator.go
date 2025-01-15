// internal/pdf/generator.go
package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

type EstimateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func pdfGenerator(filename string, estimate EstimateRequest) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 16)
	pdf.Cell(40, 10, estimate.Title)
	pdf.Ln(-1)
	pdf.SetFont("", "", 12)
	pdf.Write(8, estimate.Content)

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Fatalf("Ошибка при сохранении файла: %v\n", err)
	}
	fmt.Printf("PDF файл сохранён: %s\n", filename)
}
func sendToWhatsApp(fileName string) {
	// Здесь будет логика для отправки файла через WhatsApp
	// Например, использование WhatsApp Business API
}
