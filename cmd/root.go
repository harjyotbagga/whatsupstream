/*
Copyright © 2020 Yashvardhan Kukreja <yash.kukreja.98@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"whatsupstream/pkg/cmd/whatsupstream/notify"
	"whatsupstream/pkg/cmd/whatsupstream/stop"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "whatsupstream",
	Short: "Stay up-to-date and notified with your favorite OSS repos with whatsupstream",
	Long: `whatsupstream can keep you up-to-date and notified about the activity associated with any OSSN repository.
Say, there is an active OSS repository and you want to take up a good first issue as soon as one opens up.
Now, in real life, this is tough because you can't just keep on checking the issues of that repo every minute.
And someone might take up that issue before you too.
So, whatsupstream can notify you in almost real-time whenever such issue comes up.
And you can tweak whatsupstream to notify you for any other kind or category of issues.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World! Welcome to What'supstream!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(notify.NewCommand())
	rootCmd.AddCommand(stop.NewCommand())
}
