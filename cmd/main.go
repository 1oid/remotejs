package main

import (
	"errors"
	"github.com/1oid/remotejs/pkgs/common"
	"github.com/1oid/remotejs/pkgs/page"
	"github.com/1oid/remotejs/pkgs/web"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func runner(c *cli.Context) error {
	var tasks = chromedp.Tasks{}
	var opts = page.NewOpts(
		chromedp.Flag("headless", false))

	if common.Vars.Url != "" {
		tasks = append(tasks, chromedp.Navigate(common.Vars.Url))
	}

	if common.Vars.ChromePath != "" {
		if _, err := os.Stat(common.Vars.ChromePath); err != nil {
			log.Fatal(errors.New("error chrome path"))
		}
		opts = append(opts, chromedp.ExecPath(common.Vars.ChromePath))
	}

	if common.Vars.Proxy != "" {
		opts = append(opts, chromedp.ProxyServer(common.Vars.Proxy))
	}

	var ctx = page.LaunchBrowser(tasks, opts, common.Vars.RemoteDebuggingAddr)

	web.WebRunner(common.Vars.WebListenPort, func(t string) (*runtime.RemoteObject, error) {
		return page.EvaluateCall(ctx, page.CallFrameId, t)
	})
	return nil
}

func main() {
	common.NewRunner(runner)
}
