import * as pulumi from "@pulumi/pulumi";
import * as railway from "@pulumi/railway";

const myProject = new railway.Project("myProject", {apiToken: "9fc34a78-1e12-4453-ae87-055803d35715"});
const myEnvironment = new railway.Environment("myEnvironment", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    projectId: myProject.projectId,
});

const testEnvironment = new railway.Environment("test", {
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
    projectId: myProject.projectId,
});

const myService = new railway.Service("myService", {
    projectId: myProject.projectId,
    environmentId: myEnvironment.environmentId,
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
});

const testService = new railway.Service("testService", {
    projectId: myProject.projectId,
    environmentId: testEnvironment.environmentId,
    apiToken: "9fc34a78-1e12-4453-ae87-055803d35715",
});

export const output = {
    project: myProject,
    environment: myEnvironment,
    service: myService,
};
