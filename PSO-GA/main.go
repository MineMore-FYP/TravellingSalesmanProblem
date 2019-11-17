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

func pythonCall(progName string){
	cmd := exec.Command("python3", progName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os. Stderr
	log.Println(cmd.Run())
}


func pythonCallFourParams(progName string, para1 string, para2 string, para3 string, para4 string) {
	cmd := exec.Command("python3", progName, para1, para2, para3, para4)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}

func main() {

	int_lb_iterations := userVariableImports.DefineLBiterationsPSO()
   
	int_ub_iterations := userVariableImports.DefineUBiterationsPSO()
    	
	int_lb_size_population := userVariableImports.DefineLBpopulationPSO()
	
	int_ub_size_population := userVariableImports.DefineUBpopulationPSO()
	
	float_lb_beta := userVariableImports.DefineLBbetaPSO()
	
	float_ub_beta := userVariableImports.DefineUBbetaPSO()
	
	float_lb_alfa := userVariableImports.DefineLBalfaPSO()
	
	float_ub_alfa := userVariableImports.DefineUBalfaPSO()

	for i := int_lb_iterations; i <= int_ub_iterations; i++ {
		for j := int_lb_size_population; j <= int_ub_size_population; j++ {
			for k := float_lb_beta; k <= float_ub_beta ; k += 0.1 {
				for l := float_lb_alfa ; l <= float_ub_alfa ; l += 0.1{
					iStr := strconv.Itoa(i)
					jStr := strconv.Itoa(j)
					kStr := fmt.Sprintf("%.1f", k)
					lStr := fmt.Sprintf("%.1f", l)
					go pythonCallFourParams("tsp_pso.py", iStr, jStr, kStr, lStr)
					//time.Sleep(1500 * time.Millisecond)
					fmt.Println("PSO complete. ITERATIONS = ", iStr, " POPULATION SIZE = ", jStr, " BETA = ", kStr, " ALPHA = ", lStr)
				}
			}
		}
	}
	
	time.Sleep(10000 * time.Millisecond)

	go pythonCall("tsp_ga.py")
	fmt.Println("GA complete")
	time.Sleep(10000 * time.Millisecond)

}


