package logger

import (
	"fmt"
	"time"
)

var (
	Black   = color("\033[1;30m%s\033[0m")
	Red     = color("\033[1;31m%s\033[0m")
	Green   = color("\033[1;32m%s\033[0m")
	Yellow  = color("\033[1;33m%s\033[0m")
	Purple  = color("\033[1;34m%s\033[0m")
	Magenta = color("\033[1;35m%s\033[0m")
	Teal    = color("\033[1;36m%s\033[0m")
	White   = color("\033[1;37m%s\033[0m")
)

// Изменение цвета в консоли
func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// логирование уровня info
func Success(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(Green("[success] "))
	fmt.Println(infoContent)
	loggerWriter(infoContent, "success")
}

// логирование уровня info
func Log(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(Teal("[info] "))
	fmt.Println(infoContent)
	loggerWriter(infoContent, "info")
}

// логирование уровня info
func Warn(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(Yellow("[warn] "))
	fmt.Println(infoContent)
	loggerWriter(infoContent, "warn")
}

// логирование уровня info
func Error(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(Red("[error] "))
	fmt.Println(infoContent)
	loggerWriter(infoContent, "error")
}

// инициализация логгера, создание необходимого файла
func loggerWriter(content string, level string) {
	// if err := utils.CheckFolder("./logs"); err != nil {
	// 	if err := utils.CreateFolder("./logs"); err != nil {
	// 		log.Fatal("error create logs folder")
	// 	}
	// }
	// fileName := level + "." + "log"
	// file, err := os.OpenFile("./logs/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("error opening file: ", err)
	// }
	// defer file.Close()
	// file.WriteString(content + "\n")
}
