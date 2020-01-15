package models

import "strconv"

type DefaultModel struct {
	UserId int64
}

//get real primary key name
func (defaultModel *DefaultModel) GetKey() string {
	return "{{package}}Id"
}

//get primary key in model
func (defaultModel *DefaultModel) GetKeyProperty() int64 {
	return defaultModel.UserId
}

//set primary key
func (defaultModel *DefaultModel) SetKeyProperty(id int64) {
	defaultModel.UserId = id
}

//get real table name
func (defaultModel *DefaultModel) TableName() string {
	return "{{package}}_info"
}

func (defaultModel *DefaultModel) GetById(id string) {
	defaultModel.UserId,_ = strconv.ParseInt(id, 10, 64)
}

func (defaultModel *DefaultModel) GetList(page, limit int64, condition string) (list []*DefaultModel) {
	list = make([]*DefaultModel, 2)
	list[0] = &DefaultModel{UserId:1}
	list[1] = &DefaultModel{UserId:2}
	return
}

func (defaultModel *DefaultModel) Create() {

}

func (defaultModel *DefaultModel) Update(info DefaultModel) {

}

func (defaultModel *DefaultModel) Delete() {

}