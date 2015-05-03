package goose

type Goose struct {
	config configuration
}

func New(args ...string) Goose {

	return Goose{
		config: GetDefualtConfiguration(args...),
	}
}

func (this Goose) ExtractFromUrl(url string) *Article {
	cc := Crawler{
		config: this.config,
		url:    url,
		helper: NewUrlHelper(url),
	}
	return cc.Crawl()
}

func (this Goose) ExtractFromRawHtml(url string, rawHtml string) *Article {
	cc := Crawler{
		config:  this.config,
		url:     url,
		rawHtml: rawHtml,
		helper:  NewRawHelper(url, rawHtml),
	}
	return cc.Crawl()
}

func (this Goose) ExtractFromRawHtmlWithHelper(url string, rawHtml string, helper Helper) *Article {
	cc := Crawler{
		config:  this.config,
		url:     url,
		rawHtml: rawHtml,
		helper:  helper,
	}
	return cc.Crawl()
}
