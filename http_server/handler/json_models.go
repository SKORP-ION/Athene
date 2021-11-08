package handler

type Response struct {
	Data interface{}
}

type AddPropertyRequest struct {
	Note       string
	Properties []Property
}

type Property struct {
	Id        uint32
	Inventory string
	Serial    string
	Name      string
}

type PropGroup struct {
	Ids     []uint32
	GroupId uint32
	Note    string
}

type SendToWarehouseReq struct {
	Ids         []uint32
	WarehouseId uint32
	Note        string
}

type GiveReq struct {
	EmployeeId uint32
	Ids        []uint32
	Note       string
}

type PrintReq struct {
	RecordId uint32
	Ticket   string
}

type Employee struct {
	Id         uint32
	Name       string
	Table      string
	ShortName  string
	Job        string
	Department string
	Manager    string
}

type PrintData struct {
	Employee
	Prop   Property
	Date   string
	Ticket string
}
