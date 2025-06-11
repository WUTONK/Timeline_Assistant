package time_io

import (
	"log"
	"os"
	"strings"
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

	sb := strings.Builder{}

	for _, line := range content {
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	_, err = f.WriteString(sb.String())
	if err != nil {
		log.Fatal(err)
	}
}
