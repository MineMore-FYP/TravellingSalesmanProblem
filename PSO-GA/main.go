package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "time"
  userVariableImports "./userVariableImports"
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
	//go pythonCall("tsp_pso.py")
	fmt.Println("PSO complete")


	userVariableImports.DefinePSOVariables()
	//fmt.Println(L.int_lb_iterations)
	
/*

	for i := startWithNumberOfClustersInt; i <= endWithNumberOfClustersInt; i++ {
		go pythonCall("tsp_pso.py")
		fmt.Println("PSO complete ", i)
	}

	//go pythonCall("tsp_ga.py")
	//fmt.Println("GA complete")

	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Particle Swarm Optimization complete for all iterations")*/

	time.Sleep(1 * time.Millisecond)


}


