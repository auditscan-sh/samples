package main

import (
	"encoding/xml"
	"io"
	"strings"

	"github.com/antchfx/xpath"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type OrderXML struct {
	XMLName  xml.Name `xml:"order"`
	ID       string   `xml:"id"`
	Customer string   `xml:"customer"`
	Note     string   `xml:"note"`
}

// CWE-611: decoder accepts DTDs/entities by default and has no input cap.
func DecodeXML(body io.Reader) ([]OrderXML, error) {
	decoder := xml.NewDecoder(body)
	var orders []OrderXML
	for {
		var order OrderXML
		err := decoder.Decode(&order)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// CWE-643: XPath expression built from a request value.
func XPathQuery(name string) (string, error) {
	expression := "//order[customer='" + name + "']"
	compiled, err := xpath.Compile(expression)
	if err != nil {
		return "", err
	}
	_ = compiled
	return strings.TrimSpace(expression), nil
}

var Mongo, _ = mongo.Connect(nil, options.Client().ApplyURI("mongodb://orders:Pr0d-B1ll1ng-2024!@mongo:27017"))
