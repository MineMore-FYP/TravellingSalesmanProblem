package userVariableImports

import (
  "os/exec"
  "fmt"
  "strconv"
)

func DefineVariable(importCommand string) string{
	cmd_str := "import userScript; print userScript." + importCommand
	cmd_importCommand := exec.Command("python", "-c", cmd_str)
	out_importCommand, err_importCommand := cmd_importCommand.CombinedOutput()
	if err_importCommand != nil {
		fmt.Println(err_importCommand)
	}

	importedVariable := string(out_importCommand)[:len(out_importCommand)-1]
	//int_lb_iterations, err_int_lb_iterations := strconv.Atoi(lb_iterations)
	//if err_int_lb_iterations != nil {
	//	fmt.Println(err_int_lb_iterations)
	//}

	fmt.Println(importedVariable)
	
	return importedVariable
	//return int_lb_iterations
}

func ImportAsInt(command string) int{
	str_x := DefineVariable(command)
	int_x, err_x := strconv.Atoi(str_x)	
	if err_x != nil {
		fmt.Println(err_x)
	}
	return int_x	
}


func ImportAsFloat(command string) float64{
	str_x := DefineVariable(command)
	float_x, err_x := strconv.ParseFloat(str_x, 64)	
	if err_x != nil {
		fmt.Println(err_x)
	}
	return float_x	
}


