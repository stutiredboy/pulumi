// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAssets(args: GetAssetsArgs, opts?: pulumi.InvokeOptions): Promise<GetAssetsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("example::GetAssets", {
        "archive": args.archive,
        "source": args.source,
    }, opts);
}

export interface GetAssetsArgs {
    archive: pulumi.asset.Archive;
    source: pulumi.asset.Asset | pulumi.asset.Archive;
}

export interface GetAssetsResult {
    readonly archive: pulumi.asset.Archive;
    readonly source: pulumi.asset.Asset | pulumi.asset.Archive;
}
export function getAssetsOutput(args: GetAssetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssetsResult> {
    return pulumi.output(args).apply((a: any) => getAssets(a, opts))
}

export interface GetAssetsOutputArgs {
    archive: pulumi.Input<pulumi.asset.Archive>;
    source: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
}
