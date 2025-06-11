package time_io

import (
	"fmt"
	"log"
	"os"
)

// 初始化文件
func WriteFile(path string, content []string) {
	// fmt.Println("-------------")
	// fmt.Print(WriteContext)
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, line := range content {
		fmt.Printf("写入行内容: '%s', 长度: %d\n", line, len(line))
		_, err := f.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}
