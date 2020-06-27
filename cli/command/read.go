package command

import (
	"fmt"
	request "github.com/alessiosavi/Requests"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

func init(){
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use: "read",
	Short: "get single users with parameter id",
	Run: read,
}

func read(cmd *cobra.Command,args []string){
	baseUrl := fmt.Sprintf("%v",viper.Get("baseurl"))
	token := fmt.Sprintf("%v",viper.Get("token"))
	url := baseUrl+`/user/`+args[0]
	var req request.Request
	resp := req.SendRequest(url,"GET",nil,[]string{`Accept`, `application/json`, "Authorization", "Bearer "+ token},false, time.Duration(time.Duration.Milliseconds(1000)))
	fmt.Println(string(resp.Body))
}

