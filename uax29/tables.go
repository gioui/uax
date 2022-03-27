package uax29

// Code generated by github.com/npillmayer/uax/internal/classgen DO NOT EDIT
//
// BSD License, Copyright (c) 2018, Norbert Pillmayer (norbert@pillmayer.com)

import (
	"strconv"
	"unicode"
)

// Class for uax29.
// Must be convertable to int.
type Class int

const (
	ALetterClass            Class = 0
	CRClass                 Class = 1
	Double_QuoteClass       Class = 2
	ExtendClass             Class = 3
	ExtendNumLetClass       Class = 4
	FormatClass             Class = 5
	Hebrew_LetterClass      Class = 6
	KatakanaClass           Class = 7
	LFClass                 Class = 8
	MidLetterClass          Class = 9
	MidNumClass             Class = 10
	MidNumLetClass          Class = 11
	NewlineClass            Class = 12
	NumericClass            Class = 13
	Regional_IndicatorClass Class = 14
	Single_QuoteClass       Class = 15
	WSegSpaceClass          Class = 16
	ZWJClass                Class = 17

	Other Class = -1 // pseudo class for any other
	sot   Class = -2 // pseudo class "start of text"
	eot   Class = -3 // pseudo class "end of text"
)

// Range tables for uax29 classes.
// Clients can check with unicode.Is(..., rune)
var (
	ALetter            = _ALetter
	CR                 = _CR
	Double_Quote       = _Double_Quote
	Extend             = _Extend
	ExtendNumLet       = _ExtendNumLet
	Format             = _Format
	Hebrew_Letter      = _Hebrew_Letter
	Katakana           = _Katakana
	LF                 = _LF
	MidLetter          = _MidLetter
	MidNum             = _MidNum
	MidNumLet          = _MidNumLet
	Newline            = _Newline
	Numeric            = _Numeric
	Regional_Indicator = _Regional_Indicator
	Single_Quote       = _Single_Quote
	WSegSpace          = _WSegSpace
	ZWJ                = _ZWJ
)

// String returns the Class name.
func (c Class) String() string {
	switch c {
	case Other:
		return "Other"
	case sot:
		return "sot"
	case eot:
		return "eot"
	default:
		return "Class(" + strconv.Itoa(int(c)) + ")"
	case ALetterClass:
		return "ALetterClass"
	case CRClass:
		return "CRClass"
	case Double_QuoteClass:
		return "Double_QuoteClass"
	case ExtendClass:
		return "ExtendClass"
	case ExtendNumLetClass:
		return "ExtendNumLetClass"
	case FormatClass:
		return "FormatClass"
	case Hebrew_LetterClass:
		return "Hebrew_LetterClass"
	case KatakanaClass:
		return "KatakanaClass"
	case LFClass:
		return "LFClass"
	case MidLetterClass:
		return "MidLetterClass"
	case MidNumClass:
		return "MidNumClass"
	case MidNumLetClass:
		return "MidNumLetClass"
	case NewlineClass:
		return "NewlineClass"
	case NumericClass:
		return "NumericClass"
	case Regional_IndicatorClass:
		return "Regional_IndicatorClass"
	case Single_QuoteClass:
		return "Single_QuoteClass"
	case WSegSpaceClass:
		return "WSegSpaceClass"
	case ZWJClass:
		return "ZWJClass"
	}
}

var rangeFromClass = []*unicode.RangeTable{
	ALetterClass:            ALetter,
	CRClass:                 CR,
	Double_QuoteClass:       Double_Quote,
	ExtendClass:             Extend,
	ExtendNumLetClass:       ExtendNumLet,
	FormatClass:             Format,
	Hebrew_LetterClass:      Hebrew_Letter,
	KatakanaClass:           Katakana,
	LFClass:                 LF,
	MidLetterClass:          MidLetter,
	MidNumClass:             MidNum,
	MidNumLetClass:          MidNumLet,
	NewlineClass:            Newline,
	NumericClass:            Numeric,
	Regional_IndicatorClass: Regional_Indicator,
	Single_QuoteClass:       Single_Quote,
	WSegSpaceClass:          WSegSpace,
	ZWJClass:                ZWJ,
}

