package notifier

import (
	"bytes"
	"encoding/json"
	"github.com/wso2/product-microgateway/adapter/config"
	"github.com/wso2/product-microgateway/adapter/internal/auth"
	logger "github.com/wso2/product-microgateway/adapter/internal/loggers"
	"github.com/wso2/product-microgateway/adapter/pkg/tlsutils"
	"net/http"
	"strings"
)

const (
	deployedRevisionEP string = "internal/data/v1/apis/deployed-revisions"
)

//UpdateDeployedRevisions create the DeployedAPIRevision object
func UpdateDeployedRevisions(apiID string, revisionID int, envs []string, vhost string) *DeployedAPIRevision {
	revisions := &DeployedAPIRevision{
		APIID:      apiID,
		RevisionID: revisionID,
		EnvInfo:    []DeployedEnvInfo{},
	}
	for _, env := range envs {
		info := DeployedEnvInfo{
			Name:  env,
			VHost: vhost,
		}
		revisions.EnvInfo = append(revisions.EnvInfo, info)
	}
	return revisions
}

//SendRevisionUpdate sends deployment status to the control plane
func SendRevisionUpdate(deployedRevisionList []*DeployedAPIRevision) {
	logger.LoggerNotifier.Debugf("Revision deployed message is sending to Control plane")
	conf, _ := config.ReadConfigs()
	cpConfigs := conf.ControlPlane

	revisionEP := cpConfigs.ServiceURL
	if strings.HasSuffix(revisionEP, "/") {
		revisionEP += deployedRevisionEP
	} else {
		revisionEP += "/" + deployedRevisionEP
	}

	if len(deployedRevisionList) < 1 || !cpConfigs.Enabled {
		return
	}

	jsonValue, _ := json.Marshal(deployedRevisionList)

	// Setting authorization header
	basicAuth := "Basic " + auth.GetBasicAuth(cpConfigs.Username, cpConfigs.Password)

	logger.LoggerNotifier.Debugf("Revision deployed message sending to Control plane: %v", string(jsonValue))

	// Adding 3 retries for revision update sending
	retries := 0
	for retries < 3 {
		retries++

		req, _ := http.NewRequest("PATCH", revisionEP, bytes.NewBuffer(jsonValue))
		req.Header.Set("Authorization", basicAuth)
		req.Header.Set("Content-Type", "application/json")
		resp, err := tlsutils.InvokeControlPlane(req, cpConfigs.SkipSSLVerification)

		success := true
		if err != nil {
			logger.LoggerNotifier.Warnf("Error response from %v for attempt %v : %v", revisionEP, retries, err.Error())
			success = false
		}
		if resp != nil && resp.StatusCode != http.StatusOK {
			logger.LoggerNotifier.Warnf("Error response status code %v from %v for attempt %v", resp.StatusCode, revisionEP, retries)
			success = false
		}
		if success {
			logger.LoggerNotifier.Infof("Revision deployed message sent to Control plane for attempt %v", retries)
			break
		}
	}
}
