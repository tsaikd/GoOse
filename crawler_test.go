package goose

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func ReadRawHtml(a Article) string {
	path := fmt.Sprintf("sites/%s.html", a.Domain)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("cannot read %q", path))
	}

	return string(file)
}

func ValidateArticle(a Article) error {
	g := New()
	expected := g.ExtractFromRawHtml(a.FinalUrl, ReadRawHtml(a))

	if expected.Title != a.Title {
		return fmt.Errorf("article title does not match. Got %q", expected.Title)
	}

	if expected.MetaDescription != a.MetaDescription {
		return fmt.Errorf("article metaDescription does not match. Got %q", expected.MetaDescription)
	}

	if !strings.Contains(expected.CleanedText, a.CleanedText) {
		return fmt.Errorf("article cleanedText does not contains %q", a.CleanedText)
	}

	if expected.MetaKeywords != a.MetaKeywords {
		return fmt.Errorf("article keywords does not match. Got %q", expected.MetaKeywords)
	}
	if expected.CanonicalLink != a.CanonicalLink {
		return fmt.Errorf("article CanonicalLink does not match. Got %q", expected.CanonicalLink)
	}

	if expected.TopImage != a.TopImage {
		return fmt.Errorf("article topImage does not match. Got %q", expected.TopImage)
	}

	return nil
}

func Test_GloboEsporteParse(t *testing.T) {
	article := Article{
		Domain:          "globoesporte.globo.com",
		Title:           "Rodrigo Caio treina até nas férias e tenta acelerar retorno aos gramados",
		MetaDescription: "Rodrigo Caio treina na esteira durante as férias em Dracena-SP (Foto: Divulgação)Rodrigo Caio quer ganhar tempo na recuperação da lesão que sofreu no joelho esquerdo. Apesar de ter sido liberado pelo departamento médico do São Paulo para as férias, o ...",
		CleanedText:     "Comissão técnica planeja volta dele para o fim de fevereiro ou início de março",
		MetaKeywords:    "notícias, notícia, presidente prudente região",
		CanonicalLink:   "http://globoesporte.globo.com/sp/presidente-prudente-regiao/noticia/2014/12/rodrigo-caio-treina-ate-nas-ferias-e-tenta-acelerar-retorno-aos-gramados.html",
		TopImage:        "http://s.glbimg.com/es/ge/f/original/2014/12/26/10863872_894379987249341_2406060334390226774_o.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

func Test_EditionCnnParse(t *testing.T) {
	article := Article{
		Domain:          "edition.cnn.com",
		Title:           "What if you could make anything you wanted?",
		MetaDescription: "Massimo Banzi's pocket-sized open-source circuit board has become a key building block in the creation of a huge variety of innovative devices.",
		CleanedText:     "In the 20th century, getting your child a toy car meant a trip to a shopping mall.",
		MetaKeywords:    "",
		CanonicalLink:   "http://www.cnn.com/2012/07/08/opinion/banzi-ted-open-source/index.html",
		TopImage:        "http://i2.cdn.turner.com/cnn/dam/assets/120706022111-ted-cnn-ideas-massimo-banzi-00003302-story-top.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

func Test_BbcParse(t *testing.T) {
	article := Article{
		Domain:          "bbc.com",
		Title:           "Crunch talks on new Greek bailout",
		MetaDescription: "German and Greek finance ministers meet IMF and Eurogroup chiefs ahead of a crucial finance ministers' meeting on Greece's bailout request.",
		CleanedText:     "Mr Tsipras won elections in late January on a platform of rejecting the austerity measures tied to the bailout.",
		MetaKeywords:    "keywords, added, to, test, case insensitive",
		CanonicalLink:   "http://www.bbc.com/news/business-31545115",
		TopImage:        "http://news.bbcimg.co.uk/media/images/81120000/jpg/_81120901_81120501.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

// multiple og:image, according to http://ogp.me/, the first one should be preferred
func Test_LindorffParse(t *testing.T) {
	article := Article{
		Domain:          "profit.lindorff.fi",
		Title:           "Lindorff24.fi muuttaa maksujen hoidon mobiiliksi - Lindorff Profit",
		MetaDescription: "Lindorffin verkkopalvelu kuluttajille tunnetaan nyt nimellä Lindorff24.fi. Uusien ominaisuuksien lisäksi palvelu on käytettävissä tietokoneen lisäksi älypuhelimella ja tabletilla. Verkon itseasioinnin uskotaan kasvavan lähivuosina merkittävästi nykyisestä.",
		CleanedText:     "",
		MetaKeywords:    "",
		CanonicalLink:   "http://profit.lindorff.fi/lindorff24-fi-muuttaa-maksujen-hoidon-mobiiliksi/",
		TopImage:        "http://profit.lindorff.fi/wp-content/uploads/2015/02/Iso_Lindorff24_2_600x2501.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

// Facebook photo
func Test_FacebookParse(t *testing.T) {
	article := Article{
		Domain:          "facebook.com",
		Title:           "Facebook - Facebook's Photos",
		MetaDescription: "Stay connected with all of your groups with the new Facebook Groups app. Learn more: http://www.facebookgroups.com",
		CleanedText:     "",
		MetaKeywords:    "",
		CanonicalLink:   "https://www.facebook.com/facebook/photos/a.376995711728.190761.20531316728/10153398878696729/",
		TopImage:        "https://fbcdn-sphotos-g-a.akamaihd.net/hphotos-ak-xpa1/v/t1.0-9/p180x540/10408016_10153398878696729_8237363642999953356_n.png?oh=c6ae71220447f363ec41ea54c38341e1&oe=55B6D827&__gda__=1436749528_5c72e92a5105c1cc6df97163a64e72ce",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}
