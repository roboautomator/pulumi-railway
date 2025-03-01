import pulumi
import pulumi_railway as railway

my_project = railway.Project("myProject", api_token="9fc34a78-1e12-4453-ae87-055803d35715")
my_test_environment = railway.Environment("myTestEnvironment",
    api_token="9fc34a78-1e12-4453-ae87-055803d35715",
    project_id=my_project.project_id)
my_staging_environment = railway.Environment("myStagingEnvironment",
    api_token="9fc34a78-1e12-4453-ae87-055803d35715",
    project_id=my_project.project_id)
my_test_service = railway.Service("myTestService",
    project_id=my_project.project_id,
    environment_id=my_test_environment.environment_id,
    api_token="9fc34a78-1e12-4453-ae87-055803d35715",
    name="My Test Service")
my_staging_service = railway.Service("myStagingService",
    project_id=my_project.project_id,
    environment_id=my_staging_environment.environment_id,
    api_token="9fc34a78-1e12-4453-ae87-055803d35715",
    name="My Staging Service")
pulumi.export("output", {
    "project": my_project,
    "testEnvironment": my_test_environment,
    "testService": my_test_service,
    "stagingEnvironment": my_staging_environment,
    "stagingService": my_staging_service,
})
