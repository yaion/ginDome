package v1

import (
	"testing"
)

func TestDb(t *testing.T)  {
	res,err:= DB.Exec("insert into user(name,age)value (?,?)","tome",23)
	if err !=nil{
		t.Log(err)
		t.Error("fail")
	}
	id,err:=res.LastInsertId()
	if err !=nil{
		t.Log(err)
		t.Error("fail")
	}
	t.Log(id)
}
