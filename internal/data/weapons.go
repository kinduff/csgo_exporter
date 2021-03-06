// Package data holds general data about CS:GO to be used as utils.
package data

// WeaponByID returns the formatted name of a weapon by ID.
func WeaponByID(key int) string {
	weaponsMap := map[int]string{
		1:   "Desert Eagle",
		2:   "Dual Berettas",
		3:   "Five-SeveN",
		4:   "Glock-18",
		7:   "AK-47",
		8:   "AUG",
		9:   "AWP",
		10:  "FAMAS",
		11:  "G3SG1",
		13:  "Galil AR",
		14:  "M249",
		16:  "M4A4",
		17:  "MAC-10",
		19:  "P90",
		20:  "Repulsor Device",
		23:  "MP5-SD",
		24:  "UMP-45",
		25:  "XM1014",
		26:  "PP-Bizon",
		27:  "MAG-7",
		28:  "Negev",
		29:  "Sawed-Off",
		30:  "Tec-9",
		31:  "Zeus x27",
		32:  "P2000",
		33:  "MP7",
		34:  "MP9",
		35:  "Nova",
		36:  "P250",
		37:  "Ballistic Shield",
		38:  "SCAR-20",
		39:  "SG 553",
		40:  "SSG 08",
		41:  "Knife",
		42:  "Knife",
		43:  "Flashbang",
		44:  "HE Grenade",
		45:  "Smoke Grenade",
		46:  "Molotov",
		47:  "Decoy Grenade",
		48:  "Incendiary Grenade",
		49:  "C4 Explosive",
		59:  "Knife",
		60:  "M4A1-S",
		61:  "USP-S",
		63:  "CZ75-Auto",
		64:  "R8 Revolver",
		68:  "Tactical Awareness Grenade",
		69:  "Bare Hands",
		70:  "Breach Charge",
		74:  "Knife",
		75:  "Axe",
		76:  "Hammer",
		78:  "Wrench",
		81:  "Fire Bomb",
		82:  "Diversion Device",
		83:  "Frag Grenade",
		84:  "Snowball",
		500: "Bayonet",
		503: "Classic Knife",
		505: "Flip Knife",
		506: "Gut Knife",
		507: "Karambit",
		508: "M9 Bayonet",
		509: "Huntsman Knife",
		512: "Falchion Knife",
		514: "Bowie Knife",
		515: "Butterfly Knife",
		516: "Shadow Daggers",
		517: "Paracord Knife",
		518: "Survival Knife",
		519: "Ursus Knife",
		520: "Navaja Knife",
		521: "Nomad Knife",
		522: "Stiletto Knife",
		523: "Talon Knife",
		525: "Skeleton Knife",
	}

	return weaponsMap[key]
}

// WeaponByAPIName returns the formatted name of a weapon by API name.
func WeaponByAPIName(key string) string {
	weaponsMap := map[string]string{
		"ak47":      "AK-47",
		"aug":       "AUG",
		"awp":       "AWP",
		"bizon":     "PP-Bizon",
		"deagle":    "Desert Eagle",
		"decoy":     "Decoy Grenade",
		"elite":     "Dual Berettas",
		"famas":     "FAMAS",
		"fiveseven": "Five-SeveN",
		"g3sg1":     "G3SG1",
		"galilar":   "Galil AR",
		"glock":     "Glock-18",
		"hegrenade": "HE Grenade",
		"hkp2000":   "P2000",
		"knife":     "Knife",
		"m4a1":      "M4A1-S",
		"m249":      "M249",
		"mac10":     "MAC-10",
		"mag7":      "MAG-7",
		"molotov":   "Molotov",
		"mp7":       "MP7",
		"mp9":       "MP9",
		"negev":     "Negev",
		"nova":      "Nova",
		"p90":       "P90",
		"p250":      "P250",
		"sawedoff":  "Sawed-Off",
		"scar20":    "SCAR-20",
		"sg556":     "SG 553",
		"ssg08":     "SSG 08",
		"taser":     "Zeus x27",
		"tec9":      "Tec-9",
		"ump45":     "UMP-45",
		"xm1014":    "XM1014",
	}

	return weaponsMap[key]
}
