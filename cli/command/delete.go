package command

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init(){
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "delete need id parameter",
	Run: delete,
}

func delete(cmd *cobra.Command,args []string){
	baseUrl := fmt.Sprintf("%v",viper.Get("baseurl"))
	token := fmt.Sprintf("%v",viper.Get("token"))
	url := baseUrl+`/user/`+args[0]

	req:= HttpRequest.NewRequest().Debug(false).SetHeaders(map[string]string{
		"Content-Type":"application/json",
		"Authorization":"Bearer "+token,
	}).SetTimeout(5)

	res,err := req.Delete(url)
	body,err := res.Body()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

