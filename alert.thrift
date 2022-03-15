namespace go api

struct Indicator {
  1: string IndicatorCode
  2: string Name
  3: string Expression
  4: string Description
  5: string TimeCreate
  6: string TimeUpdate
}

struct IndicatorVO {
  1: string IndicatorCode
  2: string Name
  3: list<IndicatorVO> Indicators
  4: i64 Calculate
  5: string	Value
  6: string	Description
  7: string TimeCreate
  8: string TimeUpdate
}

struct IndicatorJson {
  1: list<IndicatorVO> Indicators
  2: i64 Calculate
  3: string	Value
}
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

struct Rule {
  1: string RuleCode
  2: string RuleName
  3: i64 RoomId
  4: string Expression
  5: string Description
  6: string TimeStart
  7: string TimeEnd
  8: string TimeCreate
  9: string TimeUpdate
}

struct RuleJson {
  1: list<RuleVo> Rules
  2: i64 Logic
  3: i64 Op
  4: i64 Value
  5: string IndicatorCode
}

struct RuleVo {
	1: string RuleCode
	2: string RuleName
	3: i64 RoomId
	4: list<RuleVo> Rules
	5: i64 Logic
	6: i64 Op
	7: i64 Value
	8: string IndicatorCode
	9: string Description
	10: string TimeStart
	11: string TimeEnd
	12: string TimeCreate
	13: string TimeUpdate
}

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

struct Task {
    1: string TaskCode
	2: string TaskName
	3: string RuleCode
	4: i64 Frequency
	6: string NextTime
	7: bool Status
}

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
