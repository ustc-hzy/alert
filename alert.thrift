namespace go api



struct scheduleRequest {
  1: i64 frequency
}

struct scheduleResponse {
  1: bool success
}

service Schedule {
    scheduleResponse schedule(1: scheduleRequest req)
}
