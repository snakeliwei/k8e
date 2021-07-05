#!/bin/sh


# --- helper functions for logs ---
info()
{
    echo '[INFO] ' "$@"
}
warn()
{
    echo '[WARN] ' "$@" >&2
}
fatal()
{
    echo '[ERROR] ' "$@" >&2
    exit 1
}

BIN_DIR=/usr/local/bin

# --- nerdctl containerd run sock path ---
export CONTAINERD_ADDRESS=/run/k8e/containerd/containerd.sock

# --- add additional utility links ---
create_symlinks() {

    for cmd in kubectl crictl ctr; do
        if [ ! -e ${BIN_DIR}/${cmd} ]; then
            which_cmd=$(which ${cmd} 2>/dev/null || true)
            if [ -z "${which_cmd}" ]; then
                info "Creating ${BIN_DIR}/${cmd} symlink to k8e"
                $SUDO ln -sf /opt/k8e/k8e ${BIN_DIR}/${cmd}
            else
                info "Skipping ${BIN_DIR}/${cmd} symlink to k8e, command exists in PATH at ${which_cmd}"
            fi
        else
            info "Skipping ${BIN_DIR}/${cmd} symlink to k8e, already exists"
        fi
    done
    info "Create nerdctl symlink for k8e"
    $SUDO ln -sf /var/lib/k8e/k8e/data/current/bin/nerdctl ${BIN_DIR}/nerdctl
}

(
create_symlinks
)
