package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "time"
  "strconv"
  userVariableImports "./userVariableImports"
)

func SendValue(s string, c chan string){
	//send value through channel c
	c <- s
}

/*
func pythonCall(progName string){
	cmd := exec.Command("python3", progName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os. Stderr
	log.Println(cmd.Run())
}
*/

func pythonCallFourParams(progName string, para1 string, para2 string, para3 string, para4 string) {
	cmd := exec.Command("python3", progName, para1, para2, para3, para4)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}

func main() {
	

	int_lb_iterations := userVariableImports.DefineLBiterationsPSO()
    	fmt.Println("lower bound of iterations", int_lb_iterations)

	int_ub_iterations := userVariableImports.DefineUBiterationsPSO()
    	fmt.Println("upper bound of iterations", int_ub_iterations)

	int_lb_size_population := userVariableImports.DefineLBpopulationPSO()
	fmt.Println("lower bound of population size", int_lb_size_population)

	int_ub_size_population := userVariableImports.DefineUBpopulationPSO()
	fmt.Println("upper bound of population size", int_ub_size_population)

	float_lb_beta := userVariableImports.DefineLBbetaPSO()
	fmt.Println("lower bound of beta", float_lb_beta)

	float_ub_beta := userVariableImports.DefineUBbetaPSO()
	fmt.Println("upper bound of beta", float_ub_beta)

	float_lb_alfa := userVariableImports.DefineLBalfaPSO()
	fmt.Println("lower bound of alpha", float_lb_alfa)

	float_ub_alfa := userVariableImports.DefineUBalfaPSO()
	fmt.Println("upper bound of alpha", float_ub_alfa)

	strLBiterations := strconv.Itoa(int_lb_iterations)
	strLBpopulation := strconv.Itoa(int_lb_size_population)

	strLBbeta := fmt.Sprintf("%f", float_lb_beta)
	strLBalfa := fmt.Sprintf("%f", float_lb_alfa)

	go pythonCallFourParams("tsp_pso.py", strLBiterations, strLBpopulation, strLBbeta, strLBalfa)
	
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("PSO complete")

/*

	for i := startWithNumberOfClustersInt; i <= endWithNumberOfClustersInt; i++ {
		go pythonCall("tsp_pso.py")
		fmt.Println("PSO complete ", i)
	}

	//go pythonCall("tsp_ga.py")
	//fmt.Println("GA complete")

	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Particle Swarm Optimization complete for all iterations")*/

	


}


