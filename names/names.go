package names

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

type ServiceDatum struct {
	Aliases           []string
	GoPackage         string
	HumanFriendly     string
	ProviderNameUpper string
}

// serviceData key is the AWS provider service package
var serviceData map[string]*ServiceDatum

//go:embed names_data.csv
var namesData string

func init() {
	serviceData = make(map[string]*ServiceDatum)

	// Data from names_data.csv
	if err := readCSVIntoServiceData(); err != nil {
		log.Fatalf("reading CSV into service data: %s", err)
	}
}

func readCSVIntoServiceData() error {
	// names_data.csv is dynamically embedded so changes, additions should be made
	// there also

	r := csv.NewReader(strings.NewReader(namesData))

	d, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("reading CSV into service data: %w", err)
	}

	for i, l := range d {
		if i < 1 { // omit header line
			continue
		}

		if l[ColExclude] != "" {
			continue
		}

		p := l[ColProviderPackage]

		serviceData[p] = &ServiceDatum{
			GoPackage:         l[ColGoPackage],
			HumanFriendly:     l[ColHumanFriendly],
			ProviderNameUpper: l[ColProviderNameUpper],
		}

		a := []string{p}

		if l[ColAliases] != "" {
			a = append(a, strings.Split(l[ColAliases], ";")...)
		}

		serviceData[p].Aliases = a
	}

	return nil
}

func ProviderPackageForAlias(serviceAlias string) (string, error) {
	for k, v := range serviceData {
		for _, hclKey := range v.Aliases {
			if serviceAlias == hclKey {
				return k, nil
			}
		}
	}

	return "", fmt.Errorf("unable to find service for service alias %s", serviceAlias)
}

func ProviderPackages() []string {
	keys := make([]string, len(serviceData))

	i := 0
	for k := range serviceData {
		keys[i] = k
		i++
	}

	return keys
}

func ProviderNameUpper(service string) (string, error) {
	if v, ok := serviceData[service]; ok {
		return v.ProviderNameUpper, nil
	}

	return "", fmt.Errorf("no service data found for %s", service)
}

func FullHumanFriendly(service string) (string, error) {
	if v, ok := serviceData[service]; ok {
		// We hard code the "Brand" as Cloudera
		return fmt.Sprintf("%s %s", "Cloudera", v.HumanFriendly), nil
	}

	if s, err := ProviderPackageForAlias(service); err == nil {
		return FullHumanFriendly(s)
	}

	return "", fmt.Errorf("no service data found for %s", service)
}

func HumanFriendly(service string) (string, error) {
	if v, ok := serviceData[service]; ok {
		return v.HumanFriendly, nil
	}

	if s, err := ProviderPackageForAlias(service); err == nil {
		return HumanFriendly(s)
	}

	return "", fmt.Errorf("no service data found for %s", service)
}

func CDPGoPackage(providerPackage string) (string, error) {
	if v, ok := serviceData[providerPackage]; ok {
		return v.GoPackage, nil
	}

	return "", fmt.Errorf("getting CDP SDK Go package, %s not found", providerPackage)
}
