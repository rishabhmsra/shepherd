package provisioning

import (
	"github.com/rancher/shepherd/extensions/provisioninginput"
	"github.com/rancher/shepherd/pkg/config"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

type CustomClusterConfig struct {
	ExternalNodeProvider ExternalNodeProvider `json:"externalNodeProvider" yaml:"externalNodeProvider"`
	NodeLabels           map[string]string    `json:"nodeLabels" yaml:"nodeLabels"`
	NodeTaints           []corev1.Taint       `json:"nodeTaints" yaml:"nodeTaints"`
	SpecifyPrivateIP     bool                 `json:"specifyPrivateIP" yaml:"specifyPrivateIP"`
	SpecifyPublicIP      bool                 `json:"specifyPublicIP" yaml:"specifyPublicIP"`
	NodeNamePrefix       string               `json:"nodeNamePrefix" yaml:"nodeNamePrefix"`
}

func GetInstallDockerURL() string {
	installDockerCfg := &provisioninginput.RKE1CustomClusterDockerInstall{}
	config.LoadConfig(provisioninginput.InstallDockerConfigFileKey, installDockerCfg)
	logrus.Info(installDockerCfg.InstallDockerURL)
	return installDockerCfg.InstallDockerURL
}
