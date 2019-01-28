package main

import (
	"./ac"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	list := []string{}

	fileName := "word.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Print("error")
	}
	buf := bufio.NewReader(file)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		list = append(list, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read finish!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}



	content := "江泽民xxxx习近平"
	aho := ac.CreateAhoCorasick()

	for _,str := range list{
		aho.Add(str)
	}
	aho.BuildAhoCorasick()
	results := aho.ScanAhoCorasick(content)
	fmt.Println("匹配成功的词: ")
	if len(results) == 0{
		fmt.Print("无敏感词")
	}else{
		fmt.Print("存在敏感词\n")
		for _, result := range results {
			//fmt.Println(result)
			fmt.Println(string([]rune(content)[result.Begin : result.End+1]))
		}
	}
}
