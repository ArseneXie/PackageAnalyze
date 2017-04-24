package main

import (
	"bufio"
	"fmt"
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
	"sort"
	"encoding/json"
)

func removeDuplicatesUnordered(elements []string,skipPatStr string) []string {
	encountered := map[string]bool{}

	patternS := regexp.MustCompile(skipPatStr)
	// Create a map of all unique elements.
	for v:= range elements {
		if patternS.MatchString(elements[v]){
			continue
		}
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func main() {
	filelocPtr := flag.String("srcfile", "C:/test.txt", "source file location")
	outlocPtr := flag.String("output", "C:/result.txt", "output file location")
	skipPattern := flag.String("skipexp", "^g_[a-z]*.", "skip pattern")

	flag.Parse()

	fmt.Println(*filelocPtr)
	fmt.Println(*outlocPtr)
	fmt.Println(*skipPattern)

	file, err := os.Open(*filelocPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fout,err := os.Create(*outlocPtr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fout.Close()

	scanner := bufio.NewScanner(file)
	var pattenA = regexp.MustCompile(`[a-zA-Z]+[a-zA-Z_0-9]*\.[a-zA-Z_0-9]+`)
	var pattenX = regexp.MustCompile(`[a-zA-Z]+[a-zA-Z_0-9]*\.[a-zA-Z_0-9]+`)
	var finalX []string
	var tmpStr string


	cmtSL := regexp.MustCompile(`^--`)
	cmtMLs := regexp.MustCompile(`^/\*`)
	cmtMLe := regexp.MustCompile(`\*/$`)
	cmtS,cmtM, cmtMpre := false, false, false
	for scanner.Scan() {
		tmpStr =  strings.ToLower(scanner.Text())

		if cmtSL.MatchString(strings.Trim(tmpStr, " ")) {
			cmtS = true
		}else{
			cmtS = false
		}

		if cmtMLs.MatchString(strings.Trim(tmpStr, " ")) {
			cmtM = true
		}

		if cmtMpre&&cmtM {
			cmtM = false
		}

		if cmtMLe.MatchString(strings.Trim(tmpStr, " ")) {
			cmtMpre = true
		}else{
			cmtMpre = false
		}

		if !cmtS&&!cmtM {
			if pattenX.MatchString(tmpStr) {
				finalX = append(finalX, pattenA.FindAllString(tmpStr, -1)...)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	finalY := removeDuplicatesUnordered(finalX,*skipPattern)
	sort.Strings(finalY)
	b, err := json.MarshalIndent(finalY, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(strings.Replace(string(b),"\"","",-1))
	fout.WriteString(strings.Replace(string(b),"\"","",-1))

}