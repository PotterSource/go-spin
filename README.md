# go-spin
A Go package that presents a stylized CLI spinner for long-running processes

## Usage
```go
func main() {
	s, _ := spinner.NewSpinner(
		"dots",
		color.Orange(),
	)
	s.Start("getting something...")

	// Simulate a long-running task in a goroutine
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second) // Simulate task duration
		done <- true
	}()

	// Handle SIGINT and SIGTERM to stop the spinner gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Wait for the task to complete or a signal
	select {
	case <-done:
		s.Stop("Success!")
	case <-sig:
		s.Stop("Interrupted!")
	}
}
```

## Spinner Types

- `bar`
- `dots`
- `dots2`
- `lines`

## Colors

- `color.Red()`
- `color.Green()`
- `color.Blue()`
- `color.Yellow()`
- `color.Magenta()`
- `color.Cyan()`
- `color.Orange()`
- `color.White()`
- `color.Black()`

### Custom Colors

Create your own RGB Color

```go
lightBlue, _ := color.NewColor(173, 216, 230)
```