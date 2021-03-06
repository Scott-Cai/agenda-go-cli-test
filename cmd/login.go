package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "For User login",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Login called")
		tmp_u, _ := cmd.Flags().GetString("username")
		tmp_p, _ := cmd.Flags().GetString("password")
		if tmp_u == "" || tmp_p == "" {
			fmt.Println("Please input both username and password")
			return
		}
		if _, flag := service.GetCurUser(); flag == true {
			fmt.Println("Please logout firstly!")
			return
		}
		if tf := service.UserLogin(tmp_u, tmp_p); tf == true {
			fmt.Println("Login Successfully. Current User: ", tmp_u)
		} else {
			fmt.Println("Login fail: Wrong username or password")
		}
		return
	},
}

var (
	username *string
	password *string
)
func init() {
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	username = loginCmd.Flags().StringP("username", "u", "", "agenda username")
	password = loginCmd.Flags().StringP("password", "p","","agenda password")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
