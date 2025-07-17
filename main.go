package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/page"
)

func main() {
    // Create Chrome allocator with default options (non-headless for visible browser)
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        // Comment out or remove the following line to see browser UI:
        // chromedp.Headless,
        chromedp.DisableGPU,
        chromedp.NoSandbox,
        chromedp.Flag("window-size", "1200,800"),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    // Create context
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    // Optional: timeout for entire operation
    ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
    defer cancel()

    // Listen and close JavaScript alert/pop-ups automatically
    chromedp.ListenTarget(ctx, func(ev interface{}) {
        if e, ok := ev.(*page.EventJavascriptDialogOpening); ok {
            fmt.Printf("Alert detected with message: %s\n", e.Message)
            go func() {
                _ = chromedp.Run(ctx, page.HandleJavaScriptDialog(true))
            }()
        }
    })

    url := "https://www.google.com"
    var pageBody string

    err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.WaitVisible("body", chromedp.ByQuery),
        chromedp.Text("body", &pageBody, chromedp.NodeVisible),
    )
    if err != nil {
        log.Fatalf("Failed to load the page: %v", err)
    }

    fmt.Println("Page content snippet:")
    // Print first 500 characters or less
    if len(pageBody) > 500 {
        fmt.Println(pageBody[:500])
    } else {
        fmt.Println(pageBody)
    }
}
