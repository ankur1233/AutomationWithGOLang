package main

import (
	"log"
	"sync"
	//"time"

	"github.com/playwright-community/playwright-go"
)

func runAutomation(code string, wg *sync.WaitGroup) {
	defer wg.Done()
	//start := time.Now()

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		Args:     []string{
                      "--no-sandbox",
                      "--disable-dev-shm-usage",
                      "--disable-gpu",
                      "--disable-extensions",
                      "--disable-software-rasterizer",
                      "--disable-setuid-sandbox",
                  },
	})
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Example navigation using code
	_, err = page.Goto("http://www.lg4all.com/pod/?Code=" + code)
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	page.OnDialog(func(dialog playwright.Dialog) {
		log.Printf("Alert text: %s", dialog.Message())
		dialog.Accept()
	})
	//time.Sleep(2 * time.Second)

	noRecord, err := page.QuerySelector("//td[contains(text(), 'No Record Found')]")
	//duration := time.Since(start)
    //log.Printf("Code: %s | Time taken: %.2fs", code, duration.Seconds())
	if err == nil && noRecord != nil {
		log.Println("No Record Found message. Exiting.")
		return
	}

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
}

func main() {
	var wg sync.WaitGroup

	// List of codes or parameters for each run
	codes := []string{"IN058035001H", "IN058035001I", "IN058035001J","IN058035001J","IN058035001J","IN058035001J","IN058035001J"}

	wg.Add(len(codes))
	for _, code := range codes {
		go runAutomation(code, &wg)
	}

	wg.Wait()
}
