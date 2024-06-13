//go:build no_runtime_type_checking

package occecrpattern

// Building without runtime type checking enabled, so all the below just return nil

func validateOCCEcrPattern_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_OCCEcrPattern) validateSetEcrParameters(val awsecr.IRepository) error {
	return nil
}

func (j *jsiiProxy_OCCEcrPattern) validateSetEcrArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OCCEcrPattern) validateSetEcrImageNameParameters(val *string) error {
	return nil
}

func validateNewOCCEcrPatternParameters(scope constructs.Construct, id *string, props IEcrProps) error {
	return nil
}

