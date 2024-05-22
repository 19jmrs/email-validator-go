package main

import(
	"fmt"
	"io/ioutil"
	"os"
)

func main(){

	// access directory
	files, err := ioutil.ReadDir("test")

	if err != nil {
		fmt.Println(err)
	}

	//display files in directory
	for _, file := range files {
		//display name anc check if it is a Directory
		fmt.Println(file.Name(), file.IsDir())
	}
	
	
	fmt.Println("Input the file name to move")
	var filename string
	fmt.Scanln(&filename)
	//move file from test to Folder directory

	err = os.Rename("test/"+filename, "test/Folder/"+filename)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("File Moved Successfully")
	}
}