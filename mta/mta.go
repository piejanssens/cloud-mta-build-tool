package mta

import (
	"gopkg.in/yaml.v2"

	"cloud-mta-build-tool/cmd/logs"
	"cloud-mta-build-tool/mta/models"
)

// Parse MTA file
func Parse(yamlContent []byte) (out models.MTA, err error) {
	mta := models.MTA{}
	// Format the YAML to struct's
	err = yaml.Unmarshal([]byte(yamlContent), &mta)
	if err != nil {
		logs.Logger.Error("Yaml file is not valid, Error: " + err.Error())
	}
	return mta, err
}

// Marshal - For edit purpose
func Marshal(in models.MTA) (mtads []byte, err error) {
	mtads, err = yaml.Marshal(&in)
	if err != nil {
		logs.Logger.Error(err.Error())
	}

	return mtads, err
}