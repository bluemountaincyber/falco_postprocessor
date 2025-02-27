package outputs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func WriteToMonitor(output []byte, stream_url string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancel()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}
	token, err := credential.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{"https://management.azure.com/.default"},
	})
	if err != nil {
		return err
	}
	fmt.Println(token.Token)
	os.Exit(0)
	return nil
}
