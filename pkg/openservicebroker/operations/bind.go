package operations

import (
	"net/http"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/open-service-broker-sdk/pkg/apis/broker"
	"github.com/openshift/open-service-broker-sdk/pkg/openservicebroker"
)

// Bind handles bind requests from the service catalog by returning
// a bind response with credentials for the service instance.
func (b *BrokerOperations) Bind(instanceID, bindingID string, breq *openservicebroker.BindRequest) *openservicebroker.Response {
	// Find the service instance that is being bound to
	si, err := b.Client.Broker().ServiceInstances(broker.Namespace).Get(instanceID, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return &openservicebroker.Response{Code: http.StatusGone, Body: &openservicebroker.BindResponse{}, Err: nil}
		}
		return &openservicebroker.Response{Code: http.StatusInternalServerError, Body: nil, Err: err}
	}

	// in principle, bind should alter state somewhere

	// Create some credentials to return.  In this case the credentials are
	// pulled from the service instance but a real broker might
	// return unique credentials for each binding so that multiple users
	// of a service instance are not sharing credentials.
	credentials := map[string]interface{}{}
	credentials["credential"] = si.Spec.Credential

	return &openservicebroker.Response{
		Code: http.StatusCreated,
		Body: &openservicebroker.BindResponse{Credentials: credentials},
		Err:  nil,
	}
}
