package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"strings"
)



var validateCmd = &cobra.Command{
	Use: "validate",
	Short: "Validate Email Format",
	Long: `Validate Email Formate
For example:

email-validator-go validate -e email@gmail.com`,

	Run: validateEmail,
}

func init(){
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringP("email","e","","Email to validate")
}

func validateEmail(cmd *cobra.Command, args [] string){
	valid := true
	email,_ := cmd.Flags().GetString("email")


	if(email == ""){
		fmt.Println("No email provided, try again!")
	}

	trimmedEmail := strings.Trim(email, " ")

	if(strings.HasSuffix(trimmedEmail,".") || strings.HasPrefix(trimmedEmail,".")){
		valid = false 
	}

	if(strings.Contains(trimmedEmail,"<") || strings.Contains(trimmedEmail,">") || strings.Contains(trimmedEmail, "(") || strings.Contains(trimmedEmail,")")){
		valid = false
	}

	if(valid){
		fmt.Println("Valid Email")
	} else{
		fmt.Println("Invalid Email")
	}

}