/*
Copyright © 2020 Denis Rendler <connect@rendler.me>

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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/koderhut/safenotes/internal/utilities/logs"
	"github.com/koderhut/safenotes/webapp/contracts"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Retrieve the webservice stats",
	Long: `Retrieve the stats available from the `,
	Run: func(cmd *cobra.Command, args []string) {
		statsUrl := url.URL{
			Scheme: "http",
			Host: fmt.Sprintf("%s:%s","localhost", cfg.Server.Port),
			Path: "/stats",
			User: url.UserPassword(cfg.Server.Auth.User, cfg.Server.Auth.Pass),
		}

		req, err := http.NewRequest(http.MethodGet, statsUrl.String(), nil)
		if err != nil {
			logs.Writer.Critical(err.Error())
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logs.Writer.Critical(err.Error())
		}

		var result contracts.StatsMessage

		json.NewDecoder(resp.Body).Decode(&result)

		logs.Writer.Info(fmt.Sprintf("Stats for safenotes service:\n\nCurrent Number of stored notes: %d\nTotal number of stored notes: %d\n", result.StoredNotes, result.TotalNotes))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
