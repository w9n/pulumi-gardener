// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Gardener_shoot extends pulumi.CustomResource {
    /**
     * Get an existing Gardener_shoot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Gardener_shootState, opts?: pulumi.CustomResourceOptions): Gardener_shoot {
        return new Gardener_shoot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gardener:index/gardener_shoot:gardener_shoot';

    /**
     * Returns true if the given object is an instance of Gardener_shoot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gardener_shoot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gardener_shoot.__pulumiType;
    }

    /**
     * Standard shoot's metadata. More info:
     * https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.Gardener_shootMetadata>;
    /**
     * ShootSpec is the specification of a Shoot. (see
     * https://github.com/gardener/gardener/blob/master/pkg/apis/garden/v1beta1/types.go#L609)
     */
    public readonly spec!: pulumi.Output<outputs.Gardener_shootSpec>;

    /**
     * Create a Gardener_shoot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Gardener_shootArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Gardener_shootArgs | Gardener_shootState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Gardener_shootState | undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["spec"] = state ? state.spec : undefined;
        } else {
            const args = argsOrState as Gardener_shootArgs | undefined;
            if ((!args || args.metadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadata'");
            }
            if ((!args || args.spec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spec'");
            }
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["spec"] = args ? args.spec : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Gardener_shoot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering gardener_shoot resources.
 */
export interface Gardener_shootState {
    /**
     * Standard shoot's metadata. More info:
     * https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
     */
    metadata?: pulumi.Input<inputs.Gardener_shootMetadata>;
    /**
     * ShootSpec is the specification of a Shoot. (see
     * https://github.com/gardener/gardener/blob/master/pkg/apis/garden/v1beta1/types.go#L609)
     */
    spec?: pulumi.Input<inputs.Gardener_shootSpec>;
}

/**
 * The set of arguments for constructing a Gardener_shoot resource.
 */
export interface Gardener_shootArgs {
    /**
     * Standard shoot's metadata. More info:
     * https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
     */
    metadata: pulumi.Input<inputs.Gardener_shootMetadata>;
    /**
     * ShootSpec is the specification of a Shoot. (see
     * https://github.com/gardener/gardener/blob/master/pkg/apis/garden/v1beta1/types.go#L609)
     */
    spec: pulumi.Input<inputs.Gardener_shootSpec>;
}