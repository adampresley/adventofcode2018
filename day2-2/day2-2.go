package main

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
)

var input = `zihrtxagncfpbsnolxydujjmqv
zihrtxagwcfpbsoolnydukjyqv
aihrtxagwcfpbsnoleybmkjmqv
zihrtxagwcfpbsnolgyduajmrv
zihrtxgmwcfpbunoleydukjmqv
zihqtxagwcfpbsnolesdukomqv
zihgtxagwcfpbsnoleydqkjqqv
dihrtxagwcqpbsnoleydpkjmqv
qihrtvagwcfpbsnollydukjmqv
zihrtgagwcfpbknoleyrukjmqv
cinrtxagwcfpbsnoleydukjaqv
zihrtxagwcfubsneleyvukjmqv
zihrtxagwcfpbsvoleydukvmtv
zihrtpagwcffbsnolfydukjmqv
zihrtxagwcfpbsxoleydtkjyqv
zohrvxugwcfpbsnoleydukjmqv
zyhrtxagdcfpbsnodeydukjmqv
zihrtxaghffpbsnoleyduojmqv
oihrtbagwcfpbsnoleyduejmqv
zihrtnagwcvpjsnoleydukjmqv
iihrtxagwcfpbsnoliyaukjmqv
ziartxagwcfpbsnokeydukjmpv
eibrtxagwccpbsnoleydukjmqv
zihrtxagwczwbsaoleydukjmqv
ziiatuagwcfpbsnoleydukjmqv
zzhrtxagwckpbsnsleydukjmqv
cihrtxaqwcfpbsnoleydkkjmqv
zihrtxaywcfpbsnoleydukzdqv
zihrtxagwjfpbvnoleydukjmql
zihrtxagwcfpbsnoleuduksmql
zizrtxxgwcfpbsnoleydukzmqv
zihrteagwcfpbsnobeydukjmqe
zihrtxafwhfpbsgoleydukjmqv
zitrtxagwcfpbsnoleyduvymqv
zihrtxauwcfebsnoleygukjmqv
zihrtxagwcfpbsnoleydubjrqh
zihrtxauwmfpbsnoleydukjmqo
zihrtxagwcdpbsnoleydukxmov
zihrtmagwcfpbsnoleydukvmlv
ziwrtxhgwcfpbsnoleodukjmqv
zihytxagacfpbsnoceydukjmqv
zihrtxagwcfpbsnolebdugjnqv
zihrzxagwcfpbsnjleyduktmqv
zihrtxygwcfpbinoleysukjmqv
zihrtxagwcfpbmnoveydujjmqv
zidrtxagwcfpbsnolexaukjmqv
zshrtxagwcepbsnoxeydukjmqv
yibrtxagwzfpbsnoleydukjmqv
zehrtxagwclpbsnoleymukjmqv
zihruxagwcfpbsnoleyhukwmqv
zihrwxagwcfpbszolesdukjmqv
zihrtpagwcfpbwnoleyuukjmqv
ziortxagwcfpssnolewdukjmqv
zohrtxagwcfpbwnoleydukjmjv
zihrtxagwcfpbsnvleyduzcmqv
zihrvxaghcfpbswoleydukjmqv
zihrtxagwcfpssnolwydukzmqv
zjhrttagwcfpbsnolfydukjmqv
zihrtxagwjfpbsnoljydukpmqv
ziwrtxagwczpbsnoljydukjmqv
zinrtxagwcfpbvfoleydukjmqv
zihrgragwcfpbsnoleydutjmqv
zihrtxagwcfpbsnozeydukffqv
zihrtxagwcfpbsmoleydxkumqv
rihwtxagwcfpbsxoleydukjmqv
ziqrtxagwcfpbsnqlevdukjmqv
zihrtxagwchpbsnoleydufamqv
sihrtxagwcfpbsnoleldukjmqp
zihrtxagwcrpbsnoleydvojmqv
zihrtxacwcfpbsnoweyxukjmqv
zihrtxagwcfpbsnolajmukjmqv
zzfrtxagwcfpbsnoleydukjmvv
zixrtxagwcfpbqnoleydukjgqv
zihitxaqwcfpbsnoleadukjmqv
zilrtxagecfxbsnoleydukjmqv
zihrtxagwcfpbypoleycukjmqv
zidrtxagdtfpbsnoleydukjmqv
lehrtxagxcfpbsnoleydukjmqv
zihrlxagwcfpbsncneydukjmqv
zihroxagbcspbsnoleydukjmqv
zihrtxagwcfkzsnolemdukjmqv
zihrtxagwcfpbsqeleydukkmqv
zihrjxagwcfpesnolxydukjmqv
zifrtxagwcfpbsooleydukkmqv
zirwtxagwcfpbsnoleydukzmqv
zjhntxagwcfpbsnoleydunjmqv
ziorexagwcfpbsnoyeydukjmqv
zhhrtlagwcfybsnoleydukjmqv
zirrtxagwvfsbsnoleydukjmqv
bihrtxagwofpbsnoleadukjmqv
dihrtxagwcfpksnoleydukjlqv
zihrrxagecfpbsnoleydukjmyv
zijrtxagwmfpbsnoleyduljmqv
zihrtxagwcfpbsnolecdukjpqs
zchrtxagwcfpbsnolehdukjmwv
rmhrtxagwcfpbsnoleydkkjmqv
zohrotagwcfpbsnoleydukjmqv
zihwtxagsifpbsnwleydukjmqv
zihrtxagicfpbsnoleydukjxqn
zihrtxsgwcfpbsntleydumjmqv
zihrlxagzgfpbsnoleydukjmqv
aihjtxagwdfpbsnoleydukjmqv
zifrtxagwcfhbsnoleddukjmqv
zihrtyagwcfpbsooleydtkjmqv
zihrtxxgwcfpbsnolerhukjmqv
zihqtxalwcfppsnoleydukjmqv
zfkrvxagwcfpbsnoleydukjmqv
zihptxagwcfpbseoleydukjmdv
zihrtxagwcfpeonoleyiukjmqv
nidrtxagwcfpbsnoleyhukjmqv
zihrtxagwcfjbsnolsydukjmqg
zghryxagwcfgbsnoleydukjmqv
zihwtxagwcfpbsnoleydugjfqv
zihryxagwjfpbsnoleydujjmqv
zihrtxagwcfpbsnolekdukymql
zfhrtxaownfpbsnoleydukjmqv
zamrtxagwcfpbsnoleyduzjmqv
ibhrtxagwcfpbsnoleydukjmfv
zihrtxagwcfpssnoseydukjmuv
zihrtxagwcfpbsnoljydukjhqs
zihrtxagwqfmbsnoleidukjmqv
zfdrtxagwchpbsnoleydukjmqv
iihrtxagqcfpbsnoleydukjmqn
mihrtxagwcfpbsqoleydukjbqv
zihttxagwcfpbsnoleyduljmqk
zzhrtxagwcfpzseoleydukjmqv
zdhrtxagbcfpbsnoleyduyjmqv
zihxtxagwcfpbsnolwrdukjmqv
zghrtxagwcypbynoleydukjmqv
zihrtxaiwcfppsnoleydukgmqv
zitatxagwcfobsnoleydukjmqv
znhrtxagwcfpysnoleydukjqqv
zihrtxagwcfppsnoleoyukjmqv
ziorgxagwcfpbsnolekdukjmqv
zihrtxagwcfpbfnoleydwkjpqv
zihrtxnrwcfpbsnolnydukjmqv
rihrtxagwcfpbsnolepdjkjmqv
zihrtxagwcfzbsnoceydukjmkv
zihrtxagwcfpysnoaeidukjmqv
zihrmxagwcfpbsnoleydukjmuq
gihrtxagwcvpbsnoleydukcmqv
zihrtxagocfpbsnoleydukqmnv
zihrtxagwcfpesnoleyluklmqv
zghrtxagwcfzbsnoleydukjmgv
zihrtxugqqfpbsnoleydukjmqv
zirrtcagwcfpbsnoleydfkjmqv
zihitxagwcfpjsnoleydnkjmqv
zihrtxqgwcfpbsnsleydukjmqy
iihrtxagwyfpbsnoleydukjmqu
zihrsxagwcfpbsnsleydukzmqv
zihrtxawwcfpbsnoleydzkjmuv
dihrkxagwcfpbsfoleydukjmqv
zihrtxaqwcfpbvnoleydukjmqt
zihntxdgwcfpbsnogeydukjmqv
zihrtxagwcdpxsnolxydukjmqv
zihrtxagwcfpbsaoleydunjaqv
zihrtyagwcfpbsnoleyduqjmqt
zihrtxagwtfpbsnoleoyukjmqv
zihrjiagwcfpbsnobeydukjmqv
zihrtxqgwcfpbsnoleydykdmqv
zihrhxmgwcfpbsnmleydukjmqv
zihatxlgwcfpbsnoleydukpmqv
zihrtxcgwcspbsnoleypukjmqv
zihrtkagqcfpbsaoleydukjmqv
ziqrtxagwcfabsnoleydukrmqv
zihwtxagwifpbsnwleydukjmqv
zitrtnagwcfpbsnoleddukjmqv
wihrtxagwcfpbsioyeydukjmqv
zihrtxagwclpystoleydukjmqv
zihmtxagwcfpbsnolfydukjmlv
zihrtxagechpbsnoleydutjmqv
zihrtxagwcfebsnolnydukjmuv
zihrtxagncmpbsnoleydukjmqs
zihrvxagocfpbsnoleydukcmqv
zihrtxagwcjcbsnolejdukjmqv
wihrtxagwcfpbogoleydukjmqv
kivrtxagwcfpgsnoleydukjmqv
zihrtxagwafpbhnoleydukjcqv
zihrtwagtcfpbsnolxydukjmqv
vihrtxagwcfpbsneletdukjmqv
zihlnxagwcfpbsnoleydukjmqb
zihrtxagwcfpbsnoleydukjuuc
zihrtxagwcfpbwntleadukjmqv
fihrtxagwcfpbsnoleydvkjmqw
zihrtxaowcfpbunoleyduljmqv
zthrtxagwcfpbtnoleydukomqv
xihltxagwcfpbsnoleydukjrqv
ziyrnxagwcfpbsnoleydukjmhv
zihrtxazwcfpbsnileyduejmqv
zihrtxagwcfibsnoliydukjmsv
zihrtxggwcfpbsnoleydugjmqj
zrartxagwcffbsnoleydukjmqv
zidrtxaqwcfpbsnoleyduksmqv
zirrtxagwcypbsnoleydtkjmqv
rihrtxagwcrpbsnoheydukjmqv
zihrtxagwcfpbsnoleydpkjmzs
zihrtxagbcfpbsnodbydukjmqv
fihrtxaqwcfpbsnolaydukjmqv
vihrtxbgwcfpbsnolemdukjmqv
zihrtxapwcfubsnoleydukmmqv
zihrtxagwcfpbgnolfydunjmqv
zihrtxagwcypbsnokeyduvjmqv
zihntxagwcfpbsnoieydukbmqv
zihbtxagwkfpbsnolpydukjmqv
zihrtxagwcfibsnoleydikjmqb
jihrtxvgwcfpbsnoleydukjmqp
zihrtxagwcfpbjnqleydukjmlv
zibrtxagwcfpbzvoleydukjmqv
zihrtxagwafgbsnbleydukjmqv
zihjctagwcfpbsnoleydukjmqv
zahrtxagwcepbsnoleddukjmqv
zihetxagwcfpbsnoleydumjmsv
zihrtvagwcfpbbnoleydukdmqv
zbhrxxagwkfpbsnoleydukjmqv
jfhrtxagwcftbsnoleydukjmqv
yihrtxagwcfvbsnoleyduksmqv
ziartxaewcfpbsnoleyduhjmqv
zihrtxagwcfpbsnoozyduzjmqv
cihotxagwcfpysnoleydukjmqv
zihrtxagwcfpusnolwydxkjmqv
zihrtxagwcfpbsnoleedmgjmqv
zihrtxaghcfpmsnoleydukqmqv
ziortxagwcfpbsboleidukjmqv
zihrtxagwcfybsnoleyqxkjmqv
zihrtxamwcfpbsngleydukjmqx
zihrtxagwcfpbsnoleyduusmqu
zihftxagwcfpssnwleydukjmqv
zihrtxagwcfkbsnomeydukjmsv
zihrtxagwcvpbsnooeydwkjmqv
zihrtxagwcfpbsnoleycekumqv
jahrtxagwcfpbsnoleydukjmmv
zihrtxabwcfpbsnzheydukjmqv
zihrtxagwctpbsnoleydwkjmhv
zihrtpagwcfpbsnoleydzkjmqh
zihwtxagwcfpbsnollydukjrqv
zihrtxagwcfpusnoleydsvjmqv
zibrtxagwcfpasnoleydukjmbv
zchrtmagwcfpbsnoleydukjmwv
ziertxbgwyfpbsnoleydukjmqv
zitrtxagwcfpbhnoweydukjmqv
zisrtxkgwcfpbsnopeydukjmqv
zihrtxcgwdfpbynoleydukjmqv
iihrtxajwcvpbsnoleydukjmqv
zihuwxapwcfpbsnoleydukjmqv
zihrtxngwcfqbsnoleyiukjmqv
ziqrtxagjcfpbsnoleydukjmqi
zifrtxarwctpbsnoleydukjmqv
zihxgxagwcfpbpnoleydukjmqv
giprtxagwcdpbsnoleydukjmqv
zihrtxagwmfpbsnodeydukjbqv`

