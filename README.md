#vndr-updater
Small tool to update [VNDR](https://github.com/LK4D4/vndr) `vendor.conf` file versions from remote vendor.conf.
You can use this tool to update local versions to match with remote versions.

#### Example
`go run ./cmd/vndr-updater/ https://raw.githubusercontent.com/containerd/containerd/v1.1.0/vendor.conf`

Update all dependencies in local `vendor.conf` to match with versions in the remote `vendor.conf` file.

