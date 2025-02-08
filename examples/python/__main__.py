import pulumi
import pulumi_railway as railway

my_random_resource = railway.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
