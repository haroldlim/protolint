package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// EnumNamesNoMsgSuffixRule verifies that all enum names are LowerSnakeCase.
type EnumNamesNoMsgSuffixRule struct{}

// NewEnumNamesNoMsgSuffixRule creates a new EnumNamesNoMsgSuffixRule.
func NewEnumNamesNoMsgSuffixRule() EnumNamesNoMsgSuffixRule {
	return EnumNamesNoMsgSuffixRule{}
}

// ID returns the ID of this rule.
func (r EnumNamesNoMsgSuffixRule) ID() string {
	return "ENUM_NAMES_NO_MSG_SUFFIX"
}

// Purpose returns the purpose of this rule.
func (r EnumNamesNoMsgSuffixRule) Purpose() string {
	return "Verifies that all enum names have no Msg suffix."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r EnumNamesNoMsgSuffixRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r EnumNamesNoMsgSuffixRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &enumNamesNoMsgSuffixVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type enumNamesNoMsgSuffixVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitEnum checks the enum field.
func (v *enumNamesNoMsgSuffixVisitor) VisitEnum(e *parser.Enum) bool {
	if strings.HasSuffix(e.EnumName, "Msg") {
		v.AddFailuref(e.Meta.Pos, "Enum name %q must not be suffixed with Msg", e.EnumName)
	}
	return false
}
