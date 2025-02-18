package intfc

type Factory interface {
	FromString(entry string) (e IMarshable, err error)
	GetTypeName() string
}
