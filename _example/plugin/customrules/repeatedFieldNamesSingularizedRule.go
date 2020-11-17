package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/strs"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// RepeatedFieldNamesSingularizedRule verifies that repeated field names are pluralized names.
// See https://developers.google.com/protocol-buffers/docs/style#repeated-fields.
type RepeatedFieldNamesSingularizedRule struct {
}

// NewRepeatedFieldNamesSingularizedRule creates a new RepeatedFieldNamesSingularizedRule.
func NewRepeatedFieldNamesSingularizedRule() RepeatedFieldNamesSingularizedRule {
	return RepeatedFieldNamesSingularizedRule{}
}

// ID returns the ID of this rule.
func (r RepeatedFieldNamesSingularizedRule) ID() string {
	return "REPEATED_FIELD_NAMES_SINGULARIZED"
}

// Purpose returns the purpose of this rule.
func (r RepeatedFieldNamesSingularizedRule) Purpose() string {
	return "Verifies that repeated field names are singularized names."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r RepeatedFieldNamesSingularizedRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r RepeatedFieldNamesSingularizedRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	c := strs.NewPluralizeClient()

	v := &repeatedFieldNamesSingularizedCaseVisitor{
		BaseAddVisitor:  visitor.NewBaseAddVisitor(r.ID()),
		pluralizeClient: c,
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type repeatedFieldNamesSingularizedCaseVisitor struct {
	*visitor.BaseAddVisitor
	pluralizeClient *strs.PluralizeClient
}

// VisitField checks the field.
func (v *repeatedFieldNamesSingularizedCaseVisitor) VisitField(field *parser.Field) bool {
	got := field.FieldName
	want := v.pluralizeClient.ToPlural(got)
	if field.IsRepeated && got == want {
		v.AddFailuref(field.Meta.Pos, "Repeated field name %q must be singularized name", got)
	}
	return false
}

// VisitGroupField checks the group field.
func (v *repeatedFieldNamesSingularizedCaseVisitor) VisitGroupField(field *parser.GroupField) bool {
	got := field.GroupName
	want := v.pluralizeClient.ToPlural(got)
	if field.IsRepeated && got == want {
		v.AddFailuref(field.Meta.Pos, "Repeated group name %q must be singularized name", got)
	}
	return true
}
