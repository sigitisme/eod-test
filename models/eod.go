package models

import (
	"fmt"
)

// EOD is to hold eod data for before and after
type EOD struct {
	ID               int
	Name             string
	Age              int
	Balanced         float64
	PreviousBalanced float64
	AveragedBalanced float64
	FreeTransfer     int
	No2bThreadNo     string
	No3ThreadNo      string
	No1ThreadNo      string
	No2aThreadNo     string
}

//ToString is used for csv
func (e *EOD) ToString() string {
	return fmt.Sprintf("%d;%s;%d;%.2f;%s;%s;%.2f;%.2f;%s;%d;%s;",
		e.ID, e.Name, e.Age, e.Balanced, e.No2bThreadNo, e.No3ThreadNo,
		e.PreviousBalanced, e.AveragedBalanced, e.No1ThreadNo, e.FreeTransfer, e.No2aThreadNo)
}
