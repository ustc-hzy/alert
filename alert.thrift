namespace go api

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

struct Task {}

struct AddTaskRequest {
  1: Task task
}

struct AddTaskResponse {
  1: bool success
}

struct DeleteTaskRequest {
  1: string taskCode
}

struct DeleteTaskResponse {
  1: bool success
}

struct QueryTaskRequest {
  1: string taskCode
}

struct QueryTaskResponse {
  1: Task task
}

struct ModifyTaskRequest {
  1: Task task
}

struct ModifyTaskResponse {
  1: bool success
}




service CRUD {
    AddIndicatorResponse addIndicator(1: AddIndicatorRequest req)
    DeleteIndicatorResponse deleteIndicator(1: DeleteIndicatorRequest req)
    QueryIndicatorResponse queryIndicator(1: QueryIndicatorRequest req)
    ModifyIndicatorResponse modifyIndicator(1: ModifyIndicatorRequest req)

    AddRuleResponse addRule(1: AddRuleRequest req)
    DeleteRuleResponse deleteRule(1: DeleteRuleRequest req)
    QueryRuleResponse queryRule(1: QueryRuleRequest req)
    ModifyRuleResponse modifyRule(1: ModifyRuleRequest req)

    AddTaskResponse addTask(1: AddTaskRequest req)
    DeleteTaskResponse deleteTask(1: DeleteTaskRequest req)
    QueryTaskResponse queryTask(1: QueryTaskRequest req)
    ModifyTaskResponse modifyTask(1: ModifyTaskRequest req)
}
