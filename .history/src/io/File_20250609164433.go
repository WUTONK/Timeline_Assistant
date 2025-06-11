package time_io

import (
	"fmt"
	"log"
	"os"
)

// 初始化文件
func WriteFine(FilePath string, WriteContext []string) {
	// fmt.Println("-------------")
	// fmt.Print(WriteContext)
	File, err := os.Create(FilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer File.Close()

	// 过滤掉空字符串
	var filteredContext []string
	for _, line := range WriteContext {
		if line != "" {
			filteredContext = append(filteredContext, line)
		}
	}

	for _, line := range filteredContext {
		fmt.Printf("写入行内容: '%s', 长度: %d\n", line, len(line))
		_, err := File.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
