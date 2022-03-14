namespace go api


struct scheduleRequest {
  1: i64 frequency
}

struct scheduleResponse {
  1: bool success
}

struct Indicator {}
struct IndicatorJson {}
struct AddIndicatorRequest {
  1: Indicator indicator
  2: IndicatorJson indicatorJson
}

struct AddIndicatorResponse {
  1: bool success
}

struct DeleteIndicatorRequest {
  1: string indicatorCode
}

struct DeleteIndicatorResponse {
  1: bool success
}

struct QueryIndicatorRequest {
  1: string indicatorCode
}

struct QueryIndicatorResponse {
  1: Indicator indicator
}

struct ModifyIndicatorRequest {
  1: Indicator indicator
}

struct ModifyIndicatorResponse {
  1: bool success
}

struct Rule {}
struct RuleJson {}
struct AddRuleRequest {
  1: Rule rule
  2: RuleJson ruleJson
}

struct AddRuleResponse {
  1: bool success
}

struct DeleteRuleRequest {
  1: string ruleCode
}

struct DeleteRuleResponse {
  1: bool success
}

struct QueryRuleRequest {
  1: string ruleCode
}

struct QueryRuleResponse {
  1: Rule rule
}

struct ModifyRuleRequest {
  1: Rule rule
}

struct ModifyRuleResponse {
  1: bool success
}


service Schedule {
    scheduleResponse schedule(1: scheduleRequest req)

    AddIndicatorResponse addIndicator(1: AddIndicatorRequest req)
    DeleteIndicatorResponse deleteIndicator(1:DeleteIndicatorRequest req)
    QueryIndicatorResponse queryIndicator(1:QueryIndicatorRequest req)
    ModifyIndicatorResponse modifyIndicator(1:ModifyIndicatorRequest req)

    AddRuleResponse addRule(1: AddRuleRequest req)
    DeleteRuleResponse deleteRule(1:DeleteRuleRequest req)
    QueryRuleResponse queryRule(1:QueryRuleRequest req)
    ModifyRuleResponse modifyRule(1:ModifyRuleRequest req)
}