// size 4172 bytes (4 KiB)
var _ALetter = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x41, 0x5a, 1},
		{0x61, 0x7a, 1},
		{0xaa, 0xb5, 11},
		{0xba, 0xc0, 6},
		{0xc1, 0xd6, 1},
		{0xd8, 0xf6, 1},
		{0xf8, 0x2d7, 1},
		{0x2de, 0x2e4, 1},
		{0x2ec, 0x2ff, 1},
		{0x370, 0x374, 1},
		{0x376, 0x377, 1},
		{0x37a, 0x37d, 1},
		{0x37f, 0x386, 7},
		{0x388, 0x38a, 1},
		{0x38c, 0x38e, 2},
		{0x38f, 0x3a1, 1},
		{0x3a3, 0x3f5, 1},
		{0x3f7, 0x481, 1},
		{0x48a, 0x52f, 1},
		{0x531, 0x556, 1},
		{0x559, 0x55b, 2},
		{0x55c, 0x560, 2},
		{0x561, 0x588, 1},
		{0x5f3, 0x620, 45},
		{0x621, 0x64a, 1},
		{0x66e, 0x66f, 1},
		{0x671, 0x6d3, 1},
		{0x6d5, 0x6e5, 16},
		{0x6e6, 0x6ee, 8},
		{0x6ef, 0x6fa, 11},
		{0x6fb, 0x6fc, 1},
		{0x6ff, 0x710, 17},
		{0x712, 0x72f, 1},
		{0x74d, 0x7a5, 1},
		{0x7b1, 0x7ca, 25},
		{0x7cb, 0x7ea, 1},
		{0x7f4, 0x7f5, 1},
		{0x7fa, 0x800, 6},
		{0x801, 0x815, 1},
		{0x81a, 0x824, 10},
		{0x828, 0x840, 24},
		{0x841, 0x858, 1},
		{0x860, 0x86a, 1},
		{0x8a0, 0x8b4, 1},
		{0x8b6, 0x8bd, 1},
		{0x904, 0x939, 1},
		{0x93d, 0x950, 19},
		{0x958, 0x961, 1},
		{0x971, 0x980, 1},
		{0x985, 0x98c, 1},
		{0x98f, 0x990, 1},
		{0x993, 0x9a8, 1},
		{0x9aa, 0x9b0, 1},
		{0x9b2, 0x9b6, 4},
		{0x9b7, 0x9b9, 1},
		{0x9bd, 0x9ce, 17},
		{0x9dc, 0x9dd, 1},
		{0x9df, 0x9e1, 1},
		{0x9f0, 0x9f1, 1},
		{0x9fc, 0xa05, 9},
		{0xa06, 0xa0a, 1},
		{0xa0f, 0xa10, 1},
		{0xa13, 0xa28, 1},
		{0xa2a, 0xa30, 1},
		{0xa32, 0xa33, 1},
		{0xa35, 0xa36, 1},
		{0xa38, 0xa39, 1},
		{0xa59, 0xa5c, 1},
		{0xa5e, 0xa72, 20},
		{0xa73, 0xa74, 1},
		{0xa85, 0xa8d, 1},
		{0xa8f, 0xa91, 1},
		{0xa93, 0xaa8, 1},
		{0xaaa, 0xab0, 1},
		{0xab2, 0xab3, 1},
		{0xab5, 0xab9, 1},
		{0xabd, 0xad0, 19},
		{0xae0, 0xae1, 1},
		{0xaf9, 0xb05, 12},
		{0xb06, 0xb0c, 1},
		{0xb0f, 0xb10, 1},
		{0xb13, 0xb28, 1},
		{0xb2a, 0xb30, 1},
		{0xb32, 0xb33, 1},
		{0xb35, 0xb39, 1},
		{0xb3d, 0xb5c, 31},
		{0xb5d, 0xb5f, 2},
		{0xb60, 0xb61, 1},
		{0xb71, 0xb83, 18},
		{0xb85, 0xb8a, 1},
		{0xb8e, 0xb90, 1},
		{0xb92, 0xb95, 1},
		{0xb99, 0xb9a, 1},
		{0xb9c, 0xb9e, 2},
		{0xb9f, 0xba3, 4},
		{0xba4, 0xba8, 4},
		{0xba9, 0xbaa, 1},
		{0xbae, 0xbb9, 1},
		{0xbd0, 0xc05, 53},
		{0xc06, 0xc0c, 1},
		{0xc0e, 0xc10, 1},
		{0xc12, 0xc28, 1},
		{0xc2a, 0xc39, 1},
		{0xc3d, 0xc58, 27},
		{0xc59, 0xc5a, 1},
		{0xc60, 0xc61, 1},
		{0xc80, 0xc85, 5},
		{0xc86, 0xc8c, 1},
		{0xc8e, 0xc90, 1},
		{0xc92, 0xca8, 1},
		{0xcaa, 0xcb3, 1},
		{0xcb5, 0xcb9, 1},
		{0xcbd, 0xcde, 33},
		{0xce0, 0xce1, 1},
		{0xcf1, 0xcf2, 1},
		{0xd05, 0xd0c, 1},
		{0xd0e, 0xd10, 1},
		{0xd12, 0xd3a, 1},
		{0xd3d, 0xd4e, 17},
		{0xd54, 0xd56, 1},
		{0xd5f, 0xd61, 1},
		{0xd7a, 0xd7f, 1},
		{0xd85, 0xd96, 1},
		{0xd9a, 0xdb1, 1},
		{0xdb3, 0xdbb, 1},
		{0xdbd, 0xdc0, 3},
		{0xdc1, 0xdc6, 1},
		{0xf00, 0xf40, 64},
		{0xf41, 0xf47, 1},
		{0xf49, 0xf6c, 1},
		{0xf88, 0xf8c, 1},
		{0x10a0, 0x10c5, 1},
		{0x10c7, 0x10cd, 6},
		{0x10d0, 0x10fa, 1},
		{0x10fc, 0x1248, 1},
		{0x124a, 0x124d, 1},
		{0x1250, 0x1256, 1},
		{0x1258, 0x125a, 2},
		{0x125b, 0x125d, 1},
		{0x1260, 0x1288, 1},
		{0x128a, 0x128d, 1},
		{0x1290, 0x12b0, 1},
		{0x12b2, 0x12b5, 1},
		{0x12b8, 0x12be, 1},
		{0x12c0, 0x12c2, 2},
		{0x12c3, 0x12c5, 1},
		{0x12c8, 0x12d6, 1},
		{0x12d8, 0x1310, 1},
		{0x1312, 0x1315, 1},
		{0x1318, 0x135a, 1},
		{0x1380, 0x138f, 1},
		{0x13a0, 0x13f5, 1},
		{0x13f8, 0x13fd, 1},
		{0x1401, 0x166c, 1},
		{0x166f, 0x167f, 1},
		{0x1681, 0x169a, 1},
		{0x16a0, 0x16ea, 1},
		{0x16ee, 0x16f8, 1},
		{0x1700, 0x170c, 1},
		{0x170e, 0x1711, 1},
		{0x1720, 0x1731, 1},
		{0x1740, 0x1751, 1},
		{0x1760, 0x176c, 1},
		{0x176e, 0x1770, 1},
		{0x1820, 0x1878, 1},
		{0x1880, 0x1884, 1},
		{0x1887, 0x18a8, 1},
		{0x18aa, 0x18b0, 6},
		{0x18b1, 0x18f5, 1},
		{0x1900, 0x191e, 1},
		{0x1a00, 0x1a16, 1},
		{0x1b05, 0x1b33, 1},
		{0x1b45, 0x1b4b, 1},
		{0x1b83, 0x1ba0, 1},
		{0x1bae, 0x1baf, 1},
		{0x1bba, 0x1be5, 1},
		{0x1c00, 0x1c23, 1},
		{0x1c4d, 0x1c4f, 1},
		{0x1c5a, 0x1c7d, 1},
		{0x1c80, 0x1c88, 1},
		{0x1c90, 0x1cba, 1},
		{0x1cbd, 0x1cbf, 1},
		{0x1ce9, 0x1cec, 1},
		{0x1cee, 0x1cf1, 1},
		{0x1cf5, 0x1cf6, 1},
		{0x1d00, 0x1dbf, 1},
		{0x1e00, 0x1f15, 1},
		{0x1f18, 0x1f1d, 1},
		{0x1f20, 0x1f45, 1},
		{0x1f48, 0x1f4d, 1},
		{0x1f50, 0x1f57, 1},
		{0x1f59, 0x1f5f, 2},
		{0x1f60, 0x1f7d, 1},
		{0x1f80, 0x1fb4, 1},
		{0x1fb6, 0x1fbc, 1},
		{0x1fbe, 0x1fc2, 4},
		{0x1fc3, 0x1fc4, 1},
		{0x1fc6, 0x1fcc, 1},
		{0x1fd0, 0x1fd3, 1},
		{0x1fd6, 0x1fdb, 1},
		{0x1fe0, 0x1fec, 1},
		{0x1ff2, 0x1ff4, 1},
		{0x1ff6, 0x1ffc, 1},
		{0x2071, 0x207f, 14},
		{0x2090, 0x209c, 1},
		{0x2102, 0x2107, 5},
		{0x210a, 0x2113, 1},
		{0x2115, 0x2119, 4},
		{0x211a, 0x211d, 1},
		{0x2124, 0x212a, 2},
		{0x212b, 0x212d, 1},
		{0x212f, 0x2139, 1},
		{0x213c, 0x213f, 1},
		{0x2145, 0x2149, 1},
		{0x214e, 0x2160, 18},
		{0x2161, 0x2188, 1},
		{0x24b6, 0x24e9, 1},
		{0x2c00, 0x2c2e, 1},
		{0x2c30, 0x2c5e, 1},
		{0x2c60, 0x2ce4, 1},
		{0x2ceb, 0x2cee, 1},
		{0x2cf2, 0x2cf3, 1},
		{0x2d00, 0x2d25, 1},
		{0x2d27, 0x2d2d, 6},
		{0x2d30, 0x2d67, 1},
		{0x2d6f, 0x2d80, 17},
		{0x2d81, 0x2d96, 1},
		{0x2da0, 0x2da6, 1},
		{0x2da8, 0x2dae, 1},
		{0x2db0, 0x2db6, 1},
		{0x2db8, 0x2dbe, 1},
		{0x2dc0, 0x2dc6, 1},
		{0x2dc8, 0x2dce, 1},
		{0x2dd0, 0x2dd6, 1},
		{0x2dd8, 0x2dde, 1},
		{0x2e2f, 0x3005, 470},
		{0x303b, 0x303c, 1},
		{0x3105, 0x312f, 1},
		{0x3131, 0x318e, 1},
		{0x31a0, 0x31ba, 1},
		{0xa000, 0xa48c, 1},
		{0xa4d0, 0xa4fd, 1},
		{0xa500, 0xa60c, 1},
		{0xa610, 0xa61f, 1},
		{0xa62a, 0xa62b, 1},
		{0xa640, 0xa66e, 1},
		{0xa67f, 0xa69d, 1},
		{0xa6a0, 0xa6ef, 1},
		{0xa717, 0xa7b9, 1},
		{0xa7f7, 0xa801, 1},
		{0xa803, 0xa805, 1},
		{0xa807, 0xa80a, 1},
		{0xa80c, 0xa822, 1},
		{0xa840, 0xa873, 1},
		{0xa882, 0xa8b3, 1},
		{0xa8f2, 0xa8f7, 1},
		{0xa8fb, 0xa8fd, 2},
		{0xa8fe, 0xa90a, 12},
		{0xa90b, 0xa925, 1},
		{0xa930, 0xa946, 1},
		{0xa960, 0xa97c, 1},
		{0xa984, 0xa9b2, 1},
		{0xa9cf, 0xaa00, 49},
		{0xaa01, 0xaa28, 1},
		{0xaa40, 0xaa42, 1},
		{0xaa44, 0xaa4b, 1},
		{0xaae0, 0xaaea, 1},
		{0xaaf2, 0xaaf4, 1},
		{0xab01, 0xab06, 1},
		{0xab09, 0xab0e, 1},
		{0xab11, 0xab16, 1},
		{0xab20, 0xab26, 1},
		{0xab28, 0xab2e, 1},
		{0xab30, 0xab65, 1},
		{0xab70, 0xabe2, 1},
		{0xac00, 0xd7a3, 1},
		{0xd7b0, 0xd7c6, 1},
		{0xd7cb, 0xd7fb, 1},
		{0xfb00, 0xfb06, 1},
		{0xfb13, 0xfb17, 1},
		{0xfb50, 0xfbb1, 1},
		{0xfbd3, 0xfd3d, 1},
		{0xfd50, 0xfd8f, 1},
		{0xfd92, 0xfdc7, 1},
		{0xfdf0, 0xfdfb, 1},
		{0xfe70, 0xfe74, 1},
		{0xfe76, 0xfefc, 1},
		{0xff21, 0xff3a, 1},
		{0xff41, 0xff5a, 1},
		{0xffa0, 0xffbe, 1},
		{0xffc2, 0xffc7, 1},
		{0xffca, 0xffcf, 1},
		{0xffd2, 0xffd7, 1},
		{0xffda, 0xffdc, 1},
	},
	R32: []unicode.Range32{
		{0x10000, 0x1000b, 1},
		{0x1000d, 0x10026, 1},
		{0x10028, 0x1003a, 1},
		{0x1003c, 0x1003d, 1},
		{0x1003f, 0x1004d, 1},
		{0x10050, 0x1005d, 1},
		{0x10080, 0x100fa, 1},
		{0x10140, 0x10174, 1},
		{0x10280, 0x1029c, 1},
		{0x102a0, 0x102d0, 1},
		{0x10300, 0x1031f, 1},
		{0x1032d, 0x1034a, 1},
		{0x10350, 0x10375, 1},
		{0x10380, 0x1039d, 1},
		{0x103a0, 0x103c3, 1},
		{0x103c8, 0x103cf, 1},
		{0x103d1, 0x103d5, 1},
		{0x10400, 0x1049d, 1},
		{0x104b0, 0x104d3, 1},
		{0x104d8, 0x104fb, 1},
		{0x10500, 0x10527, 1},
		{0x10530, 0x10563, 1},
		{0x10600, 0x10736, 1},
		{0x10740, 0x10755, 1},
		{0x10760, 0x10767, 1},
		{0x10800, 0x10805, 1},
		{0x10808, 0x1080a, 2},
		{0x1080b, 0x10835, 1},
		{0x10837, 0x10838, 1},
		{0x1083c, 0x1083f, 3},
		{0x10840, 0x10855, 1},
		{0x10860, 0x10876, 1},
		{0x10880, 0x1089e, 1},
		{0x108e0, 0x108f2, 1},
		{0x108f4, 0x108f5, 1},
		{0x10900, 0x10915, 1},
		{0x10920, 0x10939, 1},
		{0x10980, 0x109b7, 1},
		{0x109be, 0x109bf, 1},
		{0x10a00, 0x10a10, 16},
		{0x10a11, 0x10a13, 1},
		{0x10a15, 0x10a17, 1},
		{0x10a19, 0x10a35, 1},
		{0x10a60, 0x10a7c, 1},
		{0x10a80, 0x10a9c, 1},
		{0x10ac0, 0x10ac7, 1},
		{0x10ac9, 0x10ae4, 1},
		{0x10b00, 0x10b35, 1},
		{0x10b40, 0x10b55, 1},
		{0x10b60, 0x10b72, 1},
		{0x10b80, 0x10b91, 1},
		{0x10c00, 0x10c48, 1},
		{0x10c80, 0x10cb2, 1},
		{0x10cc0, 0x10cf2, 1},
		{0x10d00, 0x10d23, 1},
		{0x10f00, 0x10f1c, 1},
		{0x10f27, 0x10f30, 9},
		{0x10f31, 0x10f45, 1},
		{0x11003, 0x11037, 1},
		{0x11083, 0x110af, 1},
		{0x110d0, 0x110e8, 1},
		{0x11103, 0x11126, 1},
		{0x11144, 0x11150, 12},
		{0x11151, 0x11172, 1},
		{0x11176, 0x11183, 13},
		{0x11184, 0x111b2, 1},
		{0x111c1, 0x111c4, 1},
		{0x111da, 0x111dc, 2},
		{0x11200, 0x11211, 1},
		{0x11213, 0x1122b, 1},
		{0x11280, 0x11286, 1},
		{0x11288, 0x1128a, 2},
		{0x1128b, 0x1128d, 1},
		{0x1128f, 0x1129d, 1},
		{0x1129f, 0x112a8, 1},
		{0x112b0, 0x112de, 1},
		{0x11305, 0x1130c, 1},
		{0x1130f, 0x11310, 1},
		{0x11313, 0x11328, 1},
		{0x1132a, 0x11330, 1},
		{0x11332, 0x11333, 1},
		{0x11335, 0x11339, 1},
		{0x1133d, 0x11350, 19},
		{0x1135d, 0x11361, 1},
		{0x11400, 0x11434, 1},
		{0x11447, 0x1144a, 1},
		{0x11480, 0x114af, 1},
		{0x114c4, 0x114c5, 1},
		{0x114c7, 0x11580, 185},
		{0x11581, 0x115ae, 1},
		{0x115d8, 0x115db, 1},
		{0x11600, 0x1162f, 1},
		{0x11644, 0x11680, 60},
		{0x11681, 0x116aa, 1},
		{0x11800, 0x1182b, 1},
		{0x118a0, 0x118df, 1},
		{0x118ff, 0x11a00, 257},
		{0x11a0b, 0x11a32, 1},
		{0x11a3a, 0x11a50, 22},
		{0x11a5c, 0x11a83, 1},
		{0x11a86, 0x11a89, 1},
		{0x11a9d, 0x11ac0, 35},
		{0x11ac1, 0x11af8, 1},
		{0x11c00, 0x11c08, 1},
		{0x11c0a, 0x11c2e, 1},
		{0x11c40, 0x11c72, 50},
		{0x11c73, 0x11c8f, 1},
		{0x11d00, 0x11d06, 1},
		{0x11d08, 0x11d09, 1},
		{0x11d0b, 0x11d30, 1},
		{0x11d46, 0x11d60, 26},
		{0x11d61, 0x11d65, 1},
		{0x11d67, 0x11d68, 1},
		{0x11d6a, 0x11d89, 1},
		{0x11d98, 0x11ee0, 328},
		{0x11ee1, 0x11ef2, 1},
		{0x12000, 0x12399, 1},
		{0x12400, 0x1246e, 1},
		{0x12480, 0x12543, 1},
		{0x13000, 0x1342e, 1},
		{0x14400, 0x14646, 1},
		{0x16800, 0x16a38, 1},
		{0x16a40, 0x16a5e, 1},
		{0x16ad0, 0x16aed, 1},
		{0x16b00, 0x16b2f, 1},
		{0x16b40, 0x16b43, 1},
		{0x16b63, 0x16b77, 1},
		{0x16b7d, 0x16b8f, 1},
		{0x16e40, 0x16e7f, 1},
		{0x16f00, 0x16f44, 1},
		{0x16f50, 0x16f93, 67},
		{0x16f94, 0x16f9f, 1},
		{0x16fe0, 0x16fe1, 1},
		{0x1bc00, 0x1bc6a, 1},
		{0x1bc70, 0x1bc7c, 1},
		{0x1bc80, 0x1bc88, 1},
		{0x1bc90, 0x1bc99, 1},
		{0x1d400, 0x1d454, 1},
		{0x1d456, 0x1d49c, 1},
		{0x1d49e, 0x1d49f, 1},
		{0x1d4a2, 0x1d4a5, 3},
		{0x1d4a6, 0x1d4a9, 3},
		{0x1d4aa, 0x1d4ac, 1},
		{0x1d4ae, 0x1d4b9, 1},
		{0x1d4bb, 0x1d4bd, 2},
		{0x1d4be, 0x1d4c3, 1},
		{0x1d4c5, 0x1d505, 1},
		{0x1d507, 0x1d50a, 1},
		{0x1d50d, 0x1d514, 1},
		{0x1d516, 0x1d51c, 1},
		{0x1d51e, 0x1d539, 1},
		{0x1d53b, 0x1d53e, 1},
		{0x1d540, 0x1d544, 1},
		{0x1d546, 0x1d54a, 4},
		{0x1d54b, 0x1d550, 1},
		{0x1d552, 0x1d6a5, 1},
		{0x1d6a8, 0x1d6c0, 1},
		{0x1d6c2, 0x1d6da, 1},
		{0x1d6dc, 0x1d6fa, 1},
		{0x1d6fc, 0x1d714, 1},
		{0x1d716, 0x1d734, 1},
		{0x1d736, 0x1d74e, 1},
		{0x1d750, 0x1d76e, 1},
		{0x1d770, 0x1d788, 1},
		{0x1d78a, 0x1d7a8, 1},
		{0x1d7aa, 0x1d7c2, 1},
		{0x1d7c4, 0x1d7cb, 1},
		{0x1e800, 0x1e8c4, 1},
		{0x1e900, 0x1e943, 1},
		{0x1ee00, 0x1ee03, 1},
		{0x1ee05, 0x1ee1f, 1},
		{0x1ee21, 0x1ee22, 1},
		{0x1ee24, 0x1ee27, 3},
		{0x1ee29, 0x1ee32, 1},
		{0x1ee34, 0x1ee37, 1},
		{0x1ee39, 0x1ee3b, 2},
		{0x1ee42, 0x1ee47, 5},
		{0x1ee49, 0x1ee4d, 2},
		{0x1ee4e, 0x1ee4f, 1},
		{0x1ee51, 0x1ee52, 1},
		{0x1ee54, 0x1ee57, 3},
		{0x1ee59, 0x1ee61, 2},
		{0x1ee62, 0x1ee64, 2},
		{0x1ee67, 0x1ee6a, 1},
		{0x1ee6c, 0x1ee72, 1},
		{0x1ee74, 0x1ee77, 1},
		{0x1ee79, 0x1ee7c, 1},
		{0x1ee7e, 0x1ee80, 2},
		{0x1ee81, 0x1ee89, 1},
		{0x1ee8b, 0x1ee9b, 1},
		{0x1eea1, 0x1eea3, 1},
		{0x1eea5, 0x1eea9, 1},
		{0x1eeab, 0x1eebb, 1},
		{0x1f130, 0x1f149, 1},
		{0x1f150, 0x1f169, 1},
		{0x1f170, 0x1f189, 1},
	},
	LatinOffset: 6,
}

