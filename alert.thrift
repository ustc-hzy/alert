namespace go api

struct computeRequest {
  1: i64 indicatorID
  2: i64 roomID
}

struct computeResponse {
  1: i64 value
}

struct checkRequest {
  1: i64 ruleID
}

struct checkResponse {
  1: bool result
}

struct scheduleRequest {
  1: i64 frequency
}

struct scheduleResponse {
  1: bool success
}

service Compute {
    computeResponse compute(1: computeRequest req)
}

service Check {
    checkResponse check(1: checkRequest req)
}

service Schedule {
    scheduleResponse schedule(1: scheduleRequest req)
}
