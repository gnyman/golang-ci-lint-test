package hello

import "fmt"

func Person(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func Bad() {
	if true {
		fmt.Println("Live code")
	} else {
		fmt.Println("Dead code")
	}
}
