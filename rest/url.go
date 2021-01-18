package rest

import "net/url"

type URLBuilder struct {
	u     *url.URL
	query url.Values
}

func ParseURL(rawUrl string) *URLBuilder {
	ub := &URLBuilder{}
	ub.u, _ = url.Parse(rawUrl)
	ub.query = ub.u.Query()
	return ub
}

func (builder *URLBuilder) AddQuery(name, value string) *URLBuilder {
	builder.query.Add(name, value)
	return builder
}

func (builder *URLBuilder) AddQueries(queries map[string]string) *URLBuilder {
	for name, value := range queries {
		builder.AddQuery(name, value)
	}
	return builder
}

func (builder *URLBuilder) GetQuery() url.Values {
	return builder.query
}

func (builder *URLBuilder) GetURL() *url.URL {
	return builder.u
}

func (builder *URLBuilder) Build() *url.URL {
	builder.u.RawQuery = builder.query.Encode()
	return builder.u
}

func (builder *URLBuilder) BuildStr() string {
	return builder.Build().String()
}
