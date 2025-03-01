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

    var myEnvironment = new Railway.Environment("myEnvironment", new()
    {
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
        ProjectId = myProject.ProjectId,
    });

    var myService = new Railway.Service("myService", new()
    {
        ProjectId = myProject.ProjectId,
        EnvironmentId = myEnvironment.EnvironmentId,
        ApiToken = "9fc34a78-1e12-4453-ae87-055803d35715",
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "project", myProject },
            { "environment", myEnvironment },
            { "service", myService },
        },
    };
});

