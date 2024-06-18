func Spinner() {
	// Create a new spinner with a specific character set and delay
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)

	// Set a custom prefix and suffix for the spinner
	s.Prefix = "Loading: "
	s.Suffix = " Please wait..."

	// Change spinner color to yellow, bold, with a black background
	s.Color("yellow", "bold", "bgBlack")

	// Start the spinner animation
	s.Start()

	// Launch a goroutine to dynamically update the suffix while the spinner is running
	go func() {
		for i := 0; i < 5; i++ {
			// Update the suffix with a progress message
			s.Suffix = fmt.Sprintf(" Processing %d/5", i+1)
			// Sleep for 1 second before updating again
			time.Sleep(1 * time.Second)
		}
	}()

	// Simulate some work with a sleep of 5 seconds
	time.Sleep(5 * time.Second)

	// Stop the spinner animation
	s.Stop()

	// Print a message indicating completion
	fmt.Println("Done!")
}