func main() {
	boxIDs := strings.Split(input, "\n")
	correctBoxes := make(chan string, 5)

	wait := &sync.WaitGroup{}

	for index1 := 0; index1 < len(boxIDs)-2; index1++ {
		for index2 := index1 + 1; index2 < len(boxIDs)-1; index2++ {
			wait.Add(1)

			go func(boxID1, boxID2 string) {
				diffCount := 0

				for pos := range boxID1 {
					if boxID1[pos] == boxID2[pos] {
						continue
					}

					diffCount++
				}

				if diffCount == 1 {
					correctBoxes <- boxID1
					correctBoxes <- boxID2
				}

				wait.Done()
			}(boxIDs[index1], boxIDs[index2])
		}
	}

	wait.Wait()
	close(correctBoxes)

	fmt.Printf("\nBoxes:\n")
	boxes := make([]string, len(correctBoxes))
	boxIndex := 0

	for boxID := range correctBoxes {
		fmt.Printf("%s\n", boxID)
		boxes[boxIndex] = boxID

		boxIndex++
	}

	/*
	 * Now that we know which boxes are correct, and the position
	 * of the differing letter, build a string of letters that match
	 */
	commonLetters := &strings.Builder{}

	for index := range boxes[0] {
		if boxes[0][index] == boxes[1][index] {
			commonLetters.WriteByte(boxes[0][index])
		}
	}

	fmt.Printf("Common letters: %s\n", commonLetters.String())
}

func mapOccurences(boxID string) map[string]int {
	seen := make(map[string]int)
	ok := false

	scanner := bufio.NewScanner(strings.NewReader(boxID))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		ch := scanner.Text()

		if _, ok = seen[ch]; !ok {
			seen[ch] = 1
		} else {
			seen[ch]++
		}
	}

	return seen
}

func hasTwoOccurences(m map[string]int) bool {
	for _, count := range m {
		if count == 2 {
			return true
		}
	}

	return false
}

func hasThreeOccurences(m map[string]int) bool {
	for _, count := range m {
		if count > 2 {
			return true
		}
	}

	return false
}
