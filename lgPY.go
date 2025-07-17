package main

import (
	"log"
	//"time"

	"github.com/playwright-community/playwright-go"
)

func main() {
	// Start Playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	defer pw.Stop()

	// Launch headless chromium
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args:     []string{"--no-sandbox", "--disable-dev-shm-usage"},
	})
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navigate to the URL
	_, err = page.Goto("http://www.lg4all.com/pod/?Code=IN058035001H")
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// Handle alert popup if present
	// Playwright Go doesn't auto accept alerts; You need to register a handler
	page.OnDialog(func(dialog playwright.Dialog) {
		log.Printf("Alert text: %s", dialog.Message())
		dialog.Accept()
	})

	// Wait a little for a possible alert
	//time.Sleep(2 * time.Second)

	// Check for "No Record Found"
	noRecord, err := page.QuerySelector("//td[contains(text(), 'No Record Found')]")
	if err == nil && noRecord != nil {
		log.Println("No Record Found message. Exiting.")
		return
	}

	// Find all checkboxes
	checkboxes, err := page.QuerySelectorAll("//input[@type='checkbox']")
	if err != nil {
		log.Fatalf("checkboxes lookup failed: %v", err)
	}

	for _, checkbox := range checkboxes {
		checked, _ := checkbox.IsChecked()
		if !checked {
			checkbox.Click()
		}
	}

	// Find and click submit button
	submitBtn, err := page.QuerySelector("//input[@type='submit' or @type='button' and contains(@value, 'Submit')]")
	if err == nil && submitBtn != nil {
		err = submitBtn.Click()
		if err != nil {
			log.Println("Submit button click failed:", err)
		} else {
			log.Println("Submit button clicked.")
		}
	} else {
		log.Println("No submit button found.")
	}

    noBID, err := page.QuerySelector("#ContentPlaceHolder1_lblNoOfBid")
    if err != nil {
        log.Fatalf("Failed to locate NoOfBid element: %v", err)
    }
    if noBID != nil {
        text, err := noBID.TextContent()
        if err != nil {
            log.Printf("Failed to get text: %v", err)
        } else {
            log.Printf("No Of Bid Text: %s", text)
        }
    } else {
        log.Println("NoOfBid element not found on the page.")
    }


	// End
}
