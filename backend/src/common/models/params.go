package models

import "errors"

type SortDir string

const (
	ASC  SortDir = "ASC"
	DESC SortDir = "DESC"
)

var sortDirMap = map[string]SortDir{
	string(ASC):  ASC,
	string(DESC): DESC,
}

func (s SortDir) ToString() string {
	return string(s)
}

func SortDirFromString(value string) (SortDir, error) {
	sortDir, ok := sortDirMap[value]
	if !ok {
		return "", errors.New("invalid SortDir value")
	}
	return sortDir, nil
}

type Sort string

const (
	AMOUNT_CURRENCY Sort = "amount_currency"
)

var sortMap = map[string]Sort{
	string(AMOUNT_CURRENCY): AMOUNT_CURRENCY,
}

func (s Sort) ToString() string {
	return string(s)
}

func SortFromString(value string) (Sort, error) {
	sort, ok := sortMap[value]
	if !ok {
		return "", errors.New("invalid Sort value")
	}
	return sort, nil
}
