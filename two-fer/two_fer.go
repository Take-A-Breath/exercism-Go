package twofer

// Function to return string including a name when provided, otherwise return "you" in its place
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}
