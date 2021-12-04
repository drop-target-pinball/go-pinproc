// Package wpc provides P-ROC device number mappings for the Williams
// Pinball Controller.
package wpc

// Coils and flashers
const (
	C01 = 40
	C02 = 41
	C03 = 42
	C04 = 43
	C05 = 44
	C06 = 45
	C07 = 46
	C08 = 47
	C09 = 48
	C10 = 49
	C11 = 50
	C12 = 51
	C13 = 52
	C14 = 53
	C15 = 54
	C16 = 55
	C17 = 56
	C18 = 57
	C19 = 58
	C20 = 59
	C21 = 60
	C22 = 61
	C23 = 62
	C24 = 63
	C25 = 64
	C26 = 65
	C27 = 66
	C28 = 67
	C29 = 32
	C30 = 33
	C31 = 34
	C32 = 35
	C33 = 36
	C34 = 37
	C35 = 38
	C36 = 39
	C37 = 144
	C38 = 145
	C39 = 146
	C40 = 147
	C41 = 148
	C42 = 149
	C43 = 150
	C44 = 151
)

// General illumination
const (
	G01 = 72
	G02 = 73
	G03 = 74
	G04 = 75
	G05 = 76
)

// Lamps
const (
	L11 = 80
	L12 = 81
	L13 = 82
	L14 = 83
	L15 = 84
	L16 = 85
	L17 = 86
	L18 = 87
	L21 = 88
	L22 = 89
	L23 = 90
	L24 = 91
	L25 = 92
	L26 = 93
	L27 = 94
	L28 = 95
	L31 = 96
	L32 = 97
	L33 = 98
	L34 = 99
	L35 = 100
	L36 = 101
	L37 = 102
	L38 = 103
	L41 = 104
	L42 = 105
	L43 = 106
	L44 = 107
	L45 = 108
	L46 = 109
	L47 = 110
	L48 = 111
	L51 = 112
	L52 = 113
	L53 = 114
	L54 = 115
	L55 = 116
	L56 = 117
	L57 = 118
	L58 = 119
	L61 = 120
	L62 = 121
	L63 = 122
	L64 = 123
	L65 = 124
	L66 = 125
	L67 = 126
	L68 = 127
	L71 = 128
	L72 = 129
	L73 = 130
	L74 = 131
	L75 = 132
	L76 = 133
	L77 = 134
	L78 = 135
	L81 = 136
	L82 = 137
	L83 = 138
	L84 = 139
	L85 = 140
	L86 = 141
	L87 = 142
	L88 = 143
)

// Switch matrix
const (
	S11 = 32
	S12 = 33
	S13 = 34
	S14 = 35
	S15 = 36
	S16 = 37
	S17 = 38
	S18 = 39
	S21 = 48
	S22 = 49
	S23 = 50
	S24 = 51
	S25 = 52
	S26 = 53
	S27 = 54
	S28 = 55
	S31 = 64
	S32 = 65
	S33 = 66
	S34 = 67
	S35 = 68
	S36 = 69
	S37 = 70
	S38 = 71
	S41 = 80
	S42 = 81
	S43 = 82
	S44 = 83
	S45 = 84
	S46 = 85
	S47 = 86
	S48 = 87
	S51 = 96
	S52 = 97
	S53 = 98
	S54 = 99
	S55 = 100
	S56 = 101
	S57 = 102
	S58 = 103
	S61 = 112
	S62 = 113
	S63 = 114
	S64 = 115
	S65 = 116
	S66 = 117
	S67 = 118
	S68 = 119
	S71 = 128
	S72 = 129
	S73 = 130
	S74 = 131
	S75 = 132
	S76 = 133
	S77 = 134
	S78 = 135
	S81 = 144
	S82 = 145
	S83 = 146
	S84 = 147
	S85 = 148
	S86 = 149
	S87 = 150
	S88 = 151
)

// Dedicated switches
const (
	SD1 = 8
	SD2 = 9
	SD3 = 10
	SD4 = 11
	SD5 = 12
	SD6 = 13
	SD7 = 14
	SD8 = 15
)

// Flipper switches
const (
	SF1 = 0
	SF2 = 1
	SF3 = 2
	SF4 = 3
	SF5 = 4
	SF6 = 5
	SF7 = 6
	SF8 = 7
)

// Flipper coils
const (
	FLLH = 35
	FLLM = 36
	FLRH = 37
	FLRM = 38
	FULH = 39
	FULM = 40
	FURH = 41
	FURM = 42
)

