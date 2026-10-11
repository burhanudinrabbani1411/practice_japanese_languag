package learninghiragana

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"os"
	"os/exec"
	"slices"
	"time"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func HiraganaExcercises() {
	dict := map[string]string{
		"morning":                "asa",
		"asakusa(placed name)":   "asakusa",
		"chair":                  "isu",
		"house":                  "ie",
		"udon noodles":           "udon",
		"sea":                    "umi",
		"station":                "eki",
		"pencil":                 "enpitsu",
		"ocha":                   "tea",
		"origami":                "origami",
		"key":                    "kagi",
		"umbrella":               "kasa",
		"kimono":                 "kimono",
		"ticket":                 "kippu",
		"fruit":                  "kudamono",
		"gloves":                 "tebukuro",
		"ereaser":                "keshigomu",
		"flower arrangement":     "ikebana",
		"park":                   "kooen",
		"english tea":            "koocha",
		"fish":                   "sakana",
		"sapporo(placed name)":   "sapporo",
		"bullet train":           "shinkansen",
		"traffic light":          "shingoo",
		"sushi":                  "sushi",
		"sliding door":           "fusuma",
		"electric fan":           "senpuuki",
		"washing machine":        "sentakuki",
		"sky":                    "sora",
		"fried noodles":          "yakisoba",
		"egg":                    "tamago",
		"tatami mat":             "tatami",
		"metro":                  "chikatetsu",
		"strawberry":             "ichigo",
		"moon":                   "tsuki",
		"shoes":                  "kutsu",
		"tempura":                "tenpura",
		"stamp":                  "kitte",
		"clock":                  "tokee",
		"tokyo station":          "tookyoo eki",
		"fireworks":              "hanabi",
		"flower":                 "hana",
		"garden":                 "niwa",
		"doll":                   "ningyoo",
		"stuffed toy":            "nuigurumi",
		"dog":                    "inu",
		"cat":                    "neko",
		"onion":                  "tamanegi",
		"seaweed":                "nori",
		"alcove":                 "tokonoma",
		"festival":               "matsuri",
		"window":                 "mado",
		"water":                  "mizu",
		"miso soup":              "misoshiru",
		"rice ball":              "omusubi",
		"flour":                  "komugiko",
		"glasses":                "megane",
		"business card":          "meeshi",
		"child":                  "kodomo",
		"peach":                  "momo",
		"baseball":               "yakyuu",
		"mountain":               "yama",
		"post office":            "yuubinkyoku",
		"ring":                   "yubiwa",
		"yokohaman(placed name)": "yokohama",
		"night":                  "yoru",
		"dish/plate":             "sara",
		"akihabara(placed name)": "akihabara",
		"apple":                  "ringo",
		"bird":                   "tori",
		"spring":                 "haru",
		"car":                    "kuruma",
		"refrigerator":           "reezooko",
		"sliding door closet":    "oshiire",
		"corridor":               "rooka",
		"bathroom/bath":          "ofuro",
		"japanese horseradish":   "wasabi",
		"japanese style room":    "washitsu",
		"to wash one's hands":    "te o araimasu",
		"to telephone/call":      "denwa o shimasu",
		"station staff":          "ekiin",
		"newspaper":              "shinbun",
	}
	keysdict := slices.Collect(maps.Keys(dict))
	lenDict := len(keysdict)
	var randNum int
	for {
		randNum = rand.IntN(lenDict)
		fmt.Printf("%s = ", keysdict[randNum])
		time.Sleep(5 * time.Second)
		fmt.Printf("%s \n", dict[keysdict[randNum]])
		time.Sleep(3 * time.Second)
		clearTerminal()
	}
}
