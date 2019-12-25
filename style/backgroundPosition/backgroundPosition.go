package backgroundPosition

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/vectyUI/internal"
	"strings"
)

type Type string

const (
	Left   Type = "left"
	Center Type = "center"
	Right  Type = "right"
	Top    Type = "top"
	Bottom Type = "bottom"
)

func Len(length interface{}) Type {
	return Type(internal.Stringify(length, "px"))
}

func Percent(percent interface{}) Type {
	return Type(internal.Stringify(percent, "%"))
}

type Value []Type

func (v Value) Apply(h *vecty.HTML) {
	sb := strings.Builder{}
	for i, t := range v {
		sb.WriteString(string(t))
		if i < len(v)-1 {
			sb.WriteString(",")
		}
	}
	vecty.Style("background-position", sb.String()).Apply(h)
}