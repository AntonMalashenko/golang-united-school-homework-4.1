package reverse_string

func ReverseString(input string) (output string) {
	input_len := len(input)
	if input_len > 0 {
		temp_char_list := make([]rune, 0)
		for _, value := range input {
			temp_char_list = append(temp_char_list, value)
		}

		temp_len := len(temp_char_list)
		output_char_list := make([]rune, temp_len)
		for idx, val := range temp_char_list {
			output_char_list[temp_len-1-idx] = val
		}
		return string(output_char_list)
	}
	return output
}
