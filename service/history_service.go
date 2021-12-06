package service

import(
	"encoding/json"
	"fmt"
	"github.com/qudj/fcc_rpc/models"
)

func SaveHistory(pre, cur interface{}, table, obKey, obType, opId string) error {
	historyByte, _ := json.Marshal(pre)
	changeByte, _ := json.Marshal(cur)
	history := models.FccHistoryLog{
		Table:       table,
		ObjectKey:   obKey,
		ObjectType:  obType,
		OpId:        opId,
		ChangeData:  string(changeByte),
		HistoryData: string(historyByte),
	}
	if err := models.SaveHistory(&history); err != nil {
		fmt.Println(fmt.Sprintf("save histroy error=%v", err))
		return err
	}
	return nil
}
