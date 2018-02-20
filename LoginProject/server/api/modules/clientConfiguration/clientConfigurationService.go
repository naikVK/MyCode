package clientConfiguration

import (
	"LoginProject/server/api/common/model"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

//GET CLIENT CONFIGURATION FROM DB
func GetClientConfigService(client model.Client) (model.Client, error) {
	logginghelper.LogDebug("IN: GetClientConfigService")
	clientdbObj, err := GetClientConfigeDAO(client.ClientId)

	if err != nil {
		logginghelper.LogError("GetClientConfigService GetGetClientConfigeDAO : ", err)
		return clientdbObj, err

	}
	return clientdbObj, nil
}
