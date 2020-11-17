package customrules

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"

	"github.com/yoheimuta/protolint/linter/report"
	"github.com/yoheimuta/protolint/linter/visitor"
)

// MessageNamesSuffixedWithMsgRule verifies that all message names are CamelCase (with an initial capital).
// See https://developers.google.com/protocol-buffers/docs/style#message-and-field-names.
type MessageNamesSuffixedWithMsgRule struct{}

// NewMessageNamesSuffixedWithMsgRule creates a new MessageNamesSuffixedWithMsgRule.
func NewMessageNamesSuffixedWithMsgRule() MessageNamesSuffixedWithMsgRule {
	return MessageNamesSuffixedWithMsgRule{}
}

// ID returns the ID of this rule.
func (r MessageNamesSuffixedWithMsgRule) ID() string {
	return "MESSAGE_NAMES_SUFFIXED_WITH_MSG"
}

// Purpose returns the purpose of this rule.
func (r MessageNamesSuffixedWithMsgRule) Purpose() string {
	return "Verifies that all message names are suffixed with Msg."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r MessageNamesSuffixedWithMsgRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r MessageNamesSuffixedWithMsgRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &messageNamesSuffixedWithMsgVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(r.ID()),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type messageNamesSuffixedWithMsgVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitMessage checks the message.
func (v *messageNamesSuffixedWithMsgVisitor) VisitMessage(message *parser.Message) bool {
	if !strings.HasSuffix(message.MessageName, "Msg") {
		v.AddFailuref(message.Meta.Pos, "Message name %q must be suffixed with Msg", message.MessageName)
	}
	return true
}
