package modifiers

func ModifyText(text string) string {
	text = ConvertHexAndBinToDecimal(text)
	text = ConvertUpLowCap(text)
	text = ConvertAToAn(text)
	text = FormatText(text)
	text = HandleQuotes(text)
	return text
}
