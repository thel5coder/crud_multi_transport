package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)


func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use: "login",
	Short: "Get jwt token this command need two parameter email and password,use email and password as parameter separate with space character",
	Run: authLogin,
}

func authLogin(cmd *cobra.Command, args []string) {
	const url = "http://localhost:3001/api/v1/auth/login"
	if len(args)==0{
		fmt.Println("You need parse email and password")
		os.Exit(1)
	}

	requestBody,err := json.Marshal(map[string]interface{}{
		"email":fmt.Sprintf("%v",args[0]),
		"password":fmt.Sprintf("%v",args[1]),
	})

	response,err := http.Post(url,"application/json",bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	fmt.Println(result)
	token := result["data"].(map[string]interface{})["token"]
	viper.Set("token",token)
	viper.WriteConfig()
	log.Println("Token saved")
}
