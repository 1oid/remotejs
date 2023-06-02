package page

import (
	"context"
	"errors"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/debugger"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
)

var CallFrameId debugger.CallFrameID

func addListener(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev.(type) {
		case *debugger.EventPaused:
			var d = ev.(*debugger.EventPaused)

			if len(d.CallFrames) > 0 {
				CallFrameId = d.CallFrames[0].CallFrameID
			}
		}
	})
}

func getContext(ctx context.Context) (*chromedp.Context, error) {
	c := chromedp.FromContext(ctx)
	// If c is nil, it's not a chromedp context.
	// If c.Allocator is nil, NewContext wasn't used properly.
	// If c.cancel is nil, Run is being called directly with an allocator
	// context.
	if c == nil || c.Allocator == nil {
		return nil, chromedp.ErrInvalidContext
	}
	return c, nil
}

func EvaluateCall(ctx context.Context, callFrameId debugger.CallFrameID, expression string) (*runtime.RemoteObject, error) {
	c, _ := getContext(ctx)
	call := &debugger.EvaluateOnCallFrameParams{
		CallFrameID:           callFrameId,
		Expression:            expression,
		ObjectGroup:           "console",
		IncludeCommandLineAPI: true,
		ReturnByValue:         true,
	}
	var result, _, err = call.Do(cdp.WithExecutor(ctx, c.Target))
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("execute error")
	}
	if result.Subtype == runtime.SubtypeError {
		return nil, errors.New(result.Description)
	}
	return result, nil

}

func NewOpts(opts ...chromedp.ExecAllocatorOption) []chromedp.ExecAllocatorOption {
	return append(chromedp.DefaultExecAllocatorOptions[:], opts...)
}

func LaunchBrowser(tasks chromedp.Tasks, opts []chromedp.ExecAllocatorOption, remote string) context.Context {

	var allocCtx context.Context

	if remote != "" {
		allocCtx, _ = chromedp.NewRemoteAllocator(context.Background(), remote)
	} else {
		allocCtx, _ = chromedp.NewExecAllocator(context.Background(), opts...)
	}
	var ctx, _ = chromedp.NewContext(
		allocCtx, chromedp.WithLogf(log.Printf))

	err := chromedp.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	chromedp.Run(ctx, tasks, chromedp.ActionFunc(func(ctx context.Context) error {
		_, err := debugger.Enable().Do(ctx)
		return err
	}))
	go addListener(ctx)

	return ctx
}
