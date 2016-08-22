package mock

type EditorMock struct {
	Receives struct {
		FilePath     string
		InitialValue []byte
	}
	Returns struct {
		EndValue []byte
		Error    error
	}
}

func (j *EditorMock) RunEditor(filePath string, initialValue []byte) (endValue []byte, err error) {
	j.Receives.FilePath = filePath
	j.Receives.InitialValue = initialValue
	return j.Returns.EndValue, j.Returns.Error
}
