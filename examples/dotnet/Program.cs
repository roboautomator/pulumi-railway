using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Railway = Pulumi.Railway;

return await Deployment.RunAsync(() => 
{
    var myProject = new Railway.Project("myProject", new()
    {
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
    });

    var myTestEnvironment = new Railway.Environment("myTestEnvironment", new()
    {
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
        ProjectId = myProject.ProjectId,
    });

    var myStagingEnvironment = new Railway.Environment("myStagingEnvironment", new()
    {
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
        ProjectId = myProject.ProjectId,
    });

    var myTestService = new Railway.Service("myTestService", new()
    {
        ProjectId = myProject.ProjectId,
        EnvironmentId = myTestEnvironment.EnvironmentId,
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
    });

    var myStagingService = new Railway.Service("myStagingService", new()
    {
        ProjectId = myProject.ProjectId,
        EnvironmentId = myStagingEnvironment.EnvironmentId,
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "project", myProject },
            { "testEnvironment", myTestEnvironment },
            { "testService", myTestService },
            { "stagingEnvironment", myStagingEnvironment },
            { "stagingService", myStagingService },
        },
    };
});

