import * as pulumi from "@pulumi/pulumi";
import * as railway from "@pulumi/railway";

const myProject = new railway.Project("myProject", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    description: "My first Railway project",
    defaultEnvironmentName: "Default Environment",
    isPublic: true,
    prDeploys: true,
    runtime: "LEGACY"
});
const myTestEnvironment = new railway.Environment("myTestEnvironment", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    projectId: myProject.projectId,
    skipInitialDeploys: true,
    stageInitialChanges: true
});
const myStagingEnvironment = new railway.Environment("myStagingEnvironment", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    projectId: myProject.projectId,
});
const myTestService = new railway.Service("myTestService", {
    projectId: myProject.projectId,
    environmentId: myTestEnvironment.environmentId,
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
});
const myStagingService = new railway.Service("myStagingService", {
    projectId: myProject.projectId,
    environmentId: myStagingEnvironment.environmentId,
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
});
export const output = {
    project: myProject,
    testEnvironment: myTestEnvironment,
    testService: myTestService,
    stagingEnvironment: myStagingEnvironment,
    stagingService: myStagingService,
};
