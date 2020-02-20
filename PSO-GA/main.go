package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  //"time"
  "encoding/csv"
  "encoding/json"
  "io/ioutil"
  //"io"
  "strconv"
  //"bufio"
  //"strings"
  //"reflect"
  //userVariableImports "./userVariableImports"
)

func simplepythonCall(progName string, step string){
	cmd := exec.Command("python3", progName, step)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os. Stderr
	log.Println(cmd.Run())
}
/*
func pythonCall(progName string, inChannel chan <- string, step string) {
	cmd := exec.Command("python3", progName, step)
	out, err := cmd.CombinedOutput()
	log.Println(cmd.Run())

	if err != nil {
		fmt.Println(err)
		// Exit with status 3.
    os.Exit(3)
	}
	fmt.Println(string(out))
	//check if msg is legit
	msg := string(out)[:len(out)-1]
	//msg := ("Module Completed: " + progName)
	inChannel <- msg
}
*/
func messagePassing(inChannel <- chan string, outChannel chan <- string ){
	msg := <- inChannel
	outChannel <- msg
}

type Cost_class struct {
    Path string `json:"path"`
    Cost float64 `json:"cost"`
}

func FindMinCost(Cost_set []Cost_class) (max Cost_class) {

	min := Cost_set[0]
	for _, cost_obj := range Cost_set {

		if cost_obj.Cost < min.Cost {
			min = cost_obj
			if cost_obj.Path == "Path" {
				continue
			}
		}
	}
	return min
}

func removeIt(ss Cost_class, ssSlice []Cost_class) []Cost_class {
    for idx, v := range ssSlice {
        if v == ss {
            return append(ssSlice[0:idx], ssSlice[idx+1:]...)
        }
    }
    return ssSlice
}

func costSelection (step string) (min Cost_class) {
	fmt.Println("Cost selection started")
	//var files []string
	cmd := exec.Command("python", "-c", "import userScript; print userScript.output")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		// Exit with status 3.
		os.Exit(3)
	} else if out == nil{
		os.Exit(3)
	}

	costJsonFile := string(out)[:len(out)-1] + step +"_costJson.json"
	//csvFile, _ := os.Open("/home/mpiuser/Documents/FYP/TravellingSalesmanProblem/PSO-GA/savedFiles/pso_instances.csv")

    	csvfile, err := os.Open(string(out)[:len(out)-1] + step + "_pso_instances.csv")
   	checkError(err)
   	defer csvfile.Close()

   	reader := csv.NewReader(csvfile)

   	reader.FieldsPerRecord = -1

   	rawCSVdata, err := reader.ReadAll()

   	if err != nil {
      		fmt.Println(err)
      		os.Exit(1)
   	}

   	var cost_obj Cost_class
   	var cost_set []Cost_class

   	for _, record := range rawCSVdata {
		//fmt.Println(record[4])
		//fmt.Println(reflect.TypeOf(record[4]))
      		cost_obj.Path = record[4]
      		cost_obj.Cost, _ = strconv.ParseFloat(record[5],64)
      		cost_set = append(cost_set, cost_obj)
   	}

   	//fmt.Println(cost_set)
	cost_set = removeIt(Cost_class{"Path", 0}, cost_set)

	min = FindMinCost(cost_set)
	writeCostFile(min, costJsonFile, )
	//fmt.Println(min)
	return min

}

func checkError(err error) {
   if err != nil {
      fmt.Println("Error:",err)
      os.Exit(-1)
   }
}

func writeCostFile(cost_obj Cost_class, costJsonFile string) {

    costJson, _ := json.Marshal(cost_obj)
    ioutil.WriteFile(costJsonFile, costJson, 0644)
    fmt.Println(string(costJson))
}


func pythonCall(progName string, sendChannel chan <- string, itr string) {
	cmd := exec.Command("python3", progName, itr)
	out, err := cmd.CombinedOutput()
	log.Println(cmd.Run())

	if err != nil {
		fmt.Println(err)
		// Exit with status 3.
    		os.Exit(3)
	}
	
	fmt.Println(string(out))
	msg := string(out)[:len(out)-1]
	sendChannel <- msg
}


func loop(program string, goSteps int){
	for i:=1; i<=goSteps; i++ {
		sendChannel := make(chan string, 1)
		receiveChannel := make(chan string, 1)
		pythonCall(program, sendChannel, strconv.Itoa(i))
		minCostObj := costSelection(strconv.Itoa(i))
    		fmt.Println(minCostObj)
		msg := <- sendChannel 	
		receiveChannel <- msg
	}
	
}

func main() {
	loop("tsp_pso.py", 10)
}
