package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "time"
  "encoding/csv"
  "io"	
  //"strconv"
  //userVariableImports "./userVariableImports"
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
	

	go pythonCall("tsp_pso.py")
	fmt.Println("PSO 1st level started")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("PSO 1st level ended")
	
	
	// Open the file
	csvfile, err := os.Open("/home/mpiuser/Documents/FYP/TravellingSalesmanProblem/PSO-GA/savedFiles/pso_instances.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		path := record[4]
		cost := record[5]
		fmt.Printf(paths)
	}

	

}


