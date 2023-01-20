package fetchemail

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func Task_creator(token string) string {
	// here we can define a role variable
	var role1 string
	// here we can use for loop for split the token by dot(.) and i will store index value and part will store the value
	for index, part := range strings.Split(token, ".") {
		// here we can decode the string.
		decoded, err := base64.RawURLEncoding.DecodeString(part)
		if err != nil {
			panic(err)
		}
		// here if i == 1 then it work on payload
		if index != 1 {
			continue // i == 1 is the payload
		}
		// here we can declare a M variable were we can store the decode value.
		var m map[string]interface{}
		if err := json.Unmarshal(decoded, &m); err != nil {
			fmt.Println("json decoding failed:", err)
			continue
		}

		// here we can save the value in role variable for return the string value
		role1 = m["email"].(string)

	}
	return role1
}
