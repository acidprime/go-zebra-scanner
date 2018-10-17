package snapi

var codeTypes = map[uint16]string{
	1:  "code39",
	2:  "codabar",
	3:  "code128",
	4:  "2of5",
	5:  "iata",
	6:  "2of5int",
	7:  "code93",
	8:  "upc-a",
	9:  "upc-e0",
	10: "ean-8",
	11: "ean-13",
	12: "code11",
	13: "code49",
	14: "msi",
	15: "ean-128",
	16: "upc-e1",
	17: "pdf-417",
	18: "code16k",
	19: "code39ascii",
	20: "upc-d",
	21: "code39tri",
	22: "bookland",
	23: "coupon",
	24: "nw-7",
	25: "isbt-128",
	26: "micropdf",
	27: "datamatrix",
	28: "qr",
	29: "micropdf-cca",
	30: "postnetus",
	31: "planetcode",
	32: "code32",
	33: "isbt-128con",
	34: "postal-jpn",
	35: "postal-aus",
	36: "postal-nld",
	37: "maxicode",
	38: "postal-can",
	39: "postal-gbr",
	40: "macropdf",

	44: "microqr",
	45: "aztec",

	48: "rss-14",
	49: "rss-limited",
	50: "rss-expanded",

	55: "scanlet",

	72: "upc-a-sup2",
	73: "upc-e0-sup2",
	74: "ean-8-sup2",
	75: "ean-13-sup2",

	80: "upc-e1-sup2",
	81: "cca-ean-128",
	82: "cca-ean-13",
	83: "cca-ean-8",
	84: "cca-rss-expanded",
	85: "cca-rss-limited",
	86: "cca-rss-14",
	87: "cca-upc-a",
	88: "cca-upc-e",
	89: "ccc-ean-128",
	90: "tlc-39",

	97:  "ccb-ean-128",
	98:  "ccb-ean-13",
	99:  "ccb-ean-8",
	100: "ccb-rss-expanded",
	101: "ccb-rss-limited",
	102: "ccb-rss-14",
	103: "ccb-upc-a",
	104: "ccb-upc-e",
	105: "signature",
	113: "2of5-matrix",
	114: "2of5-chn",

	136: "upc-a-sup5",
	137: "upc-e0-sup5",
	138: "ean-8-sup5",
	139: "ean-13-sup5",

	144: "upc-e1-sup5",

	154: "macro-micropdf",
}
