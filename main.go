package main
import(
	
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"bufio"
	
);

func main(){
	clear();
	choise();
}

func createfile(nameFile string){
	f, _ := os.Create(nameFile);
	f, _ = os.OpenFile(nameFile,  os.O_WRONLY, 0666);
	inputdata := bufio.NewScanner(os.Stdin);
	for;;{
		inputdata.Scan();
		if inputdata.Text() == "!"{
			break;
		}
		f.WriteString(inputdata.Text() + "\n");

	}
	defer f.Close();
		
}



func readfile(nameFile string) error{
	file, e := os.Open(nameFile); 
	if e != nil{
		return e;
	}
	bytes, _:= ioutil.ReadAll(file);
	fmt.Print(string(bytes));
	file.Close();
	return nil;
}

func clear(){
	cmd := exec.Command("clear");
	cmd.Stdin = os.Stdin;
	cmd.Stdout = os.Stdout;
	cmd.Run();
}

func choise(){
	cmd := os.Args;
	l := len(cmd);
	if l <= 2{
		help();
		os.Exit(1);
	}
	switch (cmd[1]){
	case "-c", "--create":
		if l == 3 {
			createfile(cmd[2]);
		}else{
			help();
			os.Exit(1);
		}
	case "-o", "--open":
		if l == 3 {
			readfile(cmd[2]);
		}else{
			help();
			os.Exit(1);
		}
	case "-h", "--help":
		help();
		
	default:
		help();

		
	}
}

func help(){
	fmt.Println();
	fmt.Println("");
	fmt.Println("\tblackout <options> file");
	fmt.Println("");
	fmt.Println("\tOPTIONS:");
	fmt.Println("");
	fmt.Println("\t-o, --open\tOpen file");
	fmt.Println("\t-c, --create\tCreate file. Sign ! for exit");
	fmt.Println("\t-h, --help\tThis Window");
	fmt.Println();
	fmt.Println();
}