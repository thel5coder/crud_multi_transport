package command

import (
	"fmt"
    "github.com/kirinlabs/HttpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new users param 1 full name, param 1 email, param 2 mobile phone, param 3 password",
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	baseUrl := fmt.Sprintf("%v", viper.Get("baseurl"))
	token := fmt.Sprintf("%v", viper.Get("token"))
	url := baseUrl + `/user`

	req:= HttpRequest.NewRequest().Debug(false).SetHeaders(map[string]string{
		"Content-Type":"application/json",
		"Authorization":"Bearer "+token,
	}).SetTimeout(5)

	reqBody :=map[string]interface{}{
		"full_name":args[0],
		"email":args[1],
		"mobile_phone":args[2],
		"password":args[3],
	}

	res,err := req.Post(url,reqBody)
	body,err := res.Body()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}