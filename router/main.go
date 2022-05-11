type RouterService struct {
	routes []Route
	cached map[string]*Route
}

func NewRouter() Router {

}