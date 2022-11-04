package styles

import (
	"github.com/eastym/chroma/v2"
)

// Average style.
var Average = Register(chroma.MustNewStyle("average", chroma.StyleEntries{
	chroma.Comment:                  "#757575",
	chroma.CommentHashbang:          "#757575",
	chroma.CommentMultiline:         "#757575",
	chroma.CommentPreproc:           "#757575",
	chroma.CommentSingle:            "#757575",
	chroma.CommentSpecial:           "#757575",
	chroma.Generic:                  "#757575",
	chroma.GenericDeleted:           "#ec0000",
	chroma.GenericEmph:              "#757575 underline",
	chroma.GenericError:             "#ec0000",
	chroma.GenericHeading:           "#757575 bold",
	chroma.GenericInserted:          "#757575 bold",
	chroma.GenericOutput:            "#757575",
	chroma.GenericPrompt:            "#757575",
	chroma.GenericStrong:            "#757575 italic",
	chroma.GenericSubheading:        "#757575 bold",
	chroma.GenericTraceback:         "#757575",
	chroma.GenericUnderline:         "underline",
	chroma.Error:                    "#ec0000",
	chroma.Keyword:                  "#ec0000",
	chroma.KeywordConstant:          "#ec0000",
	chroma.KeywordDeclaration:       "#ec0000",
	chroma.KeywordNamespace:         "#ec0000",
	chroma.KeywordPseudo:            "#ec0000",
	chroma.KeywordReserved:          "#ec0000",
	chroma.KeywordType:              "#5f5fff",
	chroma.Literal:                  "#757575",
	chroma.LiteralDate:              "#757575",
	chroma.Name:                     "#757575",
	chroma.NameAttribute:            "#5f5fff",
	chroma.NameBuiltin:              "#ec0000",
	chroma.NameBuiltinPseudo:        "#757575",
	chroma.NameClass:                "#5f5fff",
	chroma.NameConstant:             "#008900",
	chroma.NameDecorator:            "#008900",
	chroma.NameEntity:               "#757575",
	chroma.NameException:            "#757575",
	chroma.NameFunction:             "#5f5fff",
	chroma.NameLabel:                "#ec0000",
	chroma.NameNamespace:            "#757575",
	chroma.NameOther:                "#757575",
	chroma.NameTag:                  "#ec0000",
	chroma.NameVariable:             "#ec0000",
	chroma.NameVariableClass:        "#ec0000",
	chroma.NameVariableGlobal:       "#ec0000",
	chroma.NameVariableInstance:     "#ec0000",
	chroma.LiteralNumber:            "#008900",
	chroma.LiteralNumberBin:         "#008900",
	chroma.LiteralNumberFloat:       "#008900",
	chroma.LiteralNumberHex:         "#008900",
	chroma.LiteralNumberInteger:     "#008900",
	chroma.LiteralNumberIntegerLong: "#008900",
	chroma.LiteralNumberOct:         "#008900",
	chroma.Operator:                 "#ec0000",
	chroma.OperatorWord:             "#ec0000",
	chroma.Other:                    "#757575",
	chroma.Punctuation:              "#757575",
	chroma.LiteralString:            "#008900",
	chroma.LiteralStringBacktick:    "#008900",
	chroma.LiteralStringChar:        "#008900",
	chroma.LiteralStringDoc:         "#008900",
	chroma.LiteralStringDouble:      "#008900",
	chroma.LiteralStringEscape:      "#008900",
	chroma.LiteralStringHeredoc:     "#008900",
	chroma.LiteralStringInterpol:    "#008900",
	chroma.LiteralStringOther:       "#008900",
	chroma.LiteralStringRegex:       "#008900",
	chroma.LiteralStringSingle:      "#008900",
	chroma.LiteralStringSymbol:      "#008900",
	chroma.Text:                     "#757575",
	chroma.TextWhitespace:           "#757575",
	chroma.Background:               " bg:#000000",
}))
