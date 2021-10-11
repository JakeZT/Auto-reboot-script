package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func main() {
	gocron.Every(1).Day().Do(task)
	<-gocron.Start()
}

func task() {
	channel := make(chan int)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// ===
	task1 := createCommandList(os.Getenv("TASK"))
	task2 := createCommandList(os.Getenv("TASK2"))
	task3 := createCommandList(os.Getenv("TASK3"))
	combinedQueue := [][]string{task1, task2, task3}
	// ===
	fmt.Println("Execution start.")
	for i, cmdList := range combinedQueue {
		if len(cmdList) == 0 {
			continue
		} else {
			go func(cmdList []string) {
				channel <- i
				executeCommandList(cmdList)
				go func() { writeToLogFile("Successfully rebooted.") }()
			}(cmdList)
			<-channel

		}
	}
	go func() { sendMail() }()
}

func executeCommandList(queue []string) {
	channel := make(chan int)
	for i, cmd := range queue {
		go func(cmd string) {
			channel <- i
			fmt.Printf("running %s \n", cmd)
			err := runCommand(cmd)
			time.Sleep(2 * time.Second)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				failedCmd := fmt.Sprintf("Failed to execute this command %s \n", cmd)
				writeToLogFile(failedCmd)
			}
		}(cmd)
		<-channel
		time.Sleep(2 * time.Second)
	}
}

func createCommandList(name string) []string {
	if name == "" {
		return []string{}
	}
	queue := []string{
		fmt.Sprintf("cd /home/%s", name),
		fmt.Sprintf("pm2 stop %s", name),
		"rimraf ./logs",
		fmt.Sprintf("pm2 start /home/%s/auto.js -n %s", name, name),
	}
	return queue
}
func runCommand(command string) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// fmt.Println(stdout.String())
	return nil
}
