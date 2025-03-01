package main

import (
	"example.com/pulumi-railway/sdk/go/railway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
myProject, err := railway.NewProject(ctx, "myProject", &railway.ProjectArgs{
ApiToken: pulumi.String("9fc34a78-1e12-4453-ae87-055803d35715"),
})
if err != nil {
return err
}
myTestEnvironment, err := railway.NewEnvironment(ctx, "myTestEnvironment", &railway.EnvironmentArgs{
ApiToken: pulumi.String("9fc34a78-1e12-4453-ae87-055803d35715"),
ProjectId: myProject.ProjectId,
})
if err != nil {
return err
}
myStagingEnvironment, err := railway.NewEnvironment(ctx, "myStagingEnvironment", &railway.EnvironmentArgs{
ApiToken: pulumi.String("9fc34a78-1e12-4453-ae87-055803d35715"),
ProjectId: myProject.ProjectId,
})
if err != nil {
return err
}
myTestService, err := railway.NewService(ctx, "myTestService", &railway.ServiceArgs{
ProjectId: myProject.ProjectId,
EnvironmentId: myTestEnvironment.EnvironmentId,
ApiToken: pulumi.String("9fc34a78-1e12-4453-ae87-055803d35715"),
Name: pulumi.String("My Test Service"),
})
if err != nil {
return err
}
myStagingService, err := railway.NewService(ctx, "myStagingService", &railway.ServiceArgs{
ProjectId: myProject.ProjectId,
EnvironmentId: myStagingEnvironment.EnvironmentId,
ApiToken: pulumi.String("9fc34a78-1e12-4453-ae87-055803d35715"),
Name: pulumi.String("My Staging Service"),
})
if err != nil {
return err
}
ctx.Export("output", %!v(PANIC=Format method: fatal: An assertion has failed: tok: ))
return nil
})
}
