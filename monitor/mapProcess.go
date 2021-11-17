package monitor

import "github.com/google/uuid"

type MapProcess map[uuid.UUID]*Proceso

func (m MapProcess) ToSlice() (pros []*Proceso) {
	for _, v := range m {
		pros = append(pros, v)
	}
	return pros
}
