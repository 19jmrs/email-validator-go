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

	//move file from test to Folder directory

	err = os.Rename("test/newname.txt", "test/Folder/newname.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
}