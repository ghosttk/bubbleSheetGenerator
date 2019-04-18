package main

import (
    "baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
	"baliance.com/gooxml/color"
)

func main() {
        doc := document.New()
        table := doc.AddTable()
		// width of the page
		table.Properties().SetWidthPercent(100)
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)
        row := table.AddRow()
        cell:= row.AddCell()
		cell.Properties().SetWidthPercent(100)
		run := cell.AddParagraph().AddRun()
        sheet := "[A][B][C][D]"
        run.AddText(sheet)
        doc.SaveToFile("sheet.docx")
}
