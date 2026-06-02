package ast

// Node definitions for the Xiaoqinli AST.
// For simplicity, we represent nodes as generic structures.
// In a full implementation, each node type would have its own struct.

type Node interface{}

type Program struct {
    Statements []Node `json:"statements"`
}