// size 62 bytes (0 KiB)
var _CR = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xd, 0xd, 1},
	},
	LatinOffset: 1,
}

// size 62 bytes (0 KiB)
var _Double_Quote = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x22, 0x22, 1},
	},
	LatinOffset: 1,
}

// size 2204 bytes (2 KiB)
var _Extend = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x300, 0x36f, 1},
		{0x483, 0x489, 1},
		{0x591, 0x5bd, 1},
		{0x5bf, 0x5c1, 2},
		{0x5c2, 0x5c4, 2},
		{0x5c5, 0x5c7, 2},
		{0x610, 0x61a, 1},
		{0x64b, 0x65f, 1},
		{0x670, 0x6d6, 102},
		{0x6d7, 0x6dc, 1},
		{0x6df, 0x6e4, 1},
		{0x6e7, 0x6e8, 1},
		{0x6ea, 0x6ed, 1},
		{0x711, 0x730, 31},
		{0x731, 0x74a, 1},
		{0x7a6, 0x7b0, 1},
		{0x7eb, 0x7f3, 1},
		{0x7fd, 0x816, 25},
		{0x817, 0x819, 1},
		{0x81b, 0x823, 1},
		{0x825, 0x827, 1},
		{0x829, 0x82d, 1},
		{0x859, 0x85b, 1},
		{0x8d3, 0x8e1, 1},
		{0x8e3, 0x903, 1},
		{0x93a, 0x93c, 1},
		{0x93e, 0x94f, 1},
		{0x951, 0x957, 1},
		{0x962, 0x963, 1},
		{0x981, 0x983, 1},
		{0x9bc, 0x9be, 2},
		{0x9bf, 0x9c4, 1},
		{0x9c7, 0x9c8, 1},
		{0x9cb, 0x9cd, 1},
		{0x9d7, 0x9e2, 11},
		{0x9e3, 0x9fe, 27},
		{0xa01, 0xa03, 1},
		{0xa3c, 0xa3e, 2},
		{0xa3f, 0xa42, 1},
		{0xa47, 0xa48, 1},
		{0xa4b, 0xa4d, 1},
		{0xa51, 0xa70, 31},
		{0xa71, 0xa75, 4},
		{0xa81, 0xa83, 1},
		{0xabc, 0xabe, 2},
		{0xabf, 0xac5, 1},
		{0xac7, 0xac9, 1},
		{0xacb, 0xacd, 1},
		{0xae2, 0xae3, 1},
		{0xafa, 0xaff, 1},
		{0xb01, 0xb03, 1},
		{0xb3c, 0xb3e, 2},
		{0xb3f, 0xb44, 1},
		{0xb47, 0xb48, 1},
		{0xb4b, 0xb4d, 1},
		{0xb56, 0xb57, 1},
		{0xb62, 0xb63, 1},
		{0xb82, 0xbbe, 60},
		{0xbbf, 0xbc2, 1},
		{0xbc6, 0xbc8, 1},
		{0xbca, 0xbcd, 1},
		{0xbd7, 0xc00, 41},
		{0xc01, 0xc04, 1},
		{0xc3e, 0xc44, 1},
		{0xc46, 0xc48, 1},
		{0xc4a, 0xc4d, 1},
		{0xc55, 0xc56, 1},
		{0xc62, 0xc63, 1},
		{0xc81, 0xc83, 1},
		{0xcbc, 0xcbe, 2},
		{0xcbf, 0xcc4, 1},
		{0xcc6, 0xcc8, 1},
		{0xcca, 0xccd, 1},
		{0xcd5, 0xcd6, 1},
		{0xce2, 0xce3, 1},
		{0xd00, 0xd03, 1},
		{0xd3b, 0xd3c, 1},
		{0xd3e, 0xd44, 1},
		{0xd46, 0xd48, 1},
		{0xd4a, 0xd4d, 1},
		{0xd57, 0xd62, 11},
		{0xd63, 0xd82, 31},
		{0xd83, 0xdca, 71},
		{0xdcf, 0xdd4, 1},
		{0xdd6, 0xdd8, 2},
		{0xdd9, 0xddf, 1},
		{0xdf2, 0xdf3, 1},
		{0xe31, 0xe34, 3},
		{0xe35, 0xe3a, 1},
		{0xe47, 0xe4e, 1},
		{0xeb1, 0xeb4, 3},
		{0xeb5, 0xeb9, 1},
		{0xebb, 0xebc, 1},
		{0xec8, 0xecd, 1},
		{0xf18, 0xf19, 1},
		{0xf35, 0xf39, 2},
		{0xf3e, 0xf3f, 1},
		{0xf71, 0xf84, 1},
		{0xf86, 0xf87, 1},
		{0xf8d, 0xf97, 1},
		{0xf99, 0xfbc, 1},
		{0xfc6, 0x102b, 101},
		{0x102c, 0x103e, 1},
		{0x1056, 0x1059, 1},
		{0x105e, 0x1060, 1},
		{0x1062, 0x1064, 1},
		{0x1067, 0x106d, 1},
		{0x1071, 0x1074, 1},
		{0x1082, 0x108d, 1},
		{0x108f, 0x109a, 11},
		{0x109b, 0x109d, 1},
		{0x135d, 0x135f, 1},
		{0x1712, 0x1714, 1},
		{0x1732, 0x1734, 1},
		{0x1752, 0x1753, 1},
		{0x1772, 0x1773, 1},
		{0x17b4, 0x17d3, 1},
		{0x17dd, 0x180b, 46},
		{0x180c, 0x180d, 1},
		{0x1885, 0x1886, 1},
		{0x18a9, 0x1920, 119},
		{0x1921, 0x192b, 1},
		{0x1930, 0x193b, 1},
		{0x1a17, 0x1a1b, 1},
		{0x1a55, 0x1a5e, 1},
		{0x1a60, 0x1a7c, 1},
		{0x1a7f, 0x1ab0, 49},
		{0x1ab1, 0x1abe, 1},
		{0x1b00, 0x1b04, 1},
		{0x1b34, 0x1b44, 1},
		{0x1b6b, 0x1b73, 1},
		{0x1b80, 0x1b82, 1},
		{0x1ba1, 0x1bad, 1},
		{0x1be6, 0x1bf3, 1},
		{0x1c24, 0x1c37, 1},
		{0x1cd0, 0x1cd2, 1},
		{0x1cd4, 0x1ce8, 1},
		{0x1ced, 0x1cf2, 5},
		{0x1cf3, 0x1cf4, 1},
		{0x1cf7, 0x1cf9, 1},
		{0x1dc0, 0x1df9, 1},
		{0x1dfb, 0x1dff, 1},
		{0x200c, 0x20d0, 196},
		{0x20d1, 0x20f0, 1},
		{0x2cef, 0x2cf1, 1},
		{0x2d7f, 0x2de0, 97},
		{0x2de1, 0x2dff, 1},
		{0x302a, 0x302f, 1},
		{0x3099, 0x309a, 1},
		{0xa66f, 0xa672, 1},
		{0xa674, 0xa67d, 1},
		{0xa69e, 0xa69f, 1},
		{0xa6f0, 0xa6f1, 1},
		{0xa802, 0xa806, 4},
		{0xa80b, 0xa823, 24},
		{0xa824, 0xa827, 1},
		{0xa880, 0xa881, 1},
		{0xa8b4, 0xa8c5, 1},
		{0xa8e0, 0xa8f1, 1},
		{0xa8ff, 0xa926, 39},
		{0xa927, 0xa92d, 1},
		{0xa947, 0xa953, 1},
		{0xa980, 0xa983, 1},
		{0xa9b3, 0xa9c0, 1},
		{0xa9e5, 0xaa29, 68},
		{0xaa2a, 0xaa36, 1},
		{0xaa43, 0xaa4c, 9},
		{0xaa4d, 0xaa7b, 46},
		{0xaa7c, 0xaa7d, 1},
		{0xaab0, 0xaab2, 2},
		{0xaab3, 0xaab4, 1},
		{0xaab7, 0xaab8, 1},
		{0xaabe, 0xaabf, 1},
		{0xaac1, 0xaaeb, 42},
		{0xaaec, 0xaaef, 1},
		{0xaaf5, 0xaaf6, 1},
		{0xabe3, 0xabea, 1},
		{0xabec, 0xabed, 1},
		{0xfb1e, 0xfe00, 738},
		{0xfe01, 0xfe0f, 1},
		{0xfe20, 0xfe2f, 1},
		{0xff9e, 0xff9f, 1},
	},
	R32: []unicode.Range32{
		{0x101fd, 0x102e0, 227},
		{0x10376, 0x1037a, 1},
		{0x10a01, 0x10a03, 1},
		{0x10a05, 0x10a06, 1},
		{0x10a0c, 0x10a0f, 1},
		{0x10a38, 0x10a3a, 1},
		{0x10a3f, 0x10ae5, 166},
		{0x10ae6, 0x10d24, 574},
		{0x10d25, 0x10d27, 1},
		{0x10f46, 0x10f50, 1},
		{0x11000, 0x11002, 1},
		{0x11038, 0x11046, 1},
		{0x1107f, 0x11082, 1},
		{0x110b0, 0x110ba, 1},
		{0x11100, 0x11102, 1},
		{0x11127, 0x11134, 1},
		{0x11145, 0x11146, 1},
		{0x11173, 0x11180, 13},
		{0x11181, 0x11182, 1},
		{0x111b3, 0x111c0, 1},
		{0x111c9, 0x111cc, 1},
		{0x1122c, 0x11237, 1},
		{0x1123e, 0x112df, 161},
		{0x112e0, 0x112ea, 1},
		{0x11300, 0x11303, 1},
		{0x1133b, 0x1133c, 1},
		{0x1133e, 0x11344, 1},
		{0x11347, 0x11348, 1},
		{0x1134b, 0x1134d, 1},
		{0x11357, 0x11362, 11},
		{0x11363, 0x11366, 3},
		{0x11367, 0x1136c, 1},
		{0x11370, 0x11374, 1},
		{0x11435, 0x11446, 1},
		{0x1145e, 0x114b0, 82},
		{0x114b1, 0x114c3, 1},
		{0x115af, 0x115b5, 1},
		{0x115b8, 0x115c0, 1},
		{0x115dc, 0x115dd, 1},
		{0x11630, 0x11640, 1},
		{0x116ab, 0x116b7, 1},
		{0x1171d, 0x1172b, 1},
		{0x1182c, 0x1183a, 1},
		{0x11a01, 0x11a0a, 1},
		{0x11a33, 0x11a39, 1},
		{0x11a3b, 0x11a3e, 1},
		{0x11a47, 0x11a51, 10},
		{0x11a52, 0x11a5b, 1},
		{0x11a8a, 0x11a99, 1},
		{0x11c2f, 0x11c36, 1},
		{0x11c38, 0x11c3f, 1},
		{0x11c92, 0x11ca7, 1},
		{0x11ca9, 0x11cb6, 1},
		{0x11d31, 0x11d36, 1},
		{0x11d3a, 0x11d3c, 2},
		{0x11d3d, 0x11d3f, 2},
		{0x11d40, 0x11d45, 1},
		{0x11d47, 0x11d8a, 67},
		{0x11d8b, 0x11d8e, 1},
		{0x11d90, 0x11d91, 1},
		{0x11d93, 0x11d97, 1},
		{0x11ef3, 0x11ef6, 1},
		{0x16af0, 0x16af4, 1},
		{0x16b30, 0x16b36, 1},
		{0x16f51, 0x16f7e, 1},
		{0x16f8f, 0x16f92, 1},
		{0x1bc9d, 0x1bc9e, 1},
		{0x1d165, 0x1d169, 1},
		{0x1d16d, 0x1d172, 1},
		{0x1d17b, 0x1d182, 1},
		{0x1d185, 0x1d18b, 1},
		{0x1d1aa, 0x1d1ad, 1},
		{0x1d242, 0x1d244, 1},
		{0x1da00, 0x1da36, 1},
		{0x1da3b, 0x1da6c, 1},
		{0x1da75, 0x1da84, 15},
		{0x1da9b, 0x1da9f, 1},
		{0x1daa1, 0x1daaf, 1},
		{0x1e000, 0x1e006, 1},
		{0x1e008, 0x1e018, 1},
		{0x1e01b, 0x1e021, 1},
		{0x1e023, 0x1e024, 1},
		{0x1e026, 0x1e02a, 1},
		{0x1e8d0, 0x1e8d6, 1},
		{0x1e944, 0x1e94a, 1},
		{0x1f3fb, 0x1f3ff, 1},
		{0xe0020, 0xe007f, 1},
		{0xe0100, 0xe01ef, 1},
	},
}

