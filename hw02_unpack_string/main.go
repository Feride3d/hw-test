package main

import "github.com/Feride3d/hw-test/hw02_unpack_string/hw02unpackstring"

func main() {
	
	yourString := []string{"a4bc2d5e", "abcd", "3abc" "45", "aaa10b", "", "d\n5abc"}
    	for _, string := range yourString {
    		result, err := hw02unpackstring.Unpack(string)
    		if err != nil {
    			fmt.Printf("%s => %s (not unpacked) : %s\n", string, result, err)
    		} else {
    			fmt.Printf("%s => %s (unpacked)\n", string, result)
    		}
    
    	}
}