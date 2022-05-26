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


func (a *AccountInfo) SetAccDID(accDID string) {
	a.accDID = accDID
}

func (a *AccountInfo) SetAccName(accName string) {
	a.accName = accName
}

func (a *AccountInfo) SetAccRole(accRole uint8) {
	a.accRole = accRole
}

func (a *AccountInfo) SetLeaderDID(leaderDID string) {
	a.leaderDID = leaderDID
}

func (a *AccountInfo) SetPlatformState(platformState uint8) {
	a.platformState = platformState
}

func (a *AccountInfo) SetOperatorState(operatorState uint8) {
	a.operatorState = operatorState
}

func (a *AccountInfo) SetField(field string) {
	a.field = field
}