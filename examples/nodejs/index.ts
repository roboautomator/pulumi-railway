import * as pulumi from "@pulumi/pulumi";
import * as railway from "@pulumi/railway";

const myProject = new railway.Project("myProject", {apiToken: "9fc34a78-1e12-4453-ae87-055803d35715"});
const myTestEnvironment = new railway.Environment("myTestEnvironment", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    projectId: myProject.projectId,
});
const myStagingEnvironment = new railway.Environment("myStagingEnvironment", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    projectId: myProject.projectId,
});
const myTestService = new railway.Service("myTestService", {
    projectId: myProject.projectId,
    environmentId: myTestEnvironment.environmentId,
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    name: "My Test Service",
});
const myStagingService = new railway.Service("myStagingService", {
    projectId: myProject.projectId,
    environmentId: myStagingEnvironment.environmentId,
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    name: "My Staging Service",
});
export const output = {
    project: myProject,
    testEnvironment: myTestEnvironment,
    testService: myTestService,
    stagingEnvironment: myStagingEnvironment,
    stagingService: myStagingService,
};
