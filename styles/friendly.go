package styles

import (
	"github.com/eastym/chroma/v2"
)

// Friendly style.
var Friendly = Register(chroma.MustNewStyle("friendly", chroma.StyleEntries{
	chroma.TextWhitespace:        "#bbbbbb",
	chroma.Comment:               "italic #60a0b0",
	chroma.CommentPreproc:        "noitalic #007020",
	chroma.CommentSpecial:        "noitalic bg:#fff0f0",
	chroma.Keyword:               "bold #007020",
	chroma.KeywordPseudo:         "nobold",
	chroma.KeywordType:           "nobold #902000",
	chroma.Operator:              "#666666",
	chroma.OperatorWord:          "bold #007020",
	chroma.NameBuiltin:           "#007020",
	chroma.NameFunction:          "#06287e",
	chroma.NameClass:             "bold #0e84b5",
	chroma.NameNamespace:         "bold #0e84b5",
	chroma.NameException:         "#007020",
	chroma.NameVariable:          "#bb60d5",
	chroma.NameConstant:          "#60add5",
	chroma.NameLabel:             "bold #002070",
	chroma.NameEntity:            "bold #d55537",
	chroma.NameAttribute:         "#4070a0",
	chroma.NameTag:               "bold #062873",
	chroma.NameDecorator:         "bold #555555",
	chroma.LiteralString:         "#4070a0",
	chroma.LiteralStringDoc:      "italic",
	chroma.LiteralStringInterpol: "#70a0d0",
	chroma.LiteralStringEscape:   "bold #4070a0",
	chroma.LiteralStringRegex:    "#235388",
	chroma.LiteralStringSymbol:   "#517918",
	chroma.LiteralStringOther:    "#c65d09",
	chroma.LiteralNumber:         "#40a070",
	chroma.GenericHeading:        "bold #000080",
	chroma.GenericSubheading:     "bold #800080",
	chroma.GenericDeleted:        "#A00000",
	chroma.GenericInserted:       "#00A000",
	chroma.GenericError:          "#FF0000",
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold #c65d09",
	chroma.GenericOutput:         "#888",
	chroma.GenericTraceback:      "#04D",
	chroma.GenericUnderline:      "underline",
	chroma.Error:                 "border:#FF0000",
	chroma.Background:            " bg:#f0f0f0",
}))
