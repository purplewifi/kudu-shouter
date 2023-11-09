/*
Copyright © 2023 Jack Stupple <jack.stupple@purple.ai>
*/
package cmd

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"text/template"

	"github.com/containrrr/shoutrrr"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	kudushouter "github.com/surdaft/kudu-shouter/kudu-shouter"
)

var (
	httpAddress string
	serviceURLs []string

	messageTemplate string = `
		{{.Deployer}} just created a new deployment:
		Site: {{.HostName}}
		Success: {{.Status}}
	`

	commit  = ""
	version = "dev"
	date    = ""
	builtBy = ""

	rootCmd = &cobra.Command{
		Use:     "kudu-shouter",
		Version: fmt.Sprintf("kudu-shouter - v%s (%s) by %s - %s", version, commit, builtBy, date),
		Short:   "Forward kudu webhooks via shoutrrr",
		Run: func(cmd *cobra.Command, args []string) {
			sender, err := shoutrrr.CreateSender(serviceURLs...)
			if err != nil {
				slog.Error("error creating sender", slog.Any("error", err.Error()))
				return
			}

			r := gin.Default()

			r.POST("/capture", func(ctx *gin.Context) {
				var payload kudushouter.Payload
				ctx.BindJSON(&payload)

				t, err := template.New("message").Parse(messageTemplate)
				if err != nil {
					handleServerErr(ctx, "error creating template", err)
					return
				}

				msg := bytes.NewBufferString("")
				err = t.Execute(msg, payload)
				if err != nil {
					handleServerErr(ctx, "error executing template", err)
					return
				}

				sender.SendAsync(msg.String(), nil)
				ctx.JSON(http.StatusOK, map[string]any{"success": true})
			})

			r.GET("/health", func(ctx *gin.Context) {
				ctx.String(200, "OK")
			})

			r.Run(httpAddress)
		},
	}
)

func handleServerErr(ctx *gin.Context, msg string, err error) {
	slog.Error(msg, slog.Any("error", err.Error()))
	ctx.JSON(http.StatusInternalServerError, map[string]any{"error": "failed to forward webhook"})
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&httpAddress, "http-address", ":7890", "HTTP address to serve on")
	rootCmd.Flags().StringArrayVar(&serviceURLs, "service-url", []string{}, "use multiple times for each service url")
}
