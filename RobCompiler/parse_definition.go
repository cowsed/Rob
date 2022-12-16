package main

import "RobCompiler/tokenizer"

type TopASTNode interface {
	Where() tokenizer.SourceRange
}

type ImportDeclaration struct {
	name  string
	where tokenizer.SourceRange
}

func (i *ImportDeclaration) Where() tokenizer.SourceRange {
	return i.where
}

type ModuleDeclaration struct {
	name  string
	where tokenizer.SourceRange
}

func (m *ModuleDeclaration) Where() tokenizer.SourceRange {
	return m.where
}

type CommentNode struct {
	text  string
	where tokenizer.SourceRange
}

func (c *CommentNode) Where() tokenizer.SourceRange {
	return c.where
}
