package pdf

import(
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"strconv"
)

//temporary definitions. may be moved or replaced.
type Invoice struct {
	billingId int
	recipient Recipient
	items []BillingItem
	billing Billing
}

type Recipient struct {
	cid int
	fullName string
	street string
	number string
	zip string
	city string
	country string
}

type BillingItem struct {
	id int
	description string
	price_single float64
	amount int
	price_total float64
	discount_perc float64
	discount_abs float64
}

type Billing struct {
	price float64
	price_total float64
	discount_perc float64
	discount_abs float64
}

func InvoiceExample() {

	var rec Recipient = Recipient{0, "Max Mustermann", "Musterstr.", "42", "83454", "München", "Deutschand"}
	var item1 BillingItem = BillingItem{0, "Ski", 30, 2, 54, 10, 6}
	var item2 BillingItem = BillingItem{1, "Stöcke", 5, 4, 18, 10, 2}
	var items []BillingItem = []BillingItem{item1, item2}
	var billing Billing = Billing{80, 72, 10, 8}
	var invoice Invoice = Invoice{0, rec, items, billing}
	var err error = MakeInvoice(invoice, "testInvoice.pdf")

	if err != nil {
		fmt.Println(err)
	}
}

func MakeInvoice(invoice Invoice, fileName string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	pdf.Ln(25);

	//Sender
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 5, tr("peak sport mit system, Fischerweg 2, 87600 Kaufbeuren"))
	pdf.Ln(15);

	//Recipient
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, tr(invoice.recipient.fullName))
	pdf.Ln(5)
	pdf.Cell(40, 10, tr(invoice.recipient.street + " " + invoice.recipient.number))
	pdf.Ln(5)
	pdf.Cell(40, 10, tr(invoice.recipient.zip + " " + invoice.recipient.city))
	pdf.Ln(5)
	pdf.Cell(40, 10, tr(invoice.recipient.country))
	pdf.Ln(30)

	//Headline
	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(190, 10, tr("Rechnung"), "", 0, "C", false, 0, "")
	pdf.Ln(20)

	//Item Table
	//Headlines
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(40, 10, tr("Artikel"), "1", 0, "", false, 0, "")
	pdf.CellFormat(30, 10, tr("Anzahl"), "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, tr("Einzelpreis"), "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, tr("Rabatt"), "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, tr("Preis"), "1", 0, "", false, 0, "")
	pdf.Ln(10)

	//Items
	for _, slice := range invoice.items {
		pdf.CellFormat(40, 10, tr(slice.description), "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 10, tr(strconv.Itoa(slice.amount)), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, tr(strconv.FormatFloat(slice.price_single, 'f', 2, 64) + "€"), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, tr(strconv.FormatFloat(slice.discount_abs, 'f', 2, 64) + "€"), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, tr(strconv.FormatFloat(slice.price_total, 'f', 2, 64) + "€"), "1", 0, "", false, 0, "")
		pdf.Ln(10)
	}

	//Totals
	pdf.CellFormat(190, 1, "", "1", 1, "", false, 0, "")
	pdf.CellFormat(70, 10, "", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, tr(strconv.FormatFloat(invoice.billing.price, 'f', 2, 64) + "€"), "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, tr(strconv.FormatFloat(invoice.billing.discount_abs, 'f', 2, 64) + "€"), "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, tr(strconv.FormatFloat(invoice.billing.price_total, 'f', 2, 64) + "€"), "1", 0, "", false, 0, "")

	return pdf.OutputFileAndClose(fileName)
}