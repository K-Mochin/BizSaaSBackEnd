package bizlogics

import "fmt"

func grpcreciever(message string) bool {
	fmt.Println("grpcReciever received:", message)
	return true
}
