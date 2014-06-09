package main

import (
    "flag"
    "fmt"
)

func main() {

    // command line flags
    loopOfTimesFlagPtr := flag.Int(
        "loop",
        5,
        "Tell the program how many times to repeat.")

    userStringPtr := flag.String("s", "Some string...", "a string to repeat")

    flag.Parse()

    // testString := "Test"
    print1(userStringPtr, loopOfTimesFlagPtr)

}

func print1(userString_p *string, numOfTimes_p *int) {
    userString := *userString_p
    numOfTimes := *numOfTimes_p

    for i := 0; i < numOfTimes; i++ {
        output := fmt.Sprintf("%d: %s", i, userString)
        fmt.Println(output)
    }
}
