package handler

import "github.com/apex/log"

var ctx = log.WithFields(log.Fields{
	"pkg": "handler",
	"cmd": "proteus-orchestrate",
})

type mapStrStruct map[string]struct{}

// The list of all the supported category codes from the citizenlab test-list.
// See: https://github.com/citizenlab/test-lists/blob/master/lists/00-LEGEND-new_category_codes.csv
var allCatCodes = mapStrStruct{"ALDR": {}, "REL": {}, "PORN": {},
	"PROV": {}, "POLR": {}, "HUMR": {}, "ENV": {}, "MILX": {},
	"HATE": {}, "NEWS": {}, "XED": {}, "PUBH": {}, "GMB": {},
	"ANON": {}, "DATE": {}, "GRP": {}, "LGBT": {}, "FILE": {},
	"HACK": {}, "COMT": {}, "MMED": {}, "HOST": {}, "SRCH": {},
	"GAME": {}, "CULTR": {}, "ECON": {}, "GOVT": {}, "COMM": {},
	"CTRL": {}, "IGO": {}, "MISC": {}}

// Official ISO 3166-2 alpha 2 country codes (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
// We also use the user-assignable ZZ and XX country codes to indicate "Unknown
// country" and "All countries" respectively .
var allCountryCodes = mapStrStruct{
	"AD": {}, "AE": {}, "AF": {}, "AG": {}, "AI": {}, "AL": {},
	"AM": {}, "AN": {}, "AO": {}, "AQ": {}, "AR": {}, "AS": {},
	"AT": {}, "AU": {}, "AW": {}, "AZ": {}, "BA": {}, "BB": {},
	"BD": {}, "BE": {}, "BF": {}, "BG": {}, "BH": {}, "BI": {},
	"BJ": {}, "BM": {}, "BN": {}, "BO": {}, "BR": {}, "BS": {},
	"BT": {}, "BU": {}, "BV": {}, "BW": {}, "BY": {}, "BZ": {},
	"CA": {}, "CC": {}, "CF": {}, "CG": {}, "CH": {}, "CI": {},
	"CK": {}, "CL": {}, "CM": {}, "CN": {}, "CO": {}, "CR": {},
	"CS": {}, "CU": {}, "CV": {}, "CX": {}, "CY": {}, "CZ": {},
	"DD": {}, "DE": {}, "DJ": {}, "DK": {}, "DM": {}, "DO": {},
	"DZ": {}, "EC": {}, "EE": {}, "EG": {}, "EH": {}, "ER": {},
	"ES": {}, "ET": {}, "FI": {}, "FJ": {}, "FK": {}, "FM": {},
	"FO": {}, "FR": {}, "FX": {}, "GA": {}, "GB": {}, "GD": {},
	"GE": {}, "GF": {}, "GH": {}, "GI": {}, "GL": {}, "GM": {},
	"GN": {}, "GP": {}, "GQ": {}, "GR": {}, "GS": {}, "GT": {},
	"GU": {}, "GW": {}, "GY": {}, "HK": {}, "HM": {}, "HN": {},
	"HR": {}, "HT": {}, "HU": {}, "ID": {}, "IE": {}, "IL": {},
	"IN": {}, "IO": {}, "IQ": {}, "IR": {}, "IS": {}, "IT": {},
	"JM": {}, "JO": {}, "JP": {}, "KE": {}, "KG": {}, "KH": {},
	"KI": {}, "KM": {}, "KN": {}, "KP": {}, "KR": {}, "KW": {},
	"KY": {}, "KZ": {}, "LA": {}, "LB": {}, "LC": {}, "LI": {},
	"LK": {}, "LR": {}, "LS": {}, "LT": {}, "LU": {}, "LV": {},
	"LY": {}, "MA": {}, "MC": {}, "MD": {}, "MG": {}, "MH": {},
	"ML": {}, "MN": {}, "MM": {}, "MO": {}, "MP": {}, "MQ": {},
	"MR": {}, "MS": {}, "MT": {}, "MU": {}, "MV": {}, "MW": {},
	"MX": {}, "MY": {}, "MZ": {}, "NA": {}, "NC": {}, "NE": {},
	"NF": {}, "NG": {}, "NI": {}, "NL": {}, "NO": {}, "NP": {},
	"NR": {}, "NT": {}, "NU": {}, "NZ": {}, "OM": {}, "PA": {},
	"PE": {}, "PF": {}, "PG": {}, "PH": {}, "PK": {}, "PL": {},
	"PM": {}, "PN": {}, "PR": {}, "PT": {}, "PW": {}, "PY": {},
	"QA": {}, "RE": {}, "RO": {}, "RU": {}, "RW": {}, "SA": {},
	"SB": {}, "SC": {}, "SD": {}, "SE": {}, "SG": {}, "SH": {},
	"SI": {}, "SJ": {}, "SK": {}, "SL": {}, "SM": {}, "SN": {},
	"SO": {}, "SR": {}, "ST": {}, "SU": {}, "SV": {}, "SY": {},
	"SZ": {}, "TC": {}, "TD": {}, "TF": {}, "TG": {}, "TH": {},
	"TJ": {}, "TK": {}, "TM": {}, "TN": {}, "TO": {}, "TP": {},
	"TR": {}, "TT": {}, "TV": {}, "TW": {}, "TZ": {}, "UA": {},
	"UG": {}, "UM": {}, "US": {}, "UY": {}, "UZ": {}, "VA": {},
	"VC": {}, "VE": {}, "VG": {}, "VI": {}, "VN": {}, "VU": {},
	"WF": {}, "WS": {}, "YD": {}, "YE": {}, "YT": {}, "YU": {},
	"ZA": {}, "ZM": {}, "ZR": {}, "ZW": {}, "XX": {}}
