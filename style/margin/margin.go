package margin

import (
	"github.com/gopherjs/vecty"
	"strings"
)

type Type string

const (
	Auto    Type = "auto"
	Initial Type = "initial"
	Inherit Type = "inherit"
)

func (t Type) Apply(h *vecty.HTML) {
	vecty.Style("margin", string(t)).Apply(h)
}

type Value []string

func (v Value) Apply(h *vecty.HTML) {
	vecty.Style("margin", strings.Join(v, " ")).Apply(h)
}
