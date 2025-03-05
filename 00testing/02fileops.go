package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// File Operations
	file, err := os.Create("./00testing/02file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	n, err := file.WriteString("Hello World@")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total written bytes n:", n)

	data, err := os.ReadFile("./00testing/02file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	//err = os.Remove("./00testing/02file.txt")
	//if err != nil {
	//	fmt.Println("Error deleting file:", err)
	//}

	// Directory Operations
	//err = os.Mkdir("./00testing/example_dir", 0755)
	//if err != nil {
	//	fmt.Println("Error creating directory:", err)
	//	return
	//}
	//
	//err = os.Rename("./00testing/example_dir", "./00testing/renamed_dir")
	//if err != nil {
	//	fmt.Println("Error renaming directory:", err)
	//	return
	//}
	//
	//err = os.Remove("./00testing/renamed_dir")
	//if err != nil {
	//	fmt.Println("Error removing directory:", err)
	//	return
	//}

	// Environment Variables
	err = os.Setenv("MY_VAR", "my_value")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}

	value := os.Getenv("MY_VAR")
	fmt.Println("MY_VAR:", value)

	err = os.Unsetenv("MY_VAR")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}

	//envVars := os.Environ()
	//for _, envVar := range envVars {
	//	fmt.Println(envVar)
	//}

	DirFile, err := os.Open("./00testing")
	if err != nil {
		fmt.Println(err)
	}
	defer DirFile.Close()
	names, err := DirFile.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
	}
	for _, name := range names {
		fmt.Println(name)
	}

	// Process management
	cmd := exec.Command("sleep", "1")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting process:", err)
		return
	}

	// Get the process ID
	pid := cmd.Process.Pid
	fmt.Println("Process ID:", pid)
	// Wait for the process to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for process:", err)
		return
	}
	fmt.Println("Process finished")

	hostname, _ := os.Hostname()
	userid := os.Geteuid()
	groupid := os.Getgid()
	fmt.Printf("Hostname %v, User %v, Group %v\n", hostname, userid, groupid)

	groups, _ := os.Getgroups()
	fmt.Println("Groups:", groups)

	pagesize := os.Getpagesize()
	pid = os.Getpid()
	ppid := os.Getppid()
	uid := os.Getuid()
	fmt.Printf("Pagesize %v, PID %v, PPID %v UID %v\n", pagesize, pid, ppid, uid)

	wd, _ := os.Getwd()
	fmt.Println("Working directory:", wd)

	f, err := os.OpenFile("./00testing/notes.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	//buff := make([]byte, 1024)
	//
	//dataBytes := f.Read(buff)
}
