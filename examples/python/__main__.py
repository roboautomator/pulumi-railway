import pulumi
import pulumi_railway as railway

my_project = railway.Project("myProject", api_token="9fc34a78-1e12-4453-ae87-055803d35715")
my_environment = railway.Environment("myEnvironment",
    api_token="9fc34a78-1e12-4453-ae87-055803d35715",
    project_id=my_project.project_id)
my_service = railway.Service("myService",
    project_id=my_project.project_id,
    environment_id=my_environment.environment_id,
    api_token="9fc34a78-1e12-4453-ae87-055803d35715")
pulumi.export("output", {
    "project": my_project,
    "environment": my_environment,
    "service": my_service,
})
