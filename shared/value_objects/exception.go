package valueobjects

type Exception struct {
	Code     string
	Err      error
	Message  string
	Metadata any
}

func NewException(err error, code, message string, metadata ...any) Exception {
	var exMetadata any

	if len(metadata) > 0 {
		exMetadata = metadata[0]
	}

	return Exception{
		Code:     code,
		Err:      err,
		Message:  message,
		Metadata: exMetadata,
	}
}

func (e Exception) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return ""
}
