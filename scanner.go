package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)
 
func main() {
	file, err := os.Open("scanner1.txt")
	f, err := os.Create("file.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	
	if err != nil {
        log.Fatal(err)
    }
	
	defer f.Close()
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range txtlines {
		_, err2 := f.WriteString(eachline + "\n")
		fmt.Println(eachline)	
		
		if err2 != nil {
			log.Fatal(err2)
		}
		
		time.Sleep(2 * time.Second)
		
	}
	
	//fmt.Println("done")
	
}
