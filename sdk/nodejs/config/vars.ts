// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("railway");

export declare const itsasecret: boolean | undefined;
Object.defineProperty(exports, "itsasecret", {
    get() {
        return __config.getObject<boolean>("itsasecret");
    },
    enumerable: true,
});

