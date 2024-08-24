package main

import (
	"fmt"
	"strings"
)

func jointStrins(str1, str2 string) string {
	result := strings.Join([]string{str1, str2}, ",")
	fmt.Println("The joint result is:--", result)
	fmt.Printf("The type of result is: %T\n", result)
	return result
}

func countWordOccurance(str string) int {
	count := strings.Count(str, "one")
	return count
}

func main() {
	fmt.Println("This is strtings code file")
	//call joint function
	res := jointStrins("sambit", "Nayak")
	fmt.Println("The result of joint is:--", res)
	//call countWordOccurance function
	count := countWordOccurance("one two one threee one four")
	fmt.Println("The count result is:--", count)

	//check substring contains or not
	str := "one two three"
	check := strings.Contains(str, "one")
	fmt.Println("Check result is:--", check)

	//trim white space from left and right
	trim_str := "     Sambit Nayak   "
	fmt.Println("The trim result is:--", strings.Trim(trim_str, " "))

	//convert to lowercase
	fmt.Println("Convert to Lowercase:--", strings.ToLower("Hello, Sam"))
}
