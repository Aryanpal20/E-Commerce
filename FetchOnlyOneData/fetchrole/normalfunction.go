package fetchrole

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// here we can use normal function and get the token by login
func Is_manager(token string) string {
	// here we can define a role variable
	var role string
	// here we can use for loop for split the token by dot(.) and i will store index value and part will store the value
	for i, part := range strings.Split(token, ".") {
		// here we can decode the string.
		decoded, err := base64.RawURLEncoding.DecodeString(part)
		if err != nil {
			panic(err)
		}
		// here if i == 1 then it work on payload
		if i != 1 {
			continue // i == 1 is the payload
		}
		// here we can declare a M variable were we can store the decode value.
		var m map[string]interface{}
		if err := json.Unmarshal(decoded, &m); err != nil {
			fmt.Println("json decoding failed:", err)
			continue
		}
		// if email, ok := m["email"]; ok {
		// 	fmt.Println("email:", email)
		// }
		// fmt.Println(m)
		// fmt.Println(m["role"])
		// here we can save the value in role variable for return the string value
		role = m["role"].(string)

	}
	return role

}