// size 92 bytes (0 KiB)
var _ExtendNumLet = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x5f, 0x202f, 8144},
		{0x203f, 0x2040, 1},
		{0x2054, 0xfe33, 56799},
		{0xfe34, 0xfe4d, 25},
		{0xfe4e, 0xfe4f, 1},
		{0xff3f, 0xff3f, 1},
	},
}

// size 170 bytes (0 KiB)
var _Format = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xad, 0x600, 1363},
		{0x601, 0x605, 1},
		{0x61c, 0x6dd, 193},
		{0x70f, 0x8e2, 467},
		{0x180e, 0x200e, 2048},
		{0x200f, 0x202a, 27},
		{0x202b, 0x202e, 1},
		{0x2060, 0x2064, 1},
		{0x2066, 0x206f, 1},
		{0xfeff, 0xfff9, 250},
		{0xfffa, 0xfffb, 1},
	},
	R32: []unicode.Range32{
		{0x110bd, 0x110cd, 16},
		{0x1bca0, 0x1bca3, 1},
		{0x1d173, 0x1d17a, 1},
		{0xe0001, 0xe0001, 1},
	},
}

// size 116 bytes (0 KiB)
var _Hebrew_Letter = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x5d0, 0x5ea, 1},
		{0x5ef, 0x5f2, 1},
		{0xfb1d, 0xfb1f, 2},
		{0xfb20, 0xfb28, 1},
		{0xfb2a, 0xfb36, 1},
		{0xfb38, 0xfb3c, 1},
		{0xfb3e, 0xfb40, 2},
		{0xfb41, 0xfb43, 2},
		{0xfb44, 0xfb46, 2},
		{0xfb47, 0xfb4f, 1},
	},
}

// size 116 bytes (0 KiB)
var _Katakana = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3031, 0x3035, 1},
		{0x309b, 0x309c, 1},
		{0x30a0, 0x30fa, 1},
		{0x30fc, 0x30ff, 1},
		{0x31f0, 0x31ff, 1},
		{0x32d0, 0x32fe, 1},
		{0x3300, 0x3357, 1},
		{0xff66, 0xff9d, 1},
	},
	R32: []unicode.Range32{
		{0x1b000, 0x1b000, 1},
	},
}

// size 62 bytes (0 KiB)
var _LF = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xa, 0xa, 1},
	},
	LatinOffset: 1,
}

// size 80 bytes (0 KiB)
var _MidLetter = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x3a, 0xb7, 125},
		{0x387, 0x5f4, 621},
		{0x2027, 0xfe13, 56812},
		{0xfe55, 0xff1a, 197},
	},
	LatinOffset: 1,
}

// size 104 bytes (0 KiB)
var _MidNum = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x2c, 0x3b, 15},
		{0x37e, 0x589, 523},
		{0x60c, 0x60d, 1},
		{0x66c, 0x7f8, 396},
		{0x2044, 0xfe10, 56780},
		{0xfe14, 0xfe50, 60},
		{0xfe54, 0xff0c, 184},
		{0xff1b, 0xff1b, 1},
	},
	LatinOffset: 1,
}

// size 80 bytes (0 KiB)
var _MidNumLet = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x2e, 0x2018, 8170},
		{0x2019, 0x2024, 11},
		{0xfe52, 0xff07, 181},
		{0xff0e, 0xff0e, 1},
	},
}

// size 74 bytes (0 KiB)
var _Newline = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xb, 0xc, 1},
		{0x85, 0x2028, 8099},
		{0x2029, 0x2029, 1},
	},
	LatinOffset: 1,
}

// size 518 bytes (0 KiB)
var _Numeric = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x30, 0x39, 1},
		{0x660, 0x669, 1},
		{0x66b, 0x6f0, 133},
		{0x6f1, 0x6f9, 1},
		{0x7c0, 0x7c9, 1},
		{0x966, 0x96f, 1},
		{0x9e6, 0x9ef, 1},
		{0xa66, 0xa6f, 1},
		{0xae6, 0xaef, 1},
		{0xb66, 0xb6f, 1},
		{0xbe6, 0xbef, 1},
		{0xc66, 0xc6f, 1},
		{0xce6, 0xcef, 1},
		{0xd66, 0xd6f, 1},
		{0xde6, 0xdef, 1},
		{0xe50, 0xe59, 1},
		{0xed0, 0xed9, 1},
		{0xf20, 0xf29, 1},
		{0x1040, 0x1049, 1},
		{0x1090, 0x1099, 1},
		{0x17e0, 0x17e9, 1},
		{0x1810, 0x1819, 1},
		{0x1946, 0x194f, 1},
		{0x19d0, 0x19d9, 1},
		{0x1a80, 0x1a89, 1},
		{0x1a90, 0x1a99, 1},
		{0x1b50, 0x1b59, 1},
		{0x1bb0, 0x1bb9, 1},
		{0x1c40, 0x1c49, 1},
		{0x1c50, 0x1c59, 1},
		{0xa620, 0xa629, 1},
		{0xa8d0, 0xa8d9, 1},
		{0xa900, 0xa909, 1},
		{0xa9d0, 0xa9d9, 1},
		{0xa9f0, 0xa9f9, 1},
		{0xaa50, 0xaa59, 1},
		{0xabf0, 0xabf9, 1},
	},
	R32: []unicode.Range32{
		{0x104a0, 0x104a9, 1},
		{0x10d30, 0x10d39, 1},
		{0x11066, 0x1106f, 1},
		{0x110f0, 0x110f9, 1},
		{0x11136, 0x1113f, 1},
		{0x111d0, 0x111d9, 1},
		{0x112f0, 0x112f9, 1},
		{0x11450, 0x11459, 1},
		{0x114d0, 0x114d9, 1},
		{0x11650, 0x11659, 1},
		{0x116c0, 0x116c9, 1},
		{0x11730, 0x11739, 1},
		{0x118e0, 0x118e9, 1},
		{0x11c50, 0x11c59, 1},
		{0x11d50, 0x11d59, 1},
		{0x11da0, 0x11da9, 1},
		{0x16a60, 0x16a69, 1},
		{0x16b50, 0x16b59, 1},
		{0x1d7ce, 0x1d7ff, 1},
		{0x1e950, 0x1e959, 1},
	},
	LatinOffset: 1,
}

// size 68 bytes (0 KiB)
var _Regional_Indicator = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x1f1e6, 0x1f1ff, 1},
	},
}

// size 62 bytes (0 KiB)
var _Single_Quote = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x27, 0x27, 1},
	},
	LatinOffset: 1,
}

// size 80 bytes (0 KiB)
var _WSegSpace = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x20, 0x1680, 5728},
		{0x2000, 0x2006, 1},
		{0x2008, 0x200a, 1},
		{0x205f, 0x3000, 4001},
	},
}

// size 62 bytes (0 KiB)
var _ZWJ = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x200d, 0x200d, 1},
	},
}