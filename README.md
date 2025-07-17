# AutomationWithGOLang

  Overview
`lgGOMT.go` is a Go (Golang) automation script using the Playwright framework to interact with a website (`http://www.lg4all.com/pod`). The script is designed to perform browser automation tasks for multiple codes in parallel. Each code triggers a new browser session that automates various webpage interactions, such as navigation, handling dialogs, checking checkboxes, collecting messages, and extracting data.
Features
	•	Parallel Automation: Executes multiple browser automation tasks concurrently using Go’s goroutines.
	•	Headless and Non-Headless Modes: Runs Chromium with options to disable headless mode for visual debugging or automated runs.
	•	Robust Browser Interactions:
	•	Navigates to code-specific URLs.
	•	Handles browser dialogs automatically.
	•	Detects specific page conditions (like “No Record Found”).
	•	Selects checkboxes and attempts to submit forms.
	•	Extracts text from specific elements if available.
Prerequisites
	•	Go 1.18+
	•	Playwright-Go library : go get github.com/playwright-community/playwright-go
		playwright install 
