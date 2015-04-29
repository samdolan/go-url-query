package urlquery

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// Setup suite
type UrlQueryParamsSuite struct {
	suite.Suite
}

func TestUrlQueryParamsSuiteTest(t *testing.T) {
	suite.Run(t, new(UrlQueryParamsSuite))
}

// NewQueryParams tests
func (s *UrlQueryParamsSuite) TestNewQueryParamsSinglVal() {
	p := [][]string{
		{"foo"},
	}

	_, err := NewQueryParams(p)
	s.Error(err)
}

// NewFromQueryStr tests
func (s *UrlQueryParamsSuite) TestNewFromQueryStrEmpty() {
	queryParams, _ := NewFromQueryStr("")
	s.Equal(0, queryParams.Len())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrJustQuestion() {
	queryParams, _ := NewFromQueryStr("?")
	s.Equal(0, queryParams.Len())
	s.Equal("", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrQuestionAndSpace() {
	queryParams, _ := NewFromQueryStr(" ? ")
	s.Equal(0, queryParams.Len())
	s.Equal("", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrSingle() {
	queryParams, _ := NewFromQueryStr("foo=1")
	s.Equal(1, queryParams.Len())
	s.Equal("foo=1", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrMulti() {
	queryParams, _ := NewFromQueryStr("foo=1&baz=2")
	s.Equal(2, queryParams.Len())
	s.Equal("foo=1&baz=2", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrExtraAmpersandEnd() {
	queryParams, _ := NewFromQueryStr("foo=1&baz=2&")
	s.Equal(2, queryParams.Len())
	s.Equal("foo=1&baz=2", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrExtraAmpersandMiddle() {
	queryParams, _ := NewFromQueryStr("foo=1&&baz=2&")
	s.Equal(2, queryParams.Len())
	s.Equal("foo=1&baz=2", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrExtraEqual() {
	queryParams, _ := NewFromQueryStr("foo=1&baz=2")
	s.Equal(2, queryParams.Len())
	s.Equal("foo=1&baz=2", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrDups() {
	queryParams, _ := NewFromQueryStr("foo=1&baz=2&foo=3")
	s.Equal(3, queryParams.Len())
	s.Equal([]string{"1", "3"}, queryParams.GetAll("foo"))
	s.Equal("foo=1&baz=2&foo=3", queryParams.String())
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrKeyWithNoVal() {
	queryParams, err := NewFromQueryStr("foo")
	s.NoError(err)
	s.Equal(1, queryParams.Len())
	s.Equal("", queryParams.Get("foo"))
	s.Equal("foo", queryParams.originalRawQuery)
}

func (s *UrlQueryParamsSuite) TestNewFromQueryStrKeyWithEq_NoVal() {
	queryParams, err := NewFromQueryStr("foo=")
	s.NoError(err)
	s.Equal(1, queryParams.Len())
	s.Equal("", queryParams.Get("foo"))
	s.Equal("foo=", queryParams.String())
}

//
// QueryParams.OriginalRawQuery tests
//
func (s *UrlQueryParamsSuite) TestOriginalRawQuerySet() {
	queryStr := "somerandomstring"
	queryParams, _ := NewFromQueryStr(queryStr)
	s.Equal(queryStr, queryParams.OriginalRawQuery())
}

func (s *UrlQueryParamsSuite) TestOriginalRawQueryUnset() {
	params := [][]string{}
	queryParams, _ := NewQueryParams(params)
	s.Equal("", queryParams.OriginalRawQuery())
}

//
// QueryParams.Encode tests
//
func (s *UrlQueryParamsSuite) TestEncodeEmptyStr() {
	queryParams, _ := NewFromQueryStr("")
	s.Equal("", queryParams.Encode())
}
func (s *UrlQueryParamsSuite) TestEncodeEncoded() {
	queryParams, _ := NewFromQueryStr("foo=bar&meh=yeah")
	s.Equal("foo=bar&meh=yeah", queryParams.Encode())
}

//
// QueryParams.Escape tests
//
func (s *UrlQueryParamsSuite) TestEscapeEmptyStr() {
	queryParams, _ := NewFromQueryStr("")
	s.Equal("", queryParams.Escape())
}

func (s *UrlQueryParamsSuite) TestEscapeEmptyEscaped() {
	queryParams, _ := NewFromQueryStr("foo=%32")
	s.Equal("foo%3D2", queryParams.Escape())
}
