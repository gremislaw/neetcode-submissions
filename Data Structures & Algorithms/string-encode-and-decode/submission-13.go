type Solution struct{
	DecodeChar string
}

func (s *Solution) Encode(strs []string) string {
	s.DecodeChar = "dazhy"
	var res strings.Builder
	for i, str := range strs {
		res.WriteString(str)
		if i != len(strs) - 1 {
			res.WriteString(s.DecodeChar)
		}
	}
	if len(strs) == 0 {
		return "semenoid"
	}
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "semenoid" {
		return []string{}
	}
	res := strings.Split(encoded, s.DecodeChar)
	return res
}
