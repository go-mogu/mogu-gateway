package model

var (
	Routes []*Route
)

type Route struct {
	Id         string   `json:"id"`
	Uri        string   `json:"uri"`        //uri TODO 先支持lb
	Predicates []string `json:"predicates"` //路径配置规则 TODO 先支持Path
	Filters    []string `json:"filters"`    //
}
