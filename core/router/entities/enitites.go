package entities

type QueryParams map[string]string

type RouteParams map[string]string

func (r RouteParams) Get(key string) string {
	return r[key]
}
