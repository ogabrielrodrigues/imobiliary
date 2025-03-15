package worker

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/util"
)

func Work(page *rod.Page, url string, id string) []kind.Debit {
	page.MustNavigate(url)

	page.MustWaitDOMStable()

	// Type House ID on input.
	page.MustElement("#txtcod").MustInput(id)
	time.Sleep(time.Millisecond * 500)

	// Click in "Access" button.
	page.MustElement("[name=Acessar]").MustClick()
	time.Sleep(time.Millisecond * 500)

	page.MustWaitDOMStable()

	// Navigate and click to section "Débitos"
	page.MustElement("li.debitos").MustClick()
	time.Sleep(time.Millisecond * 500)

	page.MustWaitDOMStable()

	if exist := page.MustHas(".textoTituloImpressao7"); exist {
		util.Logln(util.ColorGreen, "✓ Skipped")
		return nil
	}

	debts := page.MustElements(".bordaServicos table tbody tr")
	debts = debts[:len(debts)-7]

	var found_debts []kind.Debit
	for _, debit := range debts {
		columns := debit.MustElements("td.textocampo")

		found_debts = append(found_debts, kind.Debit{
			MonthRef:  columns[0].MustText(),
			Value:     util.ParseValue(columns[2].MustText()),
			ExpiresIn: columns[3].MustText(),
		})
	}

	util.Logln(util.ColorGreen, "✓ OK")
	return found_debts
}
