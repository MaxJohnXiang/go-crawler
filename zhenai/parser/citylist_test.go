package parser

import (
	"testing"

	"io/ioutil"
)


func TestParserCityList(t *testing.T) {
	//contents , err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents , err := ioutil.ReadFile(
		"citylist_test_data.html")

	if err != nil {
		panic("error")
	}

	result := ParserCityList(contents)

	expectedUrls := []string {
		"http://www.zhenai.com/zhenghun/aba","http://www.zhenai.com/zhenghun/akesu","http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCited := []string {
		"阿坝","阿克苏","阿拉善盟",
	}
	for i, url  := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url #%d :%s; but " + "was %s",i, url, result.Requests[i].Url)
		}
	}

	for i, city  := range expectedCited {
		if result.Items[i].(string)!= city {
			t.Errorf("excepted url #%d :%s; but " + "was %s",i, city, result.Items[i].(string))
		}
	}

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d" + "requests; but had %d",
			resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d" + "Items; but had %d",
			resultSize, len(result.Items))
	}
}
