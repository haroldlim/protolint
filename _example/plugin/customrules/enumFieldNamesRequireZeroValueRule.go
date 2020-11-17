package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// EnumFieldNamesRequireZeroValueRule verifies that enum must have a zero value.
type EnumFieldNamesRequireZeroValueRule struct{}

// NewEnumFieldNamesRequireZeroValueRule creates a new EnumFieldNamesRequireZeroValueRule.
func NewEnumFieldNamesRequireZeroValueRule() EnumFieldNamesRequireZeroValueRule {
	return EnumFieldNamesRequireZeroValueRule{}
}

// ID returns the ID of this rule.
func (r EnumFieldNamesRequireZeroValueRule) ID() string {
	return "ENUM_FIELD_NAMES_REQUIRE_ZERO_VALUE"
}

// Purpose returns the purpose of this rule.
func (r EnumFieldNamesRequireZeroValueRule) Purpose() string {
	return "Verifies that enum has a zero value."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r EnumFieldNamesRequireZeroValueRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r EnumFieldNamesRequireZeroValueRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &enumFieldNamesRequireZeroValueVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type enumFieldNamesRequireZeroValueVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitEnum checks the enum field.
func (v *enumFieldNamesRequireZeroValueVisitor) VisitEnum(e *parser.Enum) bool {
	foundZero := false
	for _, body := range e.EnumBody {
		switch body.(type) {
		case *parser.EnumField:
			if body.(*parser.EnumField).Number == "0" {
				foundZero = true
			}
		}
	}
	if !foundZero {
		v.AddFailuref(e.Meta.Pos, "Enum %q must have a zero value", e.EnumName)
	}
	return false
}
