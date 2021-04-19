package resources

import (
	"context"

	c "infra-gw/src/cont"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

func CreateSecrets(secretName string, secretNamespace string, secretData map[string][]byte, appCtx *c.AppContext) error{
	bg := context.Background()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: secretName},
		Data: secretData,
	}
	secret, err := appCtx.K8s.Clientset.CoreV1().Secrets(secretNamespace).Create(bg, secret, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
