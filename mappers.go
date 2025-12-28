package main

import (
	"encoding/xml"
	"fmt"
)

type Mapper interface {
	Map(data []byte) (interface{}, error)
	GetType() string
}

type User struct {
	XMLName xml.Name `xml:"user"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
	Age     int      `xml:"age"`
}

type Product struct {
	XMLName     xml.Name `xml:"product"`
	ID          string   `xml:"id,attr"`
	Title       string   `xml:"title"`
	Price       float64  `xml:"price"`
	Category    string   `xml:"category"`
	InStock     bool     `xml:"inStock"`
	Description string   `xml:"description"`
}

type Order struct {
	XMLName   xml.Name    `xml:"order"`
	ID        string      `xml:"id,attr"`
	Date      string      `xml:"date"`
	Customer  string      `xml:"customer"`
	Total     float64     `xml:"total"`
	Items     []OrderItem `xml:"items>item"`
}

type OrderItem struct {
	ProductID string  `xml:"productId"`
	Quantity  int     `xml:"quantity"`
	Price     float64 `xml:"price"`
}

type UserMapper struct{}

func (um *UserMapper) Map(data []byte) (interface{}, error) {
	var user User
	err := xml.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to map user: %v", err)
	}
	return user, nil
}

func (um *UserMapper) GetType() string {
	return "user"
}

type ProductMapper struct{}

func (pm *ProductMapper) Map(data []byte) (interface{}, error) {
	var product Product
	err := xml.Unmarshal(data, &product)
	if err != nil {
		return nil, fmt.Errorf("failed to map product: %v", err)
	}
	return product, nil
}

func (pm *ProductMapper) GetType() string {
	return "product"
}

type OrderMapper struct{}

func (om *OrderMapper) Map(data []byte) (interface{}, error) {
	var order Order
	err := xml.Unmarshal(data, &order)
	if err != nil {
		return nil, fmt.Errorf("failed to map order: %v", err)
	}
	return order, nil
}

func (om *OrderMapper) GetType() string {
	return "order"
}

type MapperRegistry struct {
	mappers map[string]Mapper
}

func NewMapperRegistry() *MapperRegistry {
	return &MapperRegistry{
		mappers: make(map[string]Mapper),
	}
}

func (mr *MapperRegistry) Register(mapper Mapper) {
	mr.mappers[mapper.GetType()] = mapper
}

func (mr *MapperRegistry) Get(mapperType string) (Mapper, bool) {
	mapper, exists := mr.mappers[mapperType]
	return mapper, exists
}

func (mr *MapperRegistry) MapXML(mapperType string, xmlData []byte) (interface{}, error) {
	mapper, exists := mr.Get(mapperType)
	if !exists {
		return nil, fmt.Errorf("mapper for type '%s' not found", mapperType)
	}
	return mapper.Map(xmlData)
}