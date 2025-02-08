using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Railway = Pulumi.Railway;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Railway.Random("myRandomResource", new()
    {
        Length = 24,
    });

    var myRandomComponent = new Railway.RandomComponent("myRandomComponent", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

