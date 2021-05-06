package sib

import (
	"strconv"
)

// Country represents the country as it is recorded in the Travel App.
type Country struct {
	//nolint
	name   string
	alpha2 string
	//nolint
	countryCode string
	sibCode     int
}

// FindCountryBySibCode returns the country code for a given sibCode.
// Sib uses numbers to identify countries. These numbers are just the index
// in a list. If no number appears, we return a default string. We use US, because
// expect most people to be travelling to and from the US.
func FindCountryBySibCode(sibCode string) string {
	cd, err := strconv.Atoi(sibCode)
	if err != nil {
		return "US"
	}
	var code = "US"
	for _, c := range sibCountries {
		if c.sibCode == cd {
			code = c.alpha2
		}
	}
	return code
}

// FindCountryByName returns the country code for a given string.
// This is mostly here for mapping SIB's nationality field. They are using
// strings in this field, but it does not appear to be standardized.
// We use US, because expect most people to be travelling to and from the US.
func FindCountryByName(name string) string {
	var n = "US"
	for _, c := range sibCountries {
		if strconv.Itoa(c.sibCode) == name {
			n = c.alpha2
		}
	}
	return n
}

//nolint
var sibCountries = []Country{
	{
		name:        "Afghanistan",
		alpha2:      "AF",
		countryCode: "004",
		sibCode:     1,
	},

	{
		name:        "Albania",
		alpha2:      "AL",
		countryCode: "008",
		sibCode:     2,
	},
	{
		name:        "Algeria",
		alpha2:      "DZ",
		countryCode: "012",
		sibCode:     3,
	},
	{
		name:        "American Samoa",
		alpha2:      "AS",
		countryCode: "016",
		sibCode:     4,
	},
	{
		name:        "Andorra",
		alpha2:      "AD",
		countryCode: "020",
		sibCode:     5,
	},
	{
		name:        "Angola",
		alpha2:      "AO",
		countryCode: "024",
		sibCode:     6,
	},
	{
		name:        "Anguilla",
		alpha2:      "AI",
		countryCode: "660",
		sibCode:     7,
	},
	{
		name:        "Antarctica",
		alpha2:      "AQ",
		countryCode: "010",
		sibCode:     8,
	},
	{
		name:        "Antigua and Barbuda",
		alpha2:      "AG",
		countryCode: "028",
		sibCode:     9,
	},
	{
		name:        "Argentina",
		alpha2:      "AR",
		countryCode: "032",
		sibCode:     10,
	},
	{
		name:        "Armenia",
		alpha2:      "AM",
		countryCode: "051",
		sibCode:     11,
	},
	{
		name:        "Aruba",
		alpha2:      "AW",
		countryCode: "533",
		sibCode:     12,
	},
	{
		name:        "Australia",
		alpha2:      "AU",
		countryCode: "036",
		sibCode:     13,
	},
	{
		name:        "Austria",
		alpha2:      "AT",
		countryCode: "040",
		sibCode:     14,
	},
	{
		name:        "Azerbaijan",
		alpha2:      "AZ",
		countryCode: "031",
		sibCode:     15,
	},
	{
		name:        "Bahamas",
		alpha2:      "BS",
		countryCode: "044",
		sibCode:     16,
	},
	{
		name:        "Bahrain",
		alpha2:      "BH",
		countryCode: "048",
		sibCode:     17,
	},
	{
		name:        "Bangladesh",
		alpha2:      "BD",
		countryCode: "050",
		sibCode:     18},
	{
		name:        "Barbados",
		alpha2:      "BB",
		countryCode: "052",
		sibCode:     19,
	},
	{
		name:        "Belarus",
		alpha2:      "BY",
		countryCode: "112",
		sibCode:     20,
	},
	{
		name:        "Belgium",
		alpha2:      "BE",
		countryCode: "056",
		sibCode:     21,
	},
	{
		name:        "Belize",
		alpha2:      "BZ",
		countryCode: "084",
		sibCode:     22,
	},
	{
		name:        "Benin",
		alpha2:      "BJ",
		countryCode: "204",
		sibCode:     23,
	},
	{
		name:        "Bermuda",
		alpha2:      "BM",
		countryCode: "060",
		sibCode:     24,
	},
	{
		name:        "Bhutan",
		alpha2:      "BT",
		countryCode: "064",
		sibCode:     25,
	},
	{
		name:        "Bolivia (Plurinational State of)",
		alpha2:      "BO",
		countryCode: "068",
		sibCode:     26,
	},
	{
		name:        "Bonaire, Sint Eustatius and Saba",
		alpha2:      "BQ",
		countryCode: "535",
		sibCode:     27,
	},
	{
		name:        "Bonaire, Sint Eustatius and Saba",
		alpha2:      "BQ",
		countryCode: "535",
		sibCode:     252,
	},
	{
		name:        "Bosnia and Herzegovina",
		alpha2:      "BA",
		countryCode: "070",
		sibCode:     28,
	},
	{
		name:        "Botswana",
		alpha2:      "BW",
		countryCode: "072",
		sibCode:     29,
	},
	{
		name:        "Bouvet Island",
		alpha2:      "BV",
		countryCode: "074",
		sibCode:     30,
	},
	{
		name:        "Brazil",
		alpha2:      "BR",
		countryCode: "076",
		sibCode:     31,
	},
	{
		name:        "British Indian Ocean Territory",
		alpha2:      "IO",
		countryCode: "086",
		sibCode:     32,
	},
	{
		name:        "Brunei Darussalam",
		alpha2:      "BN",
		countryCode: "096",
		sibCode:     33,
	},
	{
		name:        "Bulgaria",
		alpha2:      "BG",
		countryCode: "100",
		sibCode:     34,
	},
	{
		name:        "Burkina Faso",
		alpha2:      "BF",
		countryCode: "854",
		sibCode:     35,
	},
	{
		name:        "Burundi",
		alpha2:      "BI",
		countryCode: "108",
		sibCode:     36,
	},
	{
		name:        "Cabo Verde",
		alpha2:      "CV",
		countryCode: "132",
		sibCode:     37,
	},
	{
		name:        "Cambodia",
		alpha2:      "KH",
		countryCode: "116",
		sibCode:     38,
	},
	{
		name:        "Cameroon",
		alpha2:      "CM",
		countryCode: "120",
		sibCode:     39,
	},
	{
		name:        "Canada",
		alpha2:      "CA",
		countryCode: "124",
		sibCode:     40,
	},
	{
		name:        "Cayman Islands",
		alpha2:      "KY",
		countryCode: "136",
		sibCode:     41,
	},
	{
		name:        "Central African Republic",
		alpha2:      "CF",
		countryCode: "140",
		sibCode:     42,
	},
	{
		name:        "Chad",
		alpha2:      "TD",
		countryCode: "148",
		sibCode:     43,
	},
	{
		name:        "Chile",
		alpha2:      "CL",
		countryCode: "152",
		sibCode:     44,
	},
	{
		name:        "China",
		alpha2:      "CN",
		countryCode: "156",
		sibCode:     45,
	},
	{
		name:        "Christmas Island",
		alpha2:      "CX",
		countryCode: "162",
		sibCode:     46,
	},
	{
		name:        "Cocos (Keeling) Islands",
		alpha2:      "CC",
		countryCode: "166",
		sibCode:     47,
	},
	{
		name:        "Colombia",
		alpha2:      "CO",
		countryCode: "170",
		sibCode:     48,
	},
	{
		name:        "Comoros",
		alpha2:      "KM",
		countryCode: "174",
		sibCode:     49,
	},
	{
		name:        "Congo",
		alpha2:      "CG",
		countryCode: "178",
		sibCode:     51,
	},
	{
		name:        "Congo, Democratic Republic of the",
		alpha2:      "CD",
		countryCode: "180",
		sibCode:     50,
	},
	{
		name:        "Cook Islands",
		alpha2:      "CK",
		countryCode: "184",
		sibCode:     52,
	},
	{
		name:        "Costa Rica",
		alpha2:      "CR",
		countryCode: "188",
		sibCode:     53,
	},
	{
		name:        "Cote d'Ivoire",
		alpha2:      "CI",
		countryCode: "384",
		sibCode:     59,
	},
	{
		name:        "Croatia",
		alpha2:      "HR",
		countryCode: "191",
		sibCode:     54,
	},
	{
		name:        "Cuba",
		alpha2:      "CU",
		countryCode: "192",
		sibCode:     55,
	},
	{
		name:        "Curacao",
		alpha2:      "CW",
		countryCode: "531",
		sibCode:     56,
	},
	{
		name:        "Cyprus",
		alpha2:      "CY",
		countryCode: "196",
		sibCode:     57,
	},
	{
		name:        "Czechia",
		alpha2:      "CZ",
		countryCode: "203",
		sibCode:     58,
	},
	{
		name:        "Denmark",
		alpha2:      "DK",
		countryCode: "208",
		sibCode:     60,
	},
	{
		name:        "Djibouti",
		alpha2:      "DJ",
		countryCode: "262",
		sibCode:     61,
	},
	{
		name:        "Dominica",
		alpha2:      "DM",
		countryCode: "212",
		sibCode:     62,
	},
	{
		name:        "Dominican Republic",
		alpha2:      "DO",
		countryCode: "214",
		sibCode:     63,
	},
	{
		name:        "Ecuador",
		alpha2:      "EC",
		countryCode: "218",
		sibCode:     64,
	},
	{
		name:        "Egypt",
		alpha2:      "EG",
		countryCode: "818",
		sibCode:     65,
	},
	{
		name:        "El Salvador",
		alpha2:      "SV",
		countryCode: "222",
		sibCode:     66,
	},
	{
		name:        "Equatorial Guinea",
		alpha2:      "GQ",
		countryCode: "226",
		sibCode:     67,
	},
	{
		name:        "Eritrea",
		alpha2:      "ER",
		countryCode: "232",
		sibCode:     68,
	},
	{
		name:        "Estonia",
		alpha2:      "EE",
		countryCode: "233",
		sibCode:     69,
	},
	{
		name:        "Eswatini",
		alpha2:      "SZ",
		countryCode: "748",
		sibCode:     70,
	},
	{
		name:        "Ethiopia",
		alpha2:      "ET",
		countryCode: "231",
		sibCode:     71,
	},
	{
		name:        "Falkland Islands (Malvinas)",
		alpha2:      "FK",
		countryCode: "238",
		sibCode:     72,
	},
	{
		name:        "Faroe Islands",
		alpha2:      "FO",
		countryCode: "234",
		sibCode:     73,
	},
	{
		name:        "Fiji",
		alpha2:      "FJ",
		countryCode: "242",
		sibCode:     74,
	},
	{
		name:        "Finland",
		alpha2:      "FI",
		countryCode: "246",
		sibCode:     75,
	},
	{
		name:        "France",
		alpha2:      "FR",
		countryCode: "250",
		sibCode:     76,
	},
	{
		name:        "French Guiana",
		alpha2:      "GF",
		countryCode: "254",
		sibCode:     77,
	},
	{
		name:        "French Polynesia",
		alpha2:      "PF",
		countryCode: "258",
		sibCode:     78,
	},
	{
		name:        "French Southern Territories",
		alpha2:      "TF",
		countryCode: "260",
		sibCode:     79,
	},
	{
		name:        "Gabon",
		alpha2:      "GA",
		countryCode: "266",
		sibCode:     80,
	},
	{
		name:        "Gambia",
		alpha2:      "GM",
		countryCode: "270",
		sibCode:     81,
	},
	{
		name:        "Georgia",
		alpha2:      "GE",
		countryCode: "268",
		sibCode:     82,
	},
	{
		name:        "Germany",
		alpha2:      "DE",
		countryCode: "276",
		sibCode:     83,
	},
	{
		name:        "Ghana",
		alpha2:      "GH",
		countryCode: "288",
		sibCode:     84,
	},
	{
		name:        "Gibraltar",
		alpha2:      "GI",
		countryCode: "292",
		sibCode:     85,
	},
	{
		name:        "Greece",
		alpha2:      "GR",
		countryCode: "300",
		sibCode:     86,
	},
	{
		name:        "Greenland",
		alpha2:      "GL",
		countryCode: "304",
		sibCode:     87,
	},
	{
		name:        "Grenada",
		alpha2:      "GD",
		countryCode: "308",
		sibCode:     88,
	},
	{
		name:        "Guadeloupe",
		alpha2:      "GP",
		countryCode: "312",
		sibCode:     89,
	},
	{
		name:        "Guam",
		alpha2:      "GU",
		countryCode: "316",
		sibCode:     90,
	},
	{
		name:        "Guatemala",
		alpha2:      "GT",
		countryCode: "320",
		sibCode:     91,
	},
	{
		name:        "Guernsey",
		alpha2:      "GG",
		countryCode: "831",
		sibCode:     92,
	},
	{
		name:        "Guinea",
		alpha2:      "GN",
		countryCode: "324",
		sibCode:     93,
	},
	{
		name:        "Guinea-Bissau",
		alpha2:      "GW",
		countryCode: "624",
		sibCode:     94,
	},
	{
		name:        "Guyana",
		alpha2:      "GY",
		countryCode: "328",
		sibCode:     95,
	},
	{
		name:        "Haiti",
		alpha2:      "HT",
		countryCode: "332",
		sibCode:     96,
	},
	{
		name:        "Heard Island and McDonald Islands",
		alpha2:      "HM",
		countryCode: "334",
		sibCode:     97,
	},
	{
		name:        "Holy See",
		alpha2:      "VA",
		countryCode: "336",
		sibCode:     98,
	},
	{
		name:        "Honduras",
		alpha2:      "HN",
		countryCode: "340",
		sibCode:     99,
	},
	{
		name:        "Hong Kong",
		alpha2:      "HK",
		countryCode: "344",
		sibCode:     100,
	},
	{
		name:        "Hungary",
		alpha2:      "HU",
		countryCode: "348",
		sibCode:     101,
	},
	{
		name:        "Iceland",
		alpha2:      "IS",
		countryCode: "352",
		sibCode:     102,
	},
	{
		name:        "India",
		alpha2:      "IN",
		countryCode: "356",
		sibCode:     103,
	},
	{
		name:        "Indonesia",
		alpha2:      "ID",
		countryCode: "360",
		sibCode:     104,
	},
	{
		name:        "Iran (Islamic Republic of)",
		alpha2:      "IR",
		countryCode: "364",
		sibCode:     105,
	},
	{
		name:        "Iraq",
		alpha2:      "IQ",
		countryCode: "368",
		sibCode:     106,
	},
	{
		name:        "Ireland",
		alpha2:      "IE",
		countryCode: "372",
		sibCode:     107,
	},
	{
		name:        "Isle of Man",
		alpha2:      "IM",
		countryCode: "833",
		sibCode:     108,
	},
	{
		name:        "Israel",
		alpha2:      "IL",
		countryCode: "376",
		sibCode:     109,
	},
	{
		name:        "Italy",
		alpha2:      "IT",
		countryCode: "380",
		sibCode:     110,
	},
	{
		name:        "Jamaica",
		alpha2:      "JM",
		countryCode: "388",
		sibCode:     111,
	},
	{
		name:        "Japan",
		alpha2:      "JP",
		countryCode: "392",
		sibCode:     112,
	},
	{
		name:        "Jersey",
		alpha2:      "JE",
		countryCode: "832",
		sibCode:     113,
	},
	{
		name:        "Jordan",
		alpha2:      "JO",
		countryCode: "400",
		sibCode:     114,
	},
	{
		name:        "Kazakhstan",
		alpha2:      "KZ",
		countryCode: "398",
		sibCode:     115,
	},
	{
		name:        "Kenya",
		alpha2:      "KE",
		countryCode: "404",
		sibCode:     116,
	},
	{
		name:        "Kiribati",
		alpha2:      "KI",
		countryCode: "296",
		sibCode:     117,
	},
	{
		name:        "Korea (Democratic People's Republic of)",
		alpha2:      "KP",
		countryCode: "408",
		sibCode:     118,
	},
	{
		name:        "Korea, Republic of",
		alpha2:      "KR",
		countryCode: "410",
		sibCode:     119,
	},
	{
		name:        "Kuwait",
		alpha2:      "KW",
		countryCode: "414",
		sibCode:     120,
	},
	{
		name:        "Kyrgyzstan",
		alpha2:      "KG",
		countryCode: "417",
		sibCode:     121,
	},
	{
		name:        "Lao People's Democratic Republic",
		alpha2:      "LA",
		countryCode: "418",
		sibCode:     122,
	},
	{
		name:        "Latvia",
		alpha2:      "LV",
		countryCode: "428",
		sibCode:     123,
	},
	{
		name:        "Lebanon",
		alpha2:      "LB",
		countryCode: "422",
		sibCode:     124,
	},
	{
		name:        "Lesotho",
		alpha2:      "LS",
		countryCode: "426",
		sibCode:     125,
	},
	{
		name:        "Liberia",
		alpha2:      "LR",
		countryCode: "430",
		sibCode:     126,
	},
	{
		name:        "Libya",
		alpha2:      "LY",
		countryCode: "434",
		sibCode:     127,
	},
	{
		name:        "Liechtenstein",
		alpha2:      "LI",
		countryCode: "438",
		sibCode:     128,
	},
	{
		name:        "Lithuania",
		alpha2:      "LT",
		countryCode: "440",
		sibCode:     129,
	},
	{
		name:        "Luxembourg",
		alpha2:      "LU",
		countryCode: "442",
		sibCode:     130,
	},
	{
		name:        "Macao",
		alpha2:      "MO",
		countryCode: "446",
		sibCode:     131,
	},
	{
		name:        "Madagascar",
		alpha2:      "MG",
		countryCode: "450",
		sibCode:     132,
	},
	{
		name:        "Malawi",
		alpha2:      "MW",
		countryCode: "454",
		sibCode:     133,
	},
	{
		name:        "Malaysia",
		alpha2:      "MY",
		countryCode: "458",
		sibCode:     134,
	},
	{
		name:        "Maldives",
		alpha2:      "MV",
		countryCode: "462",
		sibCode:     135,
	},
	{
		name:        "Mali",
		alpha2:      "ML",
		countryCode: "466",
		sibCode:     136,
	},
	{
		name:        "Malta",
		alpha2:      "MT",
		countryCode: "470",
		sibCode:     137,
	},
	{
		name:        "Marshall Islands",
		alpha2:      "MH",
		countryCode: "584",
		sibCode:     138,
	},
	{
		name:        "Martinique",
		alpha2:      "MQ",
		countryCode: "474",
		sibCode:     139,
	},
	{
		name:        "Mauritania",
		alpha2:      "MR",
		countryCode: "478",
		sibCode:     140,
	},
	{
		name:        "Mauritius",
		alpha2:      "MU",
		countryCode: "480",
		sibCode:     141,
	},
	{
		name:        "Mayotte",
		alpha2:      "YT",
		countryCode: "175",
		sibCode:     142,
	},
	{
		name:        "Mexico",
		alpha2:      "MX",
		countryCode: "484",
		sibCode:     143,
	},
	{
		name:        "Micronesia (Federated States of)",
		alpha2:      "FM",
		countryCode: "583",
		sibCode:     144,
	},
	{
		name:        "Moldova, Republic of",
		alpha2:      "MD",
		countryCode: "498",
		sibCode:     145,
	},
	{
		name:        "Monaco",
		alpha2:      "MC",
		countryCode: "492",
		sibCode:     146,
	},
	{
		name:        "Mongolia",
		alpha2:      "MN",
		countryCode: "496",
		sibCode:     147,
	},
	{
		name:        "Montenegro",
		alpha2:      "ME",
		countryCode: "499",
		sibCode:     148,
	},
	{
		name:        "Montserrat",
		alpha2:      "MS",
		countryCode: "500",
		sibCode:     149,
	},
	{
		name:        "Morocco",
		alpha2:      "MA",
		countryCode: "504",
		sibCode:     150,
	},
	{
		name:        "Mozambique",
		alpha2:      "MZ",
		countryCode: "508",
		sibCode:     151,
	},
	{
		name:        "Myanmar",
		alpha2:      "MM",
		countryCode: "104",
		sibCode:     152,
	},
	{
		name:        "Namibia",
		alpha2:      "NA",
		countryCode: "516",
		sibCode:     153,
	},
	{
		name:        "Nauru",
		alpha2:      "NR",
		countryCode: "520",
		sibCode:     154,
	},
	{
		name:        "Nepal",
		alpha2:      "NP",
		countryCode: "524",
		sibCode:     155,
	},
	{
		name:        "Netherlands",
		alpha2:      "NL",
		countryCode: "528",
		sibCode:     156,
	},
	{
		name:        "New Caledonia",
		alpha2:      "NC",
		countryCode: "540",
		sibCode:     157,
	},
	{
		name:        "New Zealand",
		alpha2:      "NZ",
		countryCode: "554",
		sibCode:     158,
	},
	{
		name:        "Nicaragua",
		alpha2:      "NI",
		countryCode: "558",
		sibCode:     159,
	},
	{
		name:        "Niger",
		alpha2:      "NE",
		countryCode: "562",
		sibCode:     160,
	},
	{
		name:        "Nigeria",
		alpha2:      "NG",
		countryCode: "566",
		sibCode:     161,
	},
	{
		name:        "Niue",
		alpha2:      "NU",
		countryCode: "570",
		sibCode:     162,
	},
	{
		name:        "Norfolk Island",
		alpha2:      "NF",
		countryCode: "574",
		sibCode:     163,
	},
	{
		name:        "North Macedonia",
		alpha2:      "MK",
		countryCode: "807",
		sibCode:     180,
	},
	{
		name:        "Northern Mariana Islands",
		alpha2:      "MP",
		countryCode: "580",
		sibCode:     164,
	},
	{
		name:        "Norway",
		alpha2:      "NO",
		countryCode: "578",
		sibCode:     165,
	},
	{
		name:        "Oman",
		alpha2:      "OM",
		countryCode: "512",
		sibCode:     166,
	},
	{
		name:        "Pakistan",
		alpha2:      "PK",
		countryCode: "586",
		sibCode:     167,
	},
	{
		name:        "Palau",
		alpha2:      "PW",
		countryCode: "585",
		sibCode:     168,
	},
	{
		name:        "Palestine, State of",
		alpha2:      "PS",
		countryCode: "275",
		sibCode:     169,
	},
	{
		name:        "Panama",
		alpha2:      "PA",
		countryCode: "591",
		sibCode:     170,
	},
	{
		name:        "Papua New Guinea",
		alpha2:      "PG",
		countryCode: "598",
		sibCode:     171,
	},
	{
		name:        "Paraguay",
		alpha2:      "PY",
		countryCode: "600",
		sibCode:     172,
	},
	{
		name:        "Peru",
		alpha2:      "PE",
		countryCode: "604",
		sibCode:     173,
	},
	{
		name:        "Philippines",
		alpha2:      "PH",
		countryCode: "608",
		sibCode:     174,
	},
	{
		name:        "Pitcairn",
		alpha2:      "PN",
		countryCode: "612",
		sibCode:     175,
	},
	{
		name:        "Poland",
		alpha2:      "PL",
		countryCode: "616",
		sibCode:     176,
	},
	{
		name:        "Portugal",
		alpha2:      "PT",
		countryCode: "620",
		sibCode:     177,
	},
	{
		name:        "Puerto Rico",
		alpha2:      "PR",
		countryCode: "630",
		sibCode:     178,
	},
	{
		name:        "Qatar",
		alpha2:      "QA",
		countryCode: "634",
		sibCode:     179,
	},
	{
		name:        "Reunion",
		alpha2:      "RE",
		countryCode: "638",
		sibCode:     184,
	},
	{
		name:        "Romania",
		alpha2:      "RO",
		countryCode: "642",
		sibCode:     181,
	},
	{
		name:        "Russian Federation",
		alpha2:      "RU",
		countryCode: "643",
		sibCode:     182,
	},
	{
		name:        "Rwanda",
		alpha2:      "RW",
		countryCode: "646",
		sibCode:     183,
	},
	{
		name:        "Saint Barthlemy",
		alpha2:      "BL",
		countryCode: "652",
		sibCode:     185,
	},
	{
		name:        "Saint Helena, Ascension and Tristan da Cunha",
		alpha2:      "SH",
		countryCode: "654",
		sibCode:     251,
	},
	{
		name:        "Saint Helena, Ascension and Tristan da Cunha",
		alpha2:      "SH",
		countryCode: "654",
		sibCode:     186,
	},
	{
		name:        "Saint Kitts and Nevis",
		alpha2:      "KN",
		countryCode: "659",
		sibCode:     187,
	},
	{
		name:        "Saint Lucia",
		alpha2:      "LC",
		countryCode: "662",
		sibCode:     188,
	},
	{
		name:        "Saint Martin (French part)",
		alpha2:      "MF",
		countryCode: "663",
		sibCode:     189,
	},
	{
		name:        "Saint Pierre and Miquelon",
		alpha2:      "PM",
		countryCode: "666",
		sibCode:     190,
	},
	{
		name:        "Saint Vincent and the Grenadines",
		alpha2:      "VC",
		countryCode: "670",
		sibCode:     191,
	},
	{
		name:        "Samoa",
		alpha2:      "WS",
		countryCode: "882",
		sibCode:     192,
	},
	{
		name:        "San Marino",
		alpha2:      "SM",
		countryCode: "674",
		sibCode:     193,
	},
	{
		name:        "Sao Tome and Principe",
		alpha2:      "ST",
		countryCode: "678",
		sibCode:     194,
	},
	{
		name:        "Saudi Arabia",
		alpha2:      "SA",
		countryCode: "682",
		sibCode:     195,
	},
	{
		name:        "Senegal",
		alpha2:      "SN",
		countryCode: "686",
		sibCode:     196,
	},
	{
		name:        "Serbia",
		alpha2:      "RS",
		countryCode: "688",
		sibCode:     197,
	},
	{
		name:        "Seychelles",
		alpha2:      "SC",
		countryCode: "690",
		sibCode:     198,
	},
	{
		name:        "Sierra Leone",
		alpha2:      "SL",
		countryCode: "694",
		sibCode:     199,
	},
	{
		name:        "Singapore",
		alpha2:      "SG",
		countryCode: "702",
		sibCode:     200,
	},
	{
		name:        "Sint Maarten (Dutch part)",
		alpha2:      "SX",
		countryCode: "534",
		sibCode:     201,
	},
	{
		name:        "Slovakia",
		alpha2:      "SK",
		countryCode: "703",
		sibCode:     202,
	},
	{
		name:        "Slovenia",
		alpha2:      "SI",
		countryCode: "705",
		sibCode:     203,
	},
	{
		name:        "Solomon Islands",
		alpha2:      "SB",
		countryCode: "090",
		sibCode:     204,
	},
	{
		name:        "Somalia",
		alpha2:      "SO",
		countryCode: "706",
		sibCode:     205,
	},
	{
		name:        "South Africa",
		alpha2:      "ZA",
		countryCode: "710",
		sibCode:     206,
	},
	{
		name:        "South Georgia and the South Sandwich Islands",
		alpha2:      "GS",
		countryCode: "239",
		sibCode:     207,
	},
	{
		name:        "South Sudan",
		alpha2:      "SS",
		countryCode: "728",
		sibCode:     208,
	},
	{
		name:        "Spain",
		alpha2:      "ES",
		countryCode: "724",
		sibCode:     209,
	},
	{
		name:        "Sri Lanka",
		alpha2:      "LK",
		countryCode: "144",
		sibCode:     210,
	},
	{
		name:        "Sudan",
		alpha2:      "SD",
		countryCode: "729",
		sibCode:     211,
	},
	{
		name:        "Suriname",
		alpha2:      "SR",
		countryCode: "740",
		sibCode:     212,
	},
	{
		name:        "Svalbard and Jan Mayen",
		alpha2:      "SJ",
		countryCode: "744",
		sibCode:     213,
	},
	{
		name:        "Sweden",
		alpha2:      "SE",
		countryCode: "752",
		sibCode:     214,
	},
	{
		name:        "Switzerland",
		alpha2:      "CH",
		countryCode: "756",
		sibCode:     215,
	},
	{
		name:        "Syrian Arab Republic",
		alpha2:      "SY",
		countryCode: "760",
		sibCode:     216,
	},
	{
		name:        "Taiwan, Province of China",
		alpha2:      "TW",
		countryCode: "158",
		sibCode:     217,
	},
	{
		name:        "Tajikistan",
		alpha2:      "TJ",
		countryCode: "762",
		sibCode:     218,
	},
	{
		name:        "Tanzania, United Republic of",
		alpha2:      "TZ",
		countryCode: "834",
		sibCode:     219,
	},
	{
		name:        "Thailand",
		alpha2:      "TH",
		countryCode: "764",
		sibCode:     220,
	},
	{
		name:        "Timor-Leste",
		alpha2:      "TL",
		countryCode: "626",
		sibCode:     221,
	},
	{
		name:        "Togo",
		alpha2:      "TG",
		countryCode: "768",
		sibCode:     222,
	},
	{
		name:        "Tokelau",
		alpha2:      "TK",
		countryCode: "772",
		sibCode:     223,
	},
	{
		name:        "Tonga",
		alpha2:      "TO",
		countryCode: "776",
		sibCode:     224,
	},
	{
		name:        "Trinidad and Tobago",
		alpha2:      "TT",
		countryCode: "780",
		sibCode:     225,
	},
	{
		name:        "Tunisia",
		alpha2:      "TN",
		countryCode: "788",
		sibCode:     226,
	},
	{
		name:        "Turkey",
		alpha2:      "TR",
		countryCode: "792",
		sibCode:     227,
	},
	{
		name:        "Turkmenistan",
		alpha2:      "TM",
		countryCode: "795",
		sibCode:     228,
	},
	{
		name:        "Turks and Caicos Islands",
		alpha2:      "TC",
		countryCode: "796",
		sibCode:     229,
	},
	{
		name:        "Tuvalu",
		alpha2:      "TV",
		countryCode: "798",
		sibCode:     230,
	},
	{
		name:        "Uganda",
		alpha2:      "UG",
		countryCode: "800",
		sibCode:     231,
	},
	{
		name:        "Ukraine",
		alpha2:      "UA",
		countryCode: "804",
		sibCode:     232,
	},
	{
		name:        "United Arab Emirates",
		alpha2:      "AE",
		countryCode: "784",
		sibCode:     234,
	},
	{
		name:        "United Kingdom of Great Britain and Northern Ireland",
		alpha2:      "GB",
		countryCode: "826",
		sibCode:     235,
	},
	{
		name:        "United States of America",
		alpha2:      "US",
		countryCode: "840",
		sibCode:     237,
	},
	{
		name:        "United States Minor Outlying Islands",
		alpha2:      "UM",
		countryCode: "581",
		sibCode:     236,
	},
	{
		name:        "Uruguay",
		alpha2:      "UY",
		countryCode: "858",
		sibCode:     238,
	},
	{
		name:        "Uzbekistan",
		alpha2:      "UZ",
		countryCode: "860",
		sibCode:     239,
	},
	{
		name:        "Vanuatu",
		alpha2:      "VU",
		countryCode: "548",
		sibCode:     240,
	},
	{
		name:        "Venezuela (Bolivarian Republic of)",
		alpha2:      "VE",
		countryCode: "862",
		sibCode:     241,
	},
	{
		name:        "Viet Nam",
		alpha2:      "VN",
		countryCode: "704",
		sibCode:     242,
	},
	{
		name:        "Virgin Islands (British)",
		alpha2:      "VG",
		countryCode: "092",
		sibCode:     243,
	},
	{
		name:        "Virgin Islands (U.S.)",
		alpha2:      "VI",
		countryCode: "850",
		sibCode:     244,
	},
	{
		name:        "Wallis and Futuna",
		alpha2:      "WF",
		countryCode: "876",
		sibCode:     245,
	},
	{
		name:        "Western Sahara",
		alpha2:      "EH",
		countryCode: "732",
		sibCode:     246,
	},
	{
		name:        "Yemen",
		alpha2:      "YE",
		countryCode: "887",
		sibCode:     247,
	},
	{
		name:        "Zambia",
		alpha2:      "ZM",
		countryCode: "894",
		sibCode:     248,
	},
	{
		name:        "Zimbabwe",
		alpha2:      "ZW",
		countryCode: "716",
		sibCode:     249,
	},
	{
		name:        "Aland Islands",
		alpha2:      "AX",
		countryCode: "248",
		sibCode:     250,
	},
}
