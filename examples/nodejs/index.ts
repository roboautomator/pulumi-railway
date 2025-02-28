import * as pulumi from "@pulumi/pulumi";
import * as railway from "@pulumi/railway";
import { env } from "process";

// const myRandomResource = new railway.Random("myRandomResource", {length: 24});
// const myRandomComponent = new railway.RandomComponent("myRandomComponent", {length: 24});
// export const output = {
//     value: myRandomResource.result,
// };

const myService = new railway.Service("myService", {
    environmentId: "73de0eab-b6dd-426b-a155-2a7e32f332df",
    projectId: "2a29134b-7629-4065-899c-7bed50d175e8",
    apiToken: "b8d195d4-ef57-470d-bd04-164fba27d09c"
})

myService.environmentId.apply((envId) => {
    console.log(envId)
})