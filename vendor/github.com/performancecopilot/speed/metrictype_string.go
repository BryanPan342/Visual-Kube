// Code generated by "stringer -type=MetricType"; DO NOT EDIT

package speed

import "fmt"

const _MetricType_name = "Int32TypeUint32TypeInt64TypeUint64TypeFloatTypeDoubleTypeStringType"

var _MetricType_index = [...]uint8{0, 9, 19, 28, 38, 47, 57, 67}

func (i MetricType) String() string {
	if i < 0 || i >= MetricType(len(_MetricType_index)-1) {
		return fmt.Sprintf("MetricType(%d)", i)
	}
	return _MetricType_name[_MetricType_index[i]:_MetricType_index[i+1]]
}
