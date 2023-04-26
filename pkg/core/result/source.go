package result

// Source holds source execution result
type Source struct {
	// Name holds the target name
	Name string
	/*
		Result holds the source result, accepted values must be one:
			* "SUCCESS"
			* "FAILURE"
			* "ATTENTION"
			* "SKIPPED"
	*/
	Result string
	// Information stores the information detected by the source execution such as a version
	Information string
	// Description stores the source execution description
	Description string
}