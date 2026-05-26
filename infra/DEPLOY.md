# Deploying baby-registry to DigitalOcean

This walks the first-time provisioning of a droplet for `storknest.baby`.

## Architecture

```
internet --443--> DO Load Balancer (TLS, HSTS) --<port>--> droplet --docker--> nginx -> backend
                                                                            -> SPA
```

A single DigitalOcean Load Balancer terminates TLS for every app on the
droplet. Each app publishes its frontend on a unique host port; the LB
forwards `host:443` to `droplet:<port>` based on the hostname. The backend
is never exposed off the docker network.

### Host port allocation

Keep this table in sync as new apps are added:

| App           | Hostname       | Host port |
| ------------- | -------------- | --------- |
| baby-registry | storknest.baby | 3035      |

## 1. Droplet

- Create an Ubuntu 24.04 droplet (1 GB RAM minimum, 2 GB recommended).
- Add your SSH key on creation.
- Put the droplet in a VPC so the LB can reach it on its private IP.

## 2. DNS

In the DigitalOcean DNS panel for `storknest.baby`:

- `A   storknest.baby      -> <load-balancer-ip>`
- `A   www.storknest.baby  -> <load-balancer-ip>`

(Don't point DNS at the droplet directly — point it at the LB.)

## 3. Server prep (run on the droplet)

```bash
adduser deploy
usermod -aG sudo deploy
rsync --archive --chown=deploy:deploy ~/.ssh /home/deploy

apt update && apt -y upgrade
apt -y install ca-certificates curl git ufw

install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg \
  -o /etc/apt/keyrings/docker.asc
chmod a+r /etc/apt/keyrings/docker.asc
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] \
  https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo $VERSION_CODENAME) stable" \
  > /etc/apt/sources.list.d/docker.list
apt update
apt -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
usermod -aG docker deploy

# Firewall: only allow SSH from anywhere; app ports are only reachable from
# the LB on the VPC private network, so we don't open them publicly.
ufw default deny incoming
ufw default allow outgoing
ufw allow OpenSSH
ufw --force enable
```

DigitalOcean Load Balancers reach the droplet over the VPC private network,
so the app port does not need to be open on the public firewall.

## 4. DigitalOcean Load Balancer

Create the LB in the same VPC as the droplet. For each app, add:

1. **Forwarding rule**: HTTPS `:443` -> HTTP `:<host port>` on the droplet.
2. **TLS certificate**: a DigitalOcean managed Let's Encrypt cert for the
   hostname (e.g. `storknest.baby` and `www.storknest.baby`).
3. **Health check**: HTTP `GET /` on the same host port. (nginx returns 200
   for `/index.html`.)
4. **Sticky sessions**: off.
5. **Redirect HTTP -> HTTPS**: on.
6. **Proxy Protocol**: off (keep `X-Forwarded-*` instead).
7. **HSTS**: enable on the LB once you've verified the site loads cleanly.

If you later run multiple apps on this droplet, add one forwarding rule per
host port. Hostname routing on a DO LB is done by separate rules per cert,
not via SNI multiplexing inside a single rule.

## 5. Secrets (sops + age)

On your workstation:

```bash
brew install sops age
age-keygen -o ~/.config/sops/age/keys.txt   # save the public key
```

Update [infra/.sops.yaml](.sops.yaml) with your age **public** key, then:

```bash
cd infra
cp secrets/baby-registry.env.example secrets/baby-registry.env
# fill in real values, then encrypt
sops --encrypt secrets/baby-registry.env > secrets/baby-registry.env.enc
git add .sops.yaml secrets/baby-registry.env.enc
```

On the droplet, install age + sops and place the **private** key at
`/home/deploy/.config/sops/age/keys.txt` (mode 0600). The deploy workflow
decrypts the encrypted env before bringing the stack up.

## 6. First deploy (manual, before CI is wired)

As `deploy` on the droplet:

```bash
git clone <repo> butterfeet
cd butterfeet/apps/baby-registry
sops --decrypt ../../infra/secrets/baby-registry.env.enc > .env
docker compose -f docker-compose.yml -f compose.prod.yml --env-file .env up -d --build
```

Sanity-check from the droplet itself:

```bash
curl -I http://localhost:3035/
```

Then from your laptop, once DNS + LB are wired:

```bash
curl -I https://storknest.baby
curl -I https://www.storknest.baby
```

## 7. CI/CD (GitHub Actions)

Two workflows live under `.github/workflows/`:

- `baby-registry-ci.yml` — on PR + push to `main`. Runs `go vet/build/test`,
  `npm ci && npm run build`, then builds the prod Docker images. Path-filtered
  to `apps/baby-registry/**` so other apps don't trigger it.
- `baby-registry-deploy.yml` — on push to `main` (and `workflow_dispatch`).
  Decrypts `infra/secrets/baby-registry.env.enc` with sops, scps the plaintext
  to the droplet, then SSHes in and re-runs `docker compose up -d --build`.

### Required GitHub secrets (Settings -> Secrets and variables -> Actions)

| Name                 | Value                                                                           |
| -------------------- | ------------------------------------------------------------------------------- |
| `SOPS_AGE_KEY`       | Contents of `~/.config/sops/age/keys.txt` (the **private** age key, full file). |
| `DEPLOY_SSH_KEY`     | Private SSH key (PEM) for the `deploy` user on the droplet.                     |
| `DEPLOY_HOST`        | Droplet public IP or hostname.                                                  |
| `DEPLOY_USER`        | `deploy`                                                                        |
| `DEPLOY_REPO_DIR`    | Absolute path of the checkout on the droplet, e.g. `/home/deploy/butterfeet`.   |
| `DEPLOY_BRANCH`      | (Optional) branch to deploy. Defaults to `main`.                                |

### Required GitHub variables (optional)

| Name                | Value                                                                |
| ------------------- | -------------------------------------------------------------------- |
| `BABY_REGISTRY_URL` | e.g. `https://storknest.baby` — used for the post-deploy smoke test. |

Generate the SSH keypair locally and add the **public** key to
`/home/deploy/.ssh/authorized_keys` on the droplet. Put the **private**
key in `DEPLOY_SSH_KEY`. The workflow runs `ssh-keyscan` against `DEPLOY_HOST`
at deploy time to pin the host key.

The deploy workflow uses a GitHub Environment called `production`, so you
can add a manual approval gate there if you want one.

## 8. Email (Resend)

In Resend:

1. Add domain `storknest.baby`.
2. Add the SPF, DKIM, and (optional) DMARC records to DigitalOcean DNS.
3. Verify the domain in Resend.
4. Generate an API key and put it in `infra/secrets/baby-registry.env`.

Trigger a sign-in from the live site and check the recipient inbox.
