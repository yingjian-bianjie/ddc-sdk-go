package dto

type AccountInfo struct {
	AccDID        string
	AccName       string
	AccRole       uint8
	LeaderDID     string
	PlatformState uint8
	OperatorState uint8
	Field         string
}

func NewAccountInfo(accDID string, accName string, accRole uint8, leaderDID string, platformState uint8, operatorState uint8, field string) *AccountInfo {
	return &AccountInfo{AccDID: accDID, AccName: accName, AccRole: accRole, LeaderDID: leaderDID, PlatformState: platformState, OperatorState: operatorState, Field: field}
}
