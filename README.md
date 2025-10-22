# Paketo Ubi 8 Buildpackless Base Builder

## `docker.io/paketobuildpacks/builder-ubi8-buildpackless-base`

This builder uses the [UBI 8 Base Stack](https://github.com/paketo-buildpacks/ubi8-base-stack) (a Red Hat Ubi 8 base image) and contains **no buildpacks nor order groups**.
To use this builder, you must specify buildpacks and extensions at build time using whatever mechanisms your CNB platform of choice offers.

For example, with the `pack` CLI, use `--buildpack` or/and `--extension` as follows:

```
pack build nodejs-app \
    --path ./path/to/nodejs/app \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
    --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

To see which versions of build and run images and the lifecycle are contained within a given builder version, see the [Releases](https://github.com/paketo-buildpacks/builder-ubi8-buildpackless-base/releases) on this repo. This information is also available in the `builder.toml`.

For more information about this builder and how to use it, visit the [Paketo builder documentation](https://paketo.io/docs/builders/).

To learn about the base images included in this builder, visit the [Buildpacks base images](https://buildpacks.io/docs/for-app-developers/concepts/base-images/) and the [Paketo UBI 8 Base stack repo](https://github.com/paketo-buildpacks/ubi8-base-stack).
