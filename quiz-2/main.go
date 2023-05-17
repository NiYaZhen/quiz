package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Process struct {
	PID      int       `json:"pid"`
	PPID     int       `json:"ppid"`
	Cmd      string    `json:"cmd"`
	Children []Process `json:"children,omitempty"`
}

func main() {
	// 檢查命令行參數
	if len(os.Args) < 2 {
		fmt.Println("請提供 CSV 檔案的路徑 Ex: go run main.go data.csv")
		return
	}

	// 開啟 CSV 檔案
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("無法開啟 CSV 檔案:", err)
		return
	}
	defer file.Close()

	// 建立 CSV Reader
	reader := csv.NewReader(file)

	// 讀取 CSV 檔案中的資料
	processes := make(map[int]*Process)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("讀取 CSV 檔案時發生錯誤:", err)
			return
		}

		// 將 CSV 資料轉換成 Process 結構
		pid, _ := strconv.Atoi(record[0])
		ppid, _ := strconv.Atoi(record[1])
		cmd := record[2]
		process := &Process{
			PID:  pid,
			PPID: ppid,
			Cmd:  cmd,
		}

		// 將 Process 存入 map
		processes[pid] = process
	}

	// 建立 Process 階層關係
	for _, process := range processes {
		parent, exists := processes[process.PPID]
		if exists {
			parent.Children = append(parent.Children, *process)
		}
	}

	// 找出根進程
	var rootProcesses []*Process
	for _, process := range processes {
		if process.PPID == 0 {
			rootProcesses = append(rootProcesses, process)
		}
	}

	// 將結果轉換成 JSON 格式
	jsonData, err := json.MarshalIndent(rootProcesses, "", "  ")
	if err != nil {
		fmt.Println("轉換為 JSON 格式時發生錯誤:", err)
		return
	}

	// 印出 JSON 資料
	fmt.Println(string(jsonData))
}
