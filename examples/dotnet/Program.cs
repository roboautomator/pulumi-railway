using System.Collections.Generic;
using System.Linq;
using Pulumi;
using railway = Pulumi.railway;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new railway.Random("myRandomResource", new()
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

