package ocrworker

const PREPROCESSOR_IDENTITY = "identity"
const PREPROCESSOR_IMGPROC = "img-proc"

type Preprocessor interface {
	preprocess(ocrRequest *OcrRequest) error
}

type IdentityPreprocessor struct {
}

func (i IdentityPreprocessor) preprocess(ocrRequest *OcrRequest) error {
	return nil
}
