package docs_demo_test

import docs_demo "example.com/docs-demo"

func Example() {
	docs_demo.PrintKb(2 * docs_demo.BytesPerKb)
	// Output:
	// 2 kB
}

func ExamplePrintKb() {
	docs_demo.PrintKb(5 * docs_demo.BytesPerKb)
	// Output:
	// 5 kB
}

// Prints one megabty on kilobytes to the terminal.
func ExamplePrintKb_megabyte() {
	docs_demo.PrintKb(docs_demo.BytesPerKb * docs_demo.BytesPerKb)
	// Output:
	// 1024 kB
}
