package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// Proto2FieldsAvoidRequiredRule verifies that all fields should avoid required for proto2.
// See https://developers.google.com/protocol-buffers/docs/style#things-to-avoid
type Proto2FieldsAvoidRequiredRule struct {
}

// NewProto2FieldsAvoidRequiredRule creates a new Proto2FieldsAvoidRequiredRule.
func NewProto2FieldsAvoidRequiredRule() Proto2FieldsAvoidRequiredRule {
	return Proto2FieldsAvoidRequiredRule{}
}

// ID returns the ID of this rule.
func (r Proto2FieldsAvoidRequiredRule) ID() string {
	return "PROTO2_FIELDS_AVOID_REQUIRED"
}

// Purpose returns the purpose of this rule.
func (r Proto2FieldsAvoidRequiredRule) Purpose() string {
	return "Verifies that all fields should avoid required for proto2."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r Proto2FieldsAvoidRequiredRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r Proto2FieldsAvoidRequiredRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &proto2FieldsAvoidRequiredVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type proto2FieldsAvoidRequiredVisitor struct {
	*visitor.BaseAddVisitor
	isProto2 bool
}

// VisitSyntax checks the syntax.
func (v *proto2FieldsAvoidRequiredVisitor) VisitSyntax(s *parser.Syntax) bool {
	v.isProto2 = s.ProtobufVersion == "proto2"
	return false
}

// VisitField checks the field.
func (v *proto2FieldsAvoidRequiredVisitor) VisitField(field *parser.Field) bool {
	if v.isProto2 && field.IsRequired {
		v.AddFailuref(field.Meta.Pos, `Field %q should avoid required for proto2`, field.FieldName)
	}
	return false
}
