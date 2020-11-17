package main

import (
	"github.com/yoheimuta/protolint/_example/plugin/customrules"
	"github.com/yoheimuta/protolint/plugin"
)

func main() {
	plugin.RegisterCustomRules(
		customrules.NewEnumFieldNamesRequireZeroValueRule(),
		customrules.NewEnumNamesNoMsgSuffixRule(),
		customrules.NewMessageNamesSuffixedWithMsgRule(),
		customrules.NewPackageNameLowerCaseUnderscoreRule(),
		customrules.NewProto2FieldsAvoidRequiredRule(),
		customrules.NewRepeatedFieldNamesSingularizedRule(),
	)
}
