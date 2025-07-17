# README for `lgGOMT.go`

## Overview

`lgGOMT.go` is a Go (Golang) automation script using the Playwright framework to interact with a website (`http://www.lg4all.com/pod`). The script is designed to perform browser automation tasks for multiple codes in parallel. Each code triggers a new browser session that automates various webpage interactions, such as navigation, handling dialogs, checking checkboxes, collecting messages, and extracting data.

## Features

- **Parallel Automation:**  
  Executes multiple browser automation tasks concurrently using Go's goroutines.

- **Headless and Non-Headless Modes:**  
  Runs Chromium with options to disable headless mode for visual debugging or automated runs.

- **Robust Browser Interactions:**  
  - Navigates to code-specific URLs.
  - Handles browser dialogs automatically.
  - Detects specific page conditions (like "No Record Found").
  - Selects checkboxes and attempts to submit forms.
  - Extracts text from specific elements if available.

## Prerequisites

- **Go 1.18+**
- **Playwright-Go library**  
  Install with:  
  ```bash
  go get github.com/playwright-community/playwright-go
  ```
- **Playwright Browsers**  
  Install Playwright browsers if not already present:  
  ```bash
  playwright install
  ```
- **Chromium Browser** (handled by Playwright install)
- **Internet Connection** (for website access)

## How It Works

1. **Code List:**  
   The script defines a list of code strings. Each code triggers a separate browser automation routine.

2. **Parallel Execution:**  
   For every code in the list, a goroutine is launched via `runAutomation`, managed by a `sync.WaitGroup` for synchronization.

3. **Browser Automation:**  
   Each goroutine:
   - Launches a new Chromium browser with specific flags for resource efficiency and sandboxing.
   - Navigates to the URL built from the code.
   - Handles unexpected dialogs automatically.
   - Checks for a "No Record Found" message and exits early if found.
   - Checks/unchecks checkboxes as needed and attempts to submit results.
   - Extracts and logs the "No Of Bid" value from a specific element if present.

4. **Graceful Shutdown:**  
   Browsers and Playwright are stopped using deferred functions to ensure resources are always cleaned.

## Usage

1. **Edit `codes` Array:**  
   Replace the hardcoded codes with your own list of identifiers to automate against.

2. **Compile and Run:**
   ```bash
   go mod tidy               # Install dependencies
   go run lgGOMT.go
   ```
   Or compile and run separately:
   ```bash
   go build -o lgGOMT
   ./lgGOMT
   ```

3. **View Output:**  
   Log outputs will be printed to the console, showing action status and any data extracted.

## Key Configuration Points

- **Headless Mode:**  
  To run browsers without UI for speed, set `Headless: playwright.Bool(true)` in the browser launch options.

- **Concurrency:**  
  The script launches one goroutine per code in the list. Adjust the `codes` list to match system resource constraints.

- **Error Handling:**  
  The script uses fatal logs on most errors, which means a single error in a goroutine can terminate the entire program. Consider refactoring to handle errors gracefully if needed for production use.

## Example Output

```
Alert text: <Dialog message>
NoOfBid element not found on the page.
Submit button clicked.
No Of Bid Text: <value>
No Record Found message. Exiting.
```

## Troubleshooting

- **Missing Dependencies:**  
  Ensure Playwright and the Playwright-Go bindings are installed.
- **Resource Usage:**  
  Launching many browsers in parallel may require tuning based on your server's CPU/memory capacity.
- **Firewall/Network Issues:**  
  Make sure the automated browser can access external web resources.

## Customization

To use the script with different URLs, element selectors, or workflows, edit the relevant sections in `runAutomation`.

## Performance Tips

- **Adjust Concurrency:** Monitor system resources and adjust the number of concurrent browser instances based on your server's capacity.
- **Use Headless Mode:** For production runs, enable headless mode to reduce resource consumption.
- **Optimize Wait Times:** Fine-tune wait times and timeouts for better performance.

## License

This project is provided as-is for educational and automation purposes.

**Note:**  
Always review automated browser usage policies and site terms to ensure compliance before running scripts against third-party websites.
