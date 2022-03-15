package core

const NIL = -1

type CaculateType int32

const (
	ADD      CaculateType = 0
	SUBTRACT CaculateType = 1
	MULTIPLY CaculateType = 2
	DIVIDE   CaculateType = 3
)

type OpType int32

const (
	LARGER   OpType = 0
	SMALLER  OpType = 1
	EQUAL    OpType = 2
	NOTEQUAL OpType = 3
)

type LogicType int32

const (
	AND LogicType = 0
	OR  LogicType = 1
)

const TIMETAMPLATE = "2006-01-02 15:04:05"
