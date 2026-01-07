package handlehashmap

import "fmt"

func PrintFullHashMap(hashMap map[string]uint){
	for key, _ := range hashMap{
		fmt.Println(key)
	}
}