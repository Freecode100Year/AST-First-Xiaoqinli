package check

// TypesChecker performs static type checking on the AST.
// It validates variable declarations, function signatures, and return types.
// In a full implementation it would traverse the AST nodes defined in ast package.

type TypesChecker struct {}

// NewTypesChecker creates a new TypesChecker.
func NewTypesChecker() *TypesChecker { return &TypesChecker{} }

// Check performs type checking on the provided AST root.
// Returns an error if any type inconsistency is found.
func (c *TypesChecker) Check(root interface{}) error {
    // TODO: implement actual type checking logic.
    // Placeholder to illustrate structure.
    return nil
}

// Effect represents inferred side‑effects of a function.
// Effects are used by the compiler to enforce purity guarantees.

type Effect string

const (
    EffectPure   Effect = "pure"
    EffectIO     Effect = "io"
    EffectFS     Effect = "fs"
    EffectNetwork Effect = "network"
)

// InferEffects walks the AST and infers the effects of each function.
// It returns a map from function name to its inferred Effect set.
func (c *TypesChecker) InferEffects(root interface{}) (map[string][]Effect, error) {
    // TODO: implement effect inference.
    return nil, nil
}
