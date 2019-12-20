// Code generated by "stringer -type=Consistency,TimeUnit,TableState,OperationState,DbType,PutOption -output types_string.go"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Absolute-1]
	_ = x[Eventual-2]
}

const _Consistency_name = "AbsoluteEventual"

var _Consistency_index = [...]uint8{0, 8, 16}

func (i Consistency) String() string {
	i -= 1
	if i < 0 || i >= Consistency(len(_Consistency_index)-1) {
		return "Consistency(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Consistency_name[_Consistency_index[i]:_Consistency_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Hours-1]
	_ = x[Days-2]
}

const _TimeUnit_name = "HoursDays"

var _TimeUnit_index = [...]uint8{0, 5, 9}

func (i TimeUnit) String() string {
	i -= 1
	if i < 0 || i >= TimeUnit(len(_TimeUnit_index)-1) {
		return "TimeUnit(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TimeUnit_name[_TimeUnit_index[i]:_TimeUnit_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Active-0]
	_ = x[Creating-1]
	_ = x[Dropped-2]
	_ = x[Dropping-3]
	_ = x[Updating-4]
}

const _TableState_name = "ActiveCreatingDroppedDroppingUpdating"

var _TableState_index = [...]uint8{0, 6, 14, 21, 29, 37}

func (i TableState) String() string {
	if i < 0 || i >= TableState(len(_TableState_index)-1) {
		return "TableState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TableState_name[_TableState_index[i]:_TableState_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownOpState-0]
	_ = x[Complete-1]
	_ = x[Working-2]
}

const _OperationState_name = "UnknownOpStateCompleteWorking"

var _OperationState_index = [...]uint8{0, 14, 22, 29}

func (i OperationState) String() string {
	if i < 0 || i >= OperationState(len(_OperationState_index)-1) {
		return "OperationState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationState_name[_OperationState_index[i]:_OperationState_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Array-0]
	_ = x[Binary-1]
	_ = x[Boolean-2]
	_ = x[Double-3]
	_ = x[Integer-4]
	_ = x[Long-5]
	_ = x[Map-6]
	_ = x[String-7]
	_ = x[Timestamp-8]
	_ = x[Number-9]
	_ = x[JSONNull-10]
	_ = x[Null-11]
	_ = x[Empty-12]
}

const _DbType_name = "ArrayBinaryBooleanDoubleIntegerLongMapStringTimestampNumberJSONNullNullEmpty"

var _DbType_index = [...]uint8{0, 5, 11, 18, 24, 31, 35, 38, 44, 53, 59, 67, 71, 76}

func (i DbType) String() string {
	if i < 0 || i >= DbType(len(_DbType_index)-1) {
		return "DbType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DbType_name[_DbType_index[i]:_DbType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PutIfAbsent-4]
	_ = x[PutIfPresent-5]
	_ = x[PutIfVersion-6]
}

const _PutOption_name = "PutIfAbsentPutIfPresentPutIfVersion"

var _PutOption_index = [...]uint8{0, 11, 23, 35}

func (i PutOption) String() string {
	i -= 4
	if i < 0 || i >= PutOption(len(_PutOption_index)-1) {
		return "PutOption(" + strconv.FormatInt(int64(i+4), 10) + ")"
	}
	return _PutOption_name[_PutOption_index[i]:_PutOption_index[i+1]]
}
