package tree_sitter_rshtml_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_rshtml "github.com/rshtml/rshtml/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_rshtml.Language())
	if language == nil {
		t.Errorf("Error loading RsHtml grammar")
	}
}
