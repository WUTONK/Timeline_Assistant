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

	for _, line := range WriteContext {
		fmt.Print(line)
		_, err := File.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
