import * as pulumi from "@pulumi/pulumi";
import * as railway from "@pulumi/railway";

const myRandomResource = new railway.Random("myRandomResource", {length: 24});
const myRandomComponent = new railway.RandomComponent("myRandomComponent", {length: 24});
export const output = {
    value: myRandomResource.result,
};
