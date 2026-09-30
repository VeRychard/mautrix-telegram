// mautrix-telegram - A Matrix-Telegram puppeting bridge.
// Copyright (C) 2026 verychard
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package matrixfmt

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialReplyQuote(t *testing.T) {
	parser := &HTMLParser{}
	ctx := NewContext(context.Background(), nil)
	for _, tc := range []struct {
		name, html, quote string
	}{
		{"quote", `<blockquote data-telegram-partial-reply>egyik érdemenyem 102%</blockquote>nem lehet rossz a mérő? xD`, "egyik érdemenyem 102%"},
		{"leading whitespace", " \n<blockquote data-telegram-partial-reply>a &amp; b<br>c</blockquote>reply", "a & b\nc"},
		{"formatted quote", `<blockquote data-telegram-partial-reply><b>bold</b> text</blockquote>reply`, "bold text"},
		{"normal blockquote", `<blockquote>quote</blockquote>reply`, ""},
		{"not at start", `reply<blockquote data-telegram-partial-reply>quote</blockquote>`, ""},
		{"empty", `<blockquote data-telegram-partial-reply></blockquote>reply`, ""},
		{"no html", `just text`, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.quote, parser.PartialReplyQuote(tc.html, ctx))
		})
	}
	parsed := parser.Parse(`<blockquote data-telegram-partial-reply>quote</blockquote>reply`, ctx)
	assert.Equal(t, "reply", parsed.String.String())
}
