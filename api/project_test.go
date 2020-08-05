package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/squarescale/squarescale-go-sdk/api"
	"github.com/squarescale/squarescale-go-sdk/models"
)

func TestCreateProject(t *testing.T) {
	t.Run("succesful creation with all options", projectCreatedAllOptions)
	t.Run("error returned by the backend", projectCreationError)
	// todo test context cancelation ?
}

func projectCreatedAllOptions(t *testing.T) {
	project := models.Project{
		Name:           "new-project",
		UUID:           "990da244-77f7-4505-a933-4a6b4ad685cd",
		Provider:       "aws",
		Region:         "eu-west-1",
		CredentialName: "cred-aws-1",
		Clusters: []models.Cluster{
			models.Cluster{
				Type:     models.SingleNode,
				NodeType: "dev",
			},
		},
		NotificationConfigurations: []models.NotificationConfiguration{
			models.NotificationConfiguration{
				Type: "slack",
				URL:  "https://my-app.hook.slack",
			},
		},
		Organization: &models.Organization{
			Name: "my-company",
		},
		MonitoringConfiguration: &models.MonitoringConfiguration{
			Engine: "netdata",
		},
		ManagedServices: models.ManagedServices{
			Databases: []models.Database{
				models.Database{
					Engine:  "postgres",
					Version: "12",
					Size:    "dev",
				},
			},
		},
	}

	apiKey := "5b426281-34f0-4dc1-9d73-601b55852826"

	c := make(chan bool, 1)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { c <- true }()
		authHeader := r.Header.Get("Authorization")

		expectedAuthHeader := "bearer " + apiKey

		if authHeader != expectedAuthHeader {
			t.Fatalf("Expect %s, got %s", expectedAuthHeader, authHeader)
		}

		w.WriteHeader(http.StatusCreated)

	}))
	defer s.Close()

	client := api.NewClient(apiKey)
	client.SetCustomAddress(s.URL)

	// when
	err := client.CreateProject(project, context.Background())

	if err != nil {
		t.Fatalf("Expect no exception, go %s", err)
	}
	<-c
}

func projectCreationError(t *testing.T) {
	project := models.Project{
		Name:           "new-project",
		UUID:           "990da244-77f7-4505-a933-4a6b4ad685cd",
		Provider:       "aws",
		Region:         "eu-west-1",
		CredentialName: "cred-aws-1",
		Clusters: []models.Cluster{
			models.Cluster{
				Type:     models.SingleNode,
				NodeType: "dev",
			},
		},
		NotificationConfigurations: []models.NotificationConfiguration{},
		Organization: &models.Organization{
			Name: "my-company",
		},
		ManagedServices: models.ManagedServices{
			Databases: []models.Database{},
		},
	}

	apiKey := "5b426281-34f0-4dc1-9d73-601b55852826"

	c := make(chan bool, 1)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { c <- true }()

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "something wrong happened"}`))

	}))
	defer s.Close()

	client := api.NewClient(apiKey)
	client.SetCustomAddress(s.URL)

	// when
	err := client.CreateProject(project, context.Background())

	if err == nil {
		t.Fatal("Expect an exception, go nil")
	}

	expectedError := `Unexpected http code receive (400) with the message: {"error": "something wrong happened"}`

	if err.Error() != expectedError {
		t.Fatalf("Expect %s, go %s", expectedError, err.Error())
	}

	<-c
}
