package dto

type AccountInfo struct {
	accDID        string
	accName       string
	accRole       uint8
	leaderDID     string
	platformState uint8
	operatorState uint8
	field         string
}

func NewAccountInfo(accDID string, accName string, accRole uint8, leaderDID string, platformState uint8, operatorState uint8, field string) *AccountInfo {
	return &AccountInfo{accDID: accDID, accName: accName, accRole: accRole, leaderDID: leaderDID, platformState: platformState, operatorState: operatorState, field: field}
}
