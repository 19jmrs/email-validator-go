package cmd

import (
	"fmt"
	"net/mail"

	"github.com/spf13/cobra"


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
	
	email,_ := cmd.Flags().GetString("email")
	_, error := mail.ParseAddress(email)
	isValid := false

	if(error != nil){
		fmt.Println(error)
	} else{
		isValid = true
	}

	if(isValid){
		fmt.Printf("valid email")
	}else {
		fmt.Printf("not valid email")
	}

}