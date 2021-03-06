package cartridge

// https://gbdev.io/pandocs/The_Cartridge_Header.html#0147---cartridge-type
type CartridgeType byte

func (c CartridgeType) String() string {
	if str, ok := cartridgeTypeString[c]; ok {
		return str
	}

	return "unknown"
}

const (
	ROM_ONLY                       CartridgeType = 0x00
	MBC1                           CartridgeType = 0x01
	MB1_RAM                        CartridgeType = 0x02
	MBC1_RAM_BATTERY               CartridgeType = 0x03
	MBC2                           CartridgeType = 0x05
	MBC2_BATTERY                   CartridgeType = 0x06
	ROM_RAM                        CartridgeType = 0x08
	ROM_RAM_BATTERY                CartridgeType = 0x09
	MMM01                          CartridgeType = 0x0B
	MMM01_RAM                      CartridgeType = 0x0C
	MMM01_RAM_BATTERY              CartridgeType = 0x0D
	MBC3_TIMER_BATTERY             CartridgeType = 0x0F
	MBC3_TIMER_RAM_BATTERY         CartridgeType = 0x10
	MBC3                           CartridgeType = 0x11
	MBC3_RAM                       CartridgeType = 0x12
	MBC3_RAM_BATTERY               CartridgeType = 0x13
	MBC5                           CartridgeType = 0x19
	MBC5_RAM                       CartridgeType = 0x1A
	MBC5_RAM_BATTERY               CartridgeType = 0x1B
	MBC5_RUMBLE                    CartridgeType = 0x1C
	MBC5_RUMBLE_RAM                CartridgeType = 0x1D
	MBC5_RUMBLE_RAM_BATTERY        CartridgeType = 0x1E
	MBC6                           CartridgeType = 0x20
	MBC7_SENSOR_RUMBLE_RAM_BATTERY CartridgeType = 0x22
	POCKET_CAMERA                  CartridgeType = 0xFC
	BANDAI_TAMA5                   CartridgeType = 0xFD
	HUC3                           CartridgeType = 0xFE
	HUC1_RAM_BATTERY               CartridgeType = 0xFF
)

var (
	// https://gbdev.io/pandocs/The_Cartridge_Header.html#0104-0133---nintendo-logo
	NintendoLogo = [...]byte{
		0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
		0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
		0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
	}

	// https://gbdev.io/pandocs/The_Cartridge_Header.html#0144-0145---new-licensee-code
	NewLicenseeToPublisher = map[string]string{
		"00": "None",
		"01": "Nintendo R&D1",
		"08": "Capcom",
		"13": "Electronic Arts",
		"18": "Hudson Soft",
		"19": "b-ai",
		"20": "kss",
		"22": "pow",
		"24": "PCM Complete",
		"25": "san-x",
		"28": "Kemco Japan",
		"29": "seta",
		"30": "Viacom",
		"31": "Nintendo",
		"32": "Bandai",
		"33": "Ocean/Acclaim",
		"34": "Konami",
		"35": "Hector",
		"37": "Taito",
		"38": "Hudson",
		"39": "Banpresto",
		"41": "Ubi Soft",
		"42": "Atlus",
		"44": "Malibu",
		"46": "angel",
		"47": "Bullet-Proof",
		"49": "irem",
		"50": "Absolute",
		"51": "Acclaim",
		"52": "Activision",
		"53": "American sammy",
		"54": "Konami",
		"55": "Hi tech entertainment",
		"56": "LJN",
		"57": "Matchbox",
		"58": "Mattel",
		"59": "Milton Bradley",
		"60": "Titus",
		"61": "Virgin",
		"64": "LucasArts",
		"67": "Ocean",
		"69": "Electronic Arts",
		"70": "Infogrames",
		"71": "Interplay",
		"72": "Broderbund",
		"73": "sculptured",
		"75": "sci",
		"78": "THQ",
		"79": "Accolade",
		"80": "misawa",
		"83": "lozc",
		"86": "Tokuma Shoten Intermedia",
		"87": "Tsukuda Original",
		"91": "Chunsoft",
		"92": "Video system",
		"93": "Ocean/Acclaim",
		"95": "Varie",
		"96": "Yonezawa/s'pal",
		"97": "Kaneko",
		"99": "Pack in soft",
		"A4": "Konami (Yu-Gi-Oh!)",
	}

	// https://raw.githubusercontent.com/gb-archive/salvage/master/txt-files/gbrom.txt
	OldLicenseeToPublisher = map[byte]string{
		0x00: "none",
		0x01: "nintendo",
		0x08: "capcom",
		0x09: "hot-b",
		0x0A: "jaleco",
		0x0B: "coconuts",
		0x0C: "elite systems",
		0x13: "electronic arts",
		0x18: "hudsonsoft",
		0x19: "itc entertainment",
		0x1A: "yanoman",
		0x1D: "clary",
		0x1F: "virgin",
		0x24: "pcm complete",
		0x25: "san-x",
		0x28: "kotobuki systems",
		0x29: "seta",
		0x30: "infogrames",
		0x31: "nintendo",
		0x32: "bandai",
		// 0x33 (reserved) indicates to use new licensee
		0x34: "konami",
		0x35: "hector",
		0x38: "capcom",
		0x39: "banpresto",
		0x3C: "*entertainment i",
		0x3E: "gremlin",
		0x41: "ubi soft",
		0x42: "atlus",
		0x44: "malibu",
		0x46: "angel",
		0x47: "spectrum holoby",
		0x49: "irem",
		0x4A: "virgin",
		0x4D: "malibu",
		0x4F: "u.s. gold",
		0x50: "absolute",
		0x51: "acclaim",
		0x52: "activision",
		0x53: "american sammy",
		0x54: "gametek",
		0x55: "park place",
		0x56: "ljn",
		0x57: "matchbox",
		0x59: "milton bradley",
		0x5A: "mindscape",
		0x5B: "romstar",
		0x5C: "naxat soft",
		0x5D: "tradewest",
		0x60: "titus",
		0x61: "virgin",
		0x67: "ocean",
		0x69: "electronic arts",
		0x6E: "elite systems",
		0x6F: "electro brain",
		0x70: "infogrames",
		0x71: "interplay",
		0x72: "broderbund",
		0x73: "sculptered soft",
		0x75: "the sales curve",
		0x78: "t*hq",
		0x79: "accolade",
		0x7A: "triffix entertainment",
		0x7C: "microprose",
		0x7F: "kemco",
		0x80: "misawa entertainment",
		0x83: "lozc",
		0x86: "*tokuma shoten i",
		0x8B: "bullet-proof software",
		0x8C: "vic tokai",
		0x8E: "ape",
		0x8F: "i'max",
		0x91: "chun soft",
		0x92: "video system",
		0x93: "tsuburava",
		0x95: "varie",
		0x96: "yonezawa/s'pal",
		0x97: "kaneko",
		0x99: "arc",
		0x9A: "nihon bussan",
		0x9B: "tecmo",
		0x9C: "imagineer",
		0x9D: "banpresto",
		0x9F: "nova",
		0xA1: "hori electric",
		0xA2: "bandai",
		0xA4: "konami",
		0xA6: "kawada",
		0xA7: "takara",
		0xA9: "technos japan",
		0xAA: "broderbund",
		0xAC: "toei animation",
		0xAD: "toho",
		0xAF: "namco",
		0xB0: "acclaim",
		0xB1: "ascii or nexoft",
		0xB2: "bandai",
		0xB4: "enix",
		0xB6: "hal",
		0xB7: "snk",
		0xB9: "pony canyon",
		0xBA: "*culture brain o",
		0xBB: "sunsoft",
		0xBD: "sony imagesoft",
		0xBF: "sammy",
		0xC0: "taito",
		0xC2: "kemco",
		0xC3: "squaresoft",
		0xC4: "*tokuma shoten i",
		0xC5: "data east",
		0xC6: "tonkin house",
		0xC8: "koei",
		0xC9: "ufl",
		0xCA: "ultra",
		0xCB: "vap",
		0xCC: "use",
		0xCD: "meldac",
		0xCE: "*pony canyon or",
		0xCF: "angel",
		0xD0: "taito",
		0xD1: "sofel",
		0xD2: "quest",
		0xD3: "sigma enterprises",
		0xD4: "ask kodansha",
		0xD6: "naxat soft",
		0xD7: "copya systems",
		0xD9: "banpresto",
		0xDA: "tomy",
		0xDB: "ljn",
		0xDD: "ncs",
		0xDE: "human",
		0xDF: "altron",
		0xE0: "jaleco",
		0xE1: "towachiki",
		0xE2: "uutaka",
		0xE3: "varie",
		0xE5: "epoch",
		0xE7: "athena",
		0xE8: "asmik",
		0xE9: "natsume",
		0xEA: "king records",
		0xEB: "atlus",
		0xEC: "epic/sony records",
		0xEE: "igs",
		0xF0: "a wave",
		0xF3: "extreme entertainment",
		0xFF: "ljn",
	}

	cartridgeTypeString = map[CartridgeType]string{
		ROM_ONLY:                       "ROM_ONLY",
		MBC1:                           "MBC1",
		MB1_RAM:                        "MB1_RAM",
		MBC1_RAM_BATTERY:               "MBC1_RAM_BATTERY",
		MBC2:                           "MBC2",
		MBC2_BATTERY:                   "MBC2_BATTERY",
		ROM_RAM:                        "ROM_RAM",
		ROM_RAM_BATTERY:                "ROM_RAM_BATTERY",
		MMM01:                          "MMM01",
		MMM01_RAM:                      "MMM01_RAM",
		MMM01_RAM_BATTERY:              "MMM01_RAM_BATTERY",
		MBC3_TIMER_BATTERY:             "MBC3_TIMER_BATTERY",
		MBC3_TIMER_RAM_BATTERY:         "MBC3_TIMER_RAM_BATTERY",
		MBC3:                           "MBC3",
		MBC3_RAM:                       "MBC3_RAM",
		MBC3_RAM_BATTERY:               "MBC3_RAM_BATTERY",
		MBC5:                           "MBC5",
		MBC5_RAM:                       "MBC5_RAM",
		MBC5_RAM_BATTERY:               "MBC5_RAM_BATTERY",
		MBC5_RUMBLE:                    "MBC5_RUMBLE",
		MBC5_RUMBLE_RAM:                "MBC5_RUMBLE_RAM",
		MBC5_RUMBLE_RAM_BATTERY:        "MBC5_RUMBLE_RAM_BATTERY",
		MBC6:                           "MBC6",
		MBC7_SENSOR_RUMBLE_RAM_BATTERY: "MBC7_SENSOR_RUMBLE_RAM_BATTERY",
		POCKET_CAMERA:                  "POCKET_CAMERA",
		BANDAI_TAMA5:                   "BANDAI_TAMA5",
		HUC3:                           "HUC3",
		HUC1_RAM_BATTERY:               "HUC1_RAM_BATTERY",
	}
)
