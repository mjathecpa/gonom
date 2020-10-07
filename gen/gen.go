package gen

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "time"
)

// '''**************GENERAL**************'''
func ClearScreen() {
	var clear map[string]func() //create a map for storing clear funcs
	clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}
    
func TitleScreen(data_dict[string]) {
    // program title information
	title := 
    data_dict["title"] + " v" + data_dict["version"] + "\n" +
    data_dict["copyright"] + "\n" +
	data_dict["title"] + " is " + data_dict["description"] + "\n" +
	"Source distributable under the " + data_dict["license"]

	fmt.println(title)
}

func ExitFarewell(data_dict) {
    // program departure information
    fmt.println("\n" + data_dict["farewell"])
    input("press ENTER to close")
	ClearScreen()
}

func IsoDate() string {
    // request, test, and return date string
    while True:
        // ask for date based on question option
        ipdate = input('Enter the pay date (YYYY-MM-DD):\n')
        // parse input date
        try:
            ipdate = parse(ipdate).strftime("%F")
            break
        except:
            fmt.println("Invalid date. Valid options include:\n",
                  "ISO date format: YYYY-MM-DD or YYYYMMDD\n",
                  "US-Style Numeric: MM/DD/YYYY or MM/DD/YY\n",
                  "Month as word: May 8, 1945 or 1 July 2017")
	return ipdate
}