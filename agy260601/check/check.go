package check

import "errors"

// RunAll performs all static checks on the AST.
// It invokes type checking, effect inference, and capability verification.
// In this minimal implementation it simply runs the type checker and returns any error.
func RunAll(ast interface{}) error {
    // Type checking (placeholder implementation)
    tc := NewTypesChecker()
    if err := tc.Check(ast); err != nil {
        return err
    }
    // Effect inference (ignored for now)
    // Capability verification would be performed here in a full implementation.
    return nil
}
