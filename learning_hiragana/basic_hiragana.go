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
	clearTerminal()
	dict := map[string]string{
		"morning":                "あさ/asa",
		"asakusa(placed name)":   "あさくさ/asakusa",
		"chair":                  "いす/isu",
		"house":                  "いえ/ie",
		"udon noodles":           "うどん/udon",
		"sea":                    "うみ/umi",
		"station":                "えき/eki",
		"pencil":                 "えんぴつ/enpitsu",
		"japanese tea":           "おちゃ/ocha",
		"origami":                "おりがみ/origami",
		"key":                    "かぎ/kagi",
		"umbrella":               "かさ/kasa",
		"kimono":                 "きもの/kimono",
		"ticket":                 "きっぷ/kippu",
		"fruit":                  "くだもの/kudamono",
		"gloves":                 "てぶくろ/tebukuro",
		"ereaser":                "けしごむ/keshigomu",
		"flower arrangement":     "いけばな/ikebana",
		"park":                   "こうえん/kooen",
		"english tea":            "こうちゃ/koocha",
		"fish":                   "さかな/sakana",
		"sapporo(placed name)":   "さっぽろ/sapporo",
		"bullet train":           "しんかんせん/shinkansen",
		"traffic light":          "しんごう/shingoo",
		"sushi":                  "すし/sushi",
		"sliding door":           "ふすま/fusuma",
		"electric fan":           "せんぷうき/senpuuki",
		"washing machine":        "せんたくき/sentakuki",
		"sky":                    "そら/sora",
		"fried noodles":          "やきそば/yakisoba",
		"egg":                    "たまご/tamago",
		"tatami mat":             "たたみ/tatami",
		"metro":                  "ちかてつ/chikatetsu",
		"strawberry":             "いちご/ichigo",
		"moon":                   "つき/tsuki",
		"shoes":                  "くつ/kutsu",
		"tempura":                "てんぷら/tenpura",
		"stamp":                  "きって/kitte",
		"clock":                  "とけい/tokee",
		"tokyo station":          "とうきょうえき/tookyoo eki",
		"fireworks":              "はなび/hanabi",
		"flower":                 "はな/hana",
		"garden":                 "にわ/niwa",
		"doll":                   "にんぎょう/ningyoo",
		"stuffed toy":            "ぬいぐるみ/nuigurumi",
		"dog":                    "いぬ/inu",
		"cat":                    "ねこ/neko",
		"onion":                  "たまねぎ/tamanegi",
		"seaweed":                "のり/nori",
		"alcove":                 "とこのま/tokonoma",
		"festival":               "まつり/matsuri",
		"window":                 "まど/mado",
		"water":                  "みず/mizu",
		"miso soup":              "みそしる/misoshiru",
		"rice ball":              "おむすび/omusubi",
		"flour":                  "こむぎこ/komugiko",
		"glasses":                "めがね/megane",
		"business card":          "めいし/meeshi",
		"child":                  "こども/kodomo",
		"peach":                  "もも/momo",
		"baseball":               "やきゅう/yakyuu",
		"mountain":               "やま/yama",
		"post office":            "ゆうびんきょく/yuubinkyoku",
		"ring":                   "ゆびわ/yubiwa",
		"yokohaman(placed name)": "よこはま/yokohama",
		"night":                  "よる/yoru",
		"dish/plate":             "さら/sara",
		"akihabara(placed name)": "あきはばら/akihabara",
		"apple":                  "りんご/ringo",
		"bird":                   "とり/tori",
		"spring":                 "はる/haru",
		"car":                    "くるま/kuruma",
		"refrigerator":           "れいぞうこ/reezooko",
		"sliding door closet":    "おしいれ/oshiire",
		"corridor":               "ろうか/rooka",
		"bathroom/bath":          "おふろ/ofuro",
		"japanese horseradish":   "わさび/wasabi",
		"japanese style room":    "わしつ/washitsu",
		"to wash one's hands":    "てをあらいます/te o araimasu",
		"to telephone/call":      "でんわをします/denwa o shimasu",
		"station staff":          "えきいん/ekiin",
		"newspaper":              "しんぶん/shinbun",
	}
	keysdict := slices.Collect(maps.Keys(dict))
	lenDict := len(keysdict)
	var randNum int
	for {
		randNum = rand.IntN(lenDict)
		fmt.Printf("%s = ", keysdict[randNum])
		time.Sleep(10 * time.Second)
		fmt.Printf("%s", dict[keysdict[randNum]])
		time.Sleep(3 * time.Second)
		clearTerminal()
	}
}
