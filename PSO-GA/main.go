package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "time"
)

func SendValue(s string, c chan string){
	//send value through channel c
	c <- s
}

func pythonCall(progName string){
	cmd := exec.Command("python3", progName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os. Stderr
	log.Println(cmd.Run())
}

func main() {
	go pythonCall("tsp_pso.py")
	fmt.Println("PSO complete")

	go pythonCall("tsp_ga.py")
	fmt.Println("GA complete")

	time.Sleep(10000 * time.Millisecond)
}


