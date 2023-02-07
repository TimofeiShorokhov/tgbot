package dao

import (
	"bufio"
	"fmt"
	"os"
	"tgbot/data"
)

func Log(result *data.Client) error {
	content := "\nID: " +string(result.ID) +",Name: " + result.Name + ", Number: " + result.Number
	file, err := os.OpenFile("file.txt", os.O_APPEND | os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	file.Write([]byte(content))
	return nil
}

func DeleteLog(id string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line != id {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}