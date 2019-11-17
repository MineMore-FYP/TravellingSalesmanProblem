package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "time"
  "strconv"
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
	//fmt.Println("PSO complete")


	//import starting LOWER BOUND OF PSO ITERATIONS from user script
	cmd_lb_iterations := exec.Command("python", "-c", "import userScript; print userScript.lb_iterations")
	out_lb_iterations,err_lb_iterations := cmd_lb_iterations.CombinedOutput()
	if err_lb_iterations != nil {
		fmt.Println(err_lb_iterations)
	}

	lb_iterations := string(out_lb_iterations)[:len(out_lb_iterations)-1]
	int_lb_iterations, err_int_lb_iterations := strconv.Atoi(lb_iterations)
	if err_int_lb_iterations != nil {
		fmt.Println(err_int_lb_iterations)
	}

	fmt.Println("lower bound of iterations", int_lb_iterations)


	//import starting UPPER BOUND OF PSO ITERATIONS from user script
	cmd_ub_iterations := exec.Command("python", "-c", "import userScript; print userScript.ub_iterations")
	out_ub_iterations,err_ub_iterations := cmd_ub_iterations.CombinedOutput()
	if err_ub_iterations != nil {
		fmt.Println(err_ub_iterations)
	}

	ub_iterations := string(out_ub_iterations)[:len(out_ub_iterations)-1]
	int_ub_iterations, err_int_ub_iterations := strconv.Atoi(ub_iterations)
	if err_int_ub_iterations != nil {
		fmt.Println(err_int_ub_iterations)
	}

	fmt.Println("upper bound of iterations", int_ub_iterations)


	//import starting LOWER BOUND OF PSO POPULATION SIZE from user script
	cmd_lb_size_population := exec.Command("python", "-c", "import userScript; print userScript.lb_size_population")
	out_lb_size_population,err_lb_size_population := cmd_lb_size_population.CombinedOutput()
	if err_lb_size_population != nil {
		fmt.Println(err_lb_size_population)
	}

	lb_size_population := string(out_lb_size_population)[:len(out_lb_size_population)-1]
	int_lb_size_population, err_int_lb_size_population := strconv.Atoi(lb_size_population)
	if err_int_lb_size_population != nil {
		fmt.Println(err_int_lb_size_population)
	}

	fmt.Println("lower bound of population size", int_lb_size_population)


	//import starting UPPER BOUND OF PSO POPULATION SIZE from user script
	cmd_ub_size_population := exec.Command("python", "-c", "import userScript; print userScript.ub_size_population")
	out_ub_size_population,err_ub_size_population := cmd_ub_size_population.CombinedOutput()
	if err_ub_size_population != nil {
		fmt.Println(err_ub_size_population)
	}

	ub_size_population := string(out_ub_size_population)[:len(out_ub_size_population)-1]
	int_ub_size_population, err_int_ub_size_population := strconv.Atoi(ub_size_population)
	if err_int_ub_size_population != nil {
		fmt.Println(err_int_ub_size_population)
	}

	fmt.Println("upper bound of population size", int_ub_size_population)


	//import starting LOWER BOUND OF PSO BETA from user script
	cmd_lb_beta := exec.Command("python", "-c", "import userScript; print userScript.lb_beta")
	out_lb_beta,err_lb_beta := cmd_lb_beta.CombinedOutput()
	if err_lb_beta != nil {
		fmt.Println(err_lb_beta)
	}

	lb_beta := string(out_lb_beta)[:len(out_lb_beta)-1]
	int_lb_beta, err_int_lb_beta := strconv.ParseFloat(lb_beta, 64)
	if err_int_lb_beta != nil {
		fmt.Println(err_int_lb_beta)
	}

	fmt.Println("lower bound of beta", int_lb_beta)


	//import starting UPPER BOUND OF PSO BETA from user script
	cmd_ub_beta := exec.Command("python", "-c", "import userScript; print userScript.ub_beta")
	out_ub_beta,err_ub_beta := cmd_ub_beta.CombinedOutput()
	if err_ub_beta != nil {
		fmt.Println(err_ub_beta)
	}

	ub_beta := string(out_ub_beta)[:len(out_ub_beta)-1]
	int_ub_beta, err_int_ub_beta := strconv.ParseFloat(ub_beta, 64)
	if err_int_ub_beta != nil {
		fmt.Println(err_int_ub_beta)
	}

	fmt.Println("upper bound of beta", int_ub_beta)


	//import starting LOWER BOUND OF PSO ALPHA from user script
	cmd_lb_alfa := exec.Command("python", "-c", "import userScript; print userScript.lb_alfa")
	out_lb_alfa,err_lb_alfa := cmd_lb_alfa.CombinedOutput()
	if err_lb_alfa != nil {
		fmt.Println(err_lb_alfa)
	}

	lb_alfa := string(out_lb_alfa)[:len(out_lb_alfa)-1]
	int_lb_alfa, err_int_lb_alfa := strconv.ParseFloat(lb_alfa, 64)
	if err_int_lb_alfa != nil {
		fmt.Println(err_int_lb_alfa)
	}

	fmt.Println("lower bound of alpha", int_lb_alfa)


	//import starting LOWER BOUND OF PSO ALPHA from user script
	cmd_ub_alfa := exec.Command("python", "-c", "import userScript; print userScript.ub_alfa")
	out_ub_alfa,err_ub_alfa := cmd_ub_alfa.CombinedOutput()
	if err_ub_alfa != nil {
		fmt.Println(err_ub_alfa)
	}

	ub_alfa := string(out_ub_alfa)[:len(out_ub_alfa)-1]
	int_ub_alfa, err_int_ub_alfa := strconv.ParseFloat(ub_alfa, 64)
	if err_int_ub_alfa != nil {
		fmt.Println(err_int_ub_alfa)
	}

	fmt.Println("upper bound of alpha", int_ub_alfa)

	
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


