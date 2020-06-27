package command

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init(){
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use: "edit",
	Short: "edit users param 1 id, param 2 full name, param 3 email, param 4 mobile phone, param 5 password",
	Run: edit,
}

func edit(cmd *cobra.Command,args []string){
	baseUrl := fmt.Sprintf("%v",viper.Get("baseurl"))
	token := fmt.Sprintf("%v",viper.Get("token"))
	url := baseUrl+`/user/`+args[0]
	reqBody := map[string]interface{}{}

	req:= HttpRequest.NewRequest().Debug(false).SetHeaders(map[string]string{
		"Content-Type":"application/json",
		"Authorization":"Bearer "+token,
	}).SetTimeout(5)

	if len(args)>5{
		reqBody =map[string]interface{}{
			"full_name":args[1],
			"email":args[2],
			"password":args[4],
			"mobile_phone":args[3],
		}
	}else{
		reqBody =map[string]interface{}{
			"full_name":args[1],
			"email":args[2],
			"mobile_phone":args[3],
		}
	}
	res,err := req.Put(url,reqBody)
	body,err := res.Body()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

