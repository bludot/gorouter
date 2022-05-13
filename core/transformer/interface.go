package transformer

type Transformer interface {
	ToJson(data interface{}) (string, error)
}
