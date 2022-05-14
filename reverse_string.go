package reverse_string

func ReverseString(input string) (output string) {
	input_len := len(input)
	if input_len > 0 {
		char_list := make([]rune, input_len, input_len)
		for index, value := range input {
			char_list[input_len-1-index] = value
		}
		return string(char_list)
	}
	return output
}
