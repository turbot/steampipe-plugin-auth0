package main

import (
	"fmt"

	"github.com/auth0/go-auth0/management"
)

func main() {
	domain := "dev-steady-lark.us.auth0.com"
	clientId := ""
	secret := ""

	m, err := management.New(domain, management.WithClientCredentials(clientId, secret))
	if err != nil {
		fmt.Printf(err.Error())
	}

	// org, err := m.Organization.Read("org_4wACPmqLhOmqwawM")
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }
	// fmt.Printf("Organization: %s\n", org.GetDisplayName())

	org, err := m.Organization.List(
	// management.RequestOption(),
	)
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, orgg := range org.Organizations {
		fmt.Printf("Organization: %s\n", orgg)
	}

	// c := &management.Client{
	// 	Name:        auth0.String("Client Name"),
	// 	Description: auth0.String("Long description of client"),
	// }

	// err = m.Client.Create(c)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }

	// fmt.Printf("Created client %s\n", c.GetClientID())
}
