package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/strs"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// OneOfNamesLowerSnakeCaseRule verifies that all field names are underscore_separated_names.
// See https://developers.google.com/protocol-buffers/docs/style#message-and-field-names.
type OneOfNamesLowerSnakeCaseRule struct{}

// NewOneOfNamesLowerSnakeCaseRule creates a new OneOfNamesLowerSnakeCaseRule.
func NewOneOfNamesLowerSnakeCaseRule() OneOfNamesLowerSnakeCaseRule {
	return OneOfNamesLowerSnakeCaseRule{}
}

// ID returns the ID of this rule.
func (r OneOfNamesLowerSnakeCaseRule) ID() string {
	return "ONEOF_NAMES_LOWER_SNAKE_CASE"
}

// Purpose returns the purpose of this rule.
func (r OneOfNamesLowerSnakeCaseRule) Purpose() string {
	return "Verifies that all oneof names are underscore_separated_names."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r OneOfNamesLowerSnakeCaseRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r OneOfNamesLowerSnakeCaseRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &oneOfNamesLowerSnakeCaseVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type oneOfNamesLowerSnakeCaseVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitOneof checks the Oneof.
func (v *oneOfNamesLowerSnakeCaseVisitor) VisitOneof(oneof *parser.Oneof) bool {
	if !strs.IsLowerSnakeCase(oneof.OneofName) {
		v.AddFailuref(oneof.Meta.Pos, "Oneof name %q must be underscore_separated_names", oneof.OneofName)
	}
	return false
}
