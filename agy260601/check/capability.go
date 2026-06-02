package check

import "errors"

// Capability represents a set of granted abilities for a function.
// In a real implementation this would be parsed from the AST and
// compared against the caller's granted capabilities.

type Capability []string

// Contains checks whether the required capability set is a subset of the
// provided capabilities. It returns an error if any required entry is
// missing.
func (c Capability) Contains(required Capability) error {
    // Build a quick lookup set for the caller's capabilities.
    set := make(map[string]struct{}, len(c))
    for _, abil := range c {
        set[abil] = struct{}{}
    }
    for _, need := range required {
        if _, ok := set[need]; !ok {
            return errors.New("missing required capability: " + need)
        }
    }
    return nil
}

// VerifyCapability ensures that the caller's capabilities (callerCap) are a
// superset of the callee's required capabilities (requiredCap). If the check
// fails, the transpiler should emit an XQL_E3xx error.
func VerifyCapability(callerCap, requiredCap Capability) error {
    return callerCap.Contains(requiredCap)
}
