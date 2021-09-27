package fileid

type GenerateId interface {
	Generate() (string, error)
}

func Factory() GenerateId {
	return UUID{}
}