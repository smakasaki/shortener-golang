package validation

func Execute(rules []Rule) []error {
	var errors []error
	for _, rule := range rules {
		if err := rule.Validate(); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