var Devices = map[string]uint8{
	"C01":  30,
	"C02":  31,
	"C03":  32,
	"C04":  33,
	"C05":  34,
	"C06":  35,
	"C07":  36,
	"C08":  37,
	"C09":  38,
	"C10":  39,
	"C11":  40,
	"C12":  41,
	"C13":  42,
	"C14":  43,
	"C15":  44,
	"C16":  45,
	"C17":  46,
	"C18":  47,
	"C19":  48,
	"C20":  49,
	"C21":  50,
	"C22":  51,
	"C23":  52,
	"C24":  53,
	"C25":  54,
	"C26":  55,
	"C27":  56,
	"C28":  57,
	"C29":  32,
	"C30":  33,
	"C31":  34,
	"C32":  35,
	"C33":  36,
	"C34":  37,
	"C35":  38,
	"C36":  39,
	"C37":  144,
	"C38":  145,
	"C39":  146,
	"C40":  147,
	"C41":  148,
	"C42":  149,
	"C43":  150,
	"C44":  151,
	"G01":  72,
	"G02":  73,
	"G03":  74,
	"G04":  75,
	"G05":  76,
	"L11":  80,
	"L12":  81,
	"L13":  82,
	"L14":  83,
	"L15":  84,
	"L16":  85,
	"L17":  86,
	"L18":  87,
	"L21":  88,
	"L22":  89,
	"L23":  90,
	"L24":  91,
	"L25":  92,
	"L26":  93,
	"L27":  94,
	"L28":  95,
	"L31":  96,
	"L32":  97,
	"L33":  98,
	"L34":  99,
	"L35":  100,
	"L36":  101,
	"L37":  102,
	"L38":  103,
	"L41":  104,
	"L42":  105,
	"L43":  106,
	"L44":  107,
	"L45":  108,
	"L46":  109,
	"L47":  110,
	"L48":  111,
	"L51":  112,
	"L52":  113,
	"L53":  114,
	"L54":  115,
	"L55":  116,
	"L56":  117,
	"L57":  118,
	"L58":  119,
	"L61":  120,
	"L62":  121,
	"L63":  122,
	"L64":  123,
	"L65":  124,
	"L66":  125,
	"L67":  126,
	"L68":  127,
	"L71":  128,
	"L72":  129,
	"L73":  130,
	"L74":  131,
	"L75":  132,
	"L76":  133,
	"L77":  134,
	"L78":  135,
	"L81":  136,
	"L82":  137,
	"L83":  138,
	"L84":  139,
	"L85":  140,
	"L86":  141,
	"L87":  142,
	"L88":  143,
	"S11":  32,
	"S12":  33,
	"S13":  34,
	"S14":  35,
	"S15":  36,
	"S16":  37,
	"S17":  38,
	"S18":  39,
	"S21":  48,
	"S22":  49,
	"S23":  50,
	"S24":  51,
	"S25":  52,
	"S26":  53,
	"S27":  54,
	"S28":  55,
	"S31":  64,
	"S32":  65,
	"S33":  66,
	"S34":  67,
	"S35":  68,
	"S36":  69,
	"S37":  70,
	"S38":  71,
	"S41":  80,
	"S42":  81,
	"S43":  82,
	"S44":  83,
	"S45":  84,
	"S46":  85,
	"S47":  86,
	"S48":  87,
	"S51":  96,
	"S52":  97,
	"S53":  98,
	"S54":  99,
	"S55":  100,
	"S56":  101,
	"S57":  102,
	"S58":  103,
	"S61":  112,
	"S62":  113,
	"S63":  114,
	"S64":  115,
	"S65":  116,
	"S66":  117,
	"S67":  118,
	"S68":  119,
	"S71":  128,
	"S72":  129,
	"S73":  130,
	"S74":  131,
	"S75":  132,
	"S76":  133,
	"S77":  134,
	"S78":  135,
	"S81":  144,
	"S82":  145,
	"S83":  146,
	"S84":  147,
	"S85":  148,
	"S86":  149,
	"S87":  150,
	"S88":  151,
	"SD1":  8,
	"SD2":  9,
	"SD3":  10,
	"SD4":  11,
	"SD5":  12,
	"SD6":  13,
	"SD7":  14,
	"SD8":  15,
	"SF1":  0,
	"SF2":  1,
	"SF3":  2,
	"SF4":  3,
	"SF5":  4,
	"SF6":  5,
	"SF7":  6,
	"SF8":  7,
	"FLLH": 35,
	"FLLM": 36,
	"FLRH": 37,
	"FLRM": 38,
	"FULH": 39,
	"FULM": 40,
	"FURH": 41,
	"FURM": 42,
}

var SwitchNames = map[uint8]string{
	32:  "S11",
	33:  "S12",
	34:  "S13",
	35:  "S14",
	36:  "S15",
	37:  "S16",
	38:  "S17",
	39:  "S18",
	48:  "S21",
	49:  "S22",
	50:  "S23",
	51:  "S24",
	52:  "S25",
	53:  "S26",
	54:  "S27",
	55:  "S28",
	64:  "S31",
	65:  "S32",
	66:  "S33",
	67:  "S34",
	68:  "S35",
	69:  "S36",
	70:  "S37",
	71:  "S38",
	80:  "S41",
	81:  "S42",
	82:  "S43",
	83:  "S44",
	84:  "S45",
	85:  "S46",
	86:  "S47",
	87:  "S48",
	96:  "S51",
	97:  "S52",
	98:  "S53",
	99:  "S54",
	100: "S55",
	101: "S56",
	102: "S57",
	103: "S58",
	112: "S61",
	113: "S62",
	114: "S63",
	115: "S64",
	116: "S65",
	117: "S66",
	118: "S67",
	119: "S68",
	128: "S71",
	129: "S72",
	130: "S73",
	131: "S74",
	132: "S75",
	133: "S76",
	134: "S77",
	135: "S78",
	144: "S81",
	145: "S82",
	146: "S83",
	147: "S84",
	148: "S85",
	149: "S86",
	150: "S87",
	151: "S88",
	8:   "SD1",
	9:   "SD2",
	10:  "SD3",
	11:  "SD4",
	12:  "SD5",
	13:  "SD6",
	14:  "SD7",
	15:  "SD8",
	0:   "SF1",
	1:   "SF2",
	2:   "SF3",
	3:   "SF4",
	4:   "SF5",
	5:   "SF6",
	6:   "SF7",
	7:   "SF8",
}
