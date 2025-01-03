/**
 * Created by Goland
 * @file   html_test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2025/1/3 14:56
 * @desc   html_test.go
 */

package html

import "testing"

func TestHtml2JS(t *testing.T) {
	htm := "<div id=\"button\" class=\"btn\">Click me</div>\n\r"
	js := string(Html2JS([]byte(htm)))
	jsR := `<div id=\"button\" class=\"btn\">Click me</div>\n`
	if js != jsR {
		t.Errorf("Html2JS:\n Expect => %s\n Got => %s\n", jsR, js)
	}
}

func BenchmarkHtml2JS(b *testing.B) {
	htm := "<div id=\"button\" class=\"btn\">Click me</div>\n\r"
	for i := 0; i < b.N; i++ {
		Html2JS([]byte(htm))
	}
}
