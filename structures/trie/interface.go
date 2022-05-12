package trie

type ITrie interface {
	Insert(key string, value interface{})
	Search(key string) (*Node, bool)
	/*Keys() []string
	Values() []interface{}
	PrefixSearch(key string) []string
	PrefixSearchValues(key string) []interface{}
	PrefixSearchKeys(key string) []string*/
}
