// Code generated by goctl. DO NOT EDIT.
package types

type Notification struct {
	NotificationID int64  `json:"id"`
	AssigneeID     int64  `json:"assigneeID"`
	AssignerID     int64  `json:"assignerID"`
	RefID          int64  `json:"refID"`
	RefType        int64  `json:"refType"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Priority       int64  `json:"priority"`
	DueDate        int64  `json:"dueDate"`
	Status         int64  `json:"status"`
	Unread         int64  `json:"unread"`
	CreatedAt      int64  `json:"createdAt"`
}

type CreateNotificationReq struct {
	Sender      int64  `form:"sender"`
	Receiver    int64  `form:"receiver"`
	RefID       int64  `form:"refID"`
	RefType     int64  `form:"refType"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Priority    int64  `form:"priority"`
	DueDate     int64  `form:"dueDate"`
}

type CreateNotificationRes struct {
	Result       Result       `json:"result"`
	Notification Notification `json:"notification"`
}

type Result struct {
	Code    int    `json:"code"`    //    Result code: 0 is success. Otherwise, getting an error
	Message string `json:"message"` // Result message: detail response code
}

type User struct {
	UserID      int64  `json:"userID"`
	Phone       string `json:"phone"`
	Role        int64  `json:"role"`
	Status      int64  `json:"status"`
	Address     string `json:"address"`
	FullName    string `json:"fullName"`
	AvatarUrl   string `json:"avatarUrl"`
	Birthday    int64  `json:"birthday"`
	Gender      int64  `json:"gender"`
	CccdNumber  string `json:"cccdNumber"`
	CccdDate    int64  `json:"cccdDate"`
	CccdAddress string `json:"cccdAddress"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

type House struct {
	HouseID     int64     `json:"houseID"`
	User        User      `json:"user"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        int64     `json:"type"`
	Status      int64     `json:"status"`
	Area        int64     `json:"area"`
	Price       int64     `json:"price"`
	BedNum      int64     `json:"bedNum"`
	LivingNum   int64     `json:"livingNum"`
	Albums      []string  `json:"albums"`
	Rooms       []Room    `json:"rooms"`
	Services    []Service `json:"services"`
	Address     string    `json:"address"`
	WardID      int64     `json:"wardID"`
	DistrictID  int64     `json:"districtID"`
	ProvinceID  int64     `json:"provinceID"`
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
	CreatedBy   int64     `json:"createdBy"`
	UpdatedBy   int64     `json:"updatedBy"`
}

type Album struct {
	AlbumID int64  `json:"albumID"`
	HouseID int64  `json:"houseID"`
	Url     string `json:"url"`
}

type Room struct {
	RoomID    int64  `json:"roomID"`
	HouseID   int64  `json:"houseID"`
	Name      string `json:"name"`
	HouseName string `json:"houseName"`
	Area      int64  `json:"area"`
	Price     int64  `json:"price"`
	Type      int64  `json:"type"`
	Status    int64  `json:"status"`
	Capacity  int64  `json:"capacity"`
	EIndex    int64  `json:"eIndex"`
	WIndex    int64  `json:"wIndex"`
}

type Service struct {
	ServiceID int64  `json:"serviceID"`
	HouseID   int64  `json:"houseID"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Unit      int64  `json:"unit"`
}

type Contract struct {
	ContractID    int64   `json:"contractID"`
	Code          string  `json:"code"`
	Status        int64   `json:"status"`
	RenterID      int64   `json:"renterID"`
	RenterPhone   string  `json:"renterPhone"`
	RenterNumber  string  `json:"renterNumber"`
	RenterDate    int64   `json:"renterDate"`
	RenterAddress string  `json:"renterAddress"`
	RenterName    string  `json:"renterName"`
	LessorID      int64   `json:"lessorID"`
	LessorPhone   string  `json:"lessorPhone"`
	LessorNumber  string  `json:"lessorNumber"`
	LessorDate    int64   `json:"lessorDate"`
	LessorAddress string  `json:"lessorAddress"`
	LessorName    string  `json:"lessorName"`
	Room          Room    `json:"room"`
	CheckIn       int64   `json:"checkIn"`
	Duration      int64   `json:"duration"`
	Purpose       string  `json:"purpose"`
	Payment       Payment `json:"payment"`
	CreatedAt     int64   `json:"createdAt"`
	UpdatedAt     int64   `json:"updatedAt"`
	CreatedBy     int64   `json:"createdBy"`
	UpdatedBy     int64   `json:"updatedBy"`
}

type PaymentRenter struct {
	ID        int64  `json:"id"`
	PaymentID int64  `json:"paymentID"`
	RenterID  int64  `json:"renterID"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
}

type PaymentDetail struct {
	ID        int64  `json:"id"`
	PaymentID int64  `json:"paymentID"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Type      int64  `json:"type"`
}

type Payment struct {
	PaymentID      int64           `json:"paymentID"`
	ContractID     int64           `json:"contractID"`
	Amount         int64           `json:"amount"`
	Discount       int64           `json:"discount"`
	Deposit        int64           `json:"deposit"`
	DepositDate    int64           `json:"depositDate"`
	NextBill       int64           `json:"nextBill"`
	PaymentRenters []PaymentRenter `json:"paymentRenters"`
	PaymentDetails []PaymentDetail `json:"paymentDetails"`
}

type Bill struct {
	BillID       int64        `json:"billID"`
	Title        string       `json:"title"`
	ContractCode string       `json:"contractCode"`
	PaymentID    int64        `json:"paymentID"`
	PaymentDate  int64        `json:"paymentDate"`
	Amount       int64        `json:"amount"`
	Remain       int64        `json:"remain"`
	Status       int64        `json:"status"`
	BillDetails  []BillDetail `json:"billDetails"`
}

type BillDetail struct {
	BillDetailID int64  `json:"billDetailID"`
	BillID       int64  `json:"billID"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Type         int64  `json:"type"`
	Quantity     int64  `json:"quantity"`
}

type BillPay struct {
	BillPayID int64 `json:"billPayID"`
	BillID    int64 `json:"billID"`
	Amount    int64 `json:"amount"`
	PayDate   int64 `json:"payDate"`
	UserID    int64 `json:"userID"`
}
