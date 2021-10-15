package model

type HisSku struct {
	Id               string      `json:"id"`
	Sku              string      `json:"sku"`
	PartId           int         `json:"partId"`
	IceasyId         int         `json:"iceasyId"`
	IceasyNo         string      `json:"iceasyNo"`
	LinecardId       int         `json:"linecardId"`
	Linecard         string      `json:"linecard"`
	LinecardEn       interface{} `json:"linecardEn"`
	PartNo           string      `json:"partNo"`
	FirstCategoryId  interface{} `json:"firstCategoryId"`
	SecondCategoryId interface{} `json:"secondCategoryId"`
	CategoryId       int         `json:"categoryId"`
	Type             string      `json:"type"`
	Unpacking        string      `json:"unpacking"`
	UpdateTimestamp  interface{} `json:"updateTimestamp"`
	IndexCreateDate  int64       `json:"indexCreateDate"`
	Restriction      int         `json:"restriction"`
	UpperTitle       interface{} `json:"upperTitle"`
	Description      interface{} `json:"description"`
	DescriptionEn    interface{} `json:"descriptionEn"`
	Enclosure        interface{} `json:"enclosure"`
	Pack             interface{} `json:"pack"`
	Rohs             interface{} `json:"rohs"`
	Datasheet        interface{} `json:"datasheet"`
	Imguri           interface{} `json:"imguri"`
	TypeOrderBy      interface{} `json:"typeOrderBy"`
	BrandLogo        interface{} `json:"brandLogo"`
	DateCode         interface{} `json:"dateCode"`
}
