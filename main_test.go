package main

import (
	"context"
	"flag"
	"os"
	"regexp"
	"testing"

	"dagger.io/dagger"
)

func TestFoo(t *testing.T) {

}

func TestDagger(t *testing.T) {
	if m := flag.Lookup("test.run").Value.String(); m == "" || !regexp.MustCompile(m).MatchString(t.Name()) {
		t.Skip("skipping as execution was not requested explicitly using go test -run")
	}

	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stderr))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	args := []string{"go", "test", "./..."}

	output, err := client.Container().
		From("golang:1.19-alpine").
		WithEnvVariable("CGO_ENABLED", "0").
		WithWorkdir("/app").
		WithMountedDirectory("/app", client.Host().Workdir()).
		WithWorkdir("/app").
		WithEnvVariable("CGO_ENABLED", "0").
		Exec(dagger.ContainerExecOpts{
			Args: args,
		}).
		Stdout().Contents(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	t.Log(output)
}
