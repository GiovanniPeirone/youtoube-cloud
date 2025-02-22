package main 

import (
    "fmt"
    "strconv"
    "strings"

    "image/color"
    "os"

    "github.com/fogleman/gg"
    

)



const (
    width = 800
    height = 600
    frames = 1
    fameRate = 30
)




func textToBinary(text string) string {
    
    binaryString := ""

    for _, char := range text {
        binaryString += fmt.Sprintf("%08b ",  char)
    }
    

    return binaryString
}

func binaryToText(binary string) (string, error) {
    
    bits := strings.Fields(binary)

    var text string

    for _, b := range bits {
        num, err := strconv.ParseInt(b, 2, 8)

        if err != nil {
            return "", err         
        }
        text += string(rune(num))
    }

    return text, nil
}

func writeVideo() {
        
}

func ReadVideo() {

}

func main() {
    
    text := "Hello World"
    
    binary := textToBinary(text)
   
    toText, err := binaryToText(binary)

    fmt.Println(binary)


    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(toText)
}
