package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
    if err := build(context.Background()); err != nil {
        fmt.Println(err)
    }
}

func build(ctx context.Context) error {
    fmt.Println("Building with Dagger")

    // initialize Dagger client
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
    if err != nil {
        return err
    }
    defer client.Close()


    // Get reference from local project
    src := client.Host().Directory(".")

    // Get `golang image`
    golang := client.Container().From("golang:latest")
    client.Container().WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-cache"))
    client.Container().WithExec([]string{"go", "mod", "download"})
    client.Container().WithMountedDirectory("/src", client.Host().Directory("."))

    // mount cloned repository into `golang` image
    golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

     // define the application build command
    path := "build/"
    golang = golang.WithExec([]string{"go", "build", "-o",path})

    // get reference to build output directory in container
    output := golang.Directory(path)

      // write contents of container build/ directory to the host
    _, err = output.Export(ctx, path)
    if err != nil {
        return err
    }

    return nil
}
