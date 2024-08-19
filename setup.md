### What is Authelia?

Authelia is an open-source authentication and authorization server that provides multi-factor authentication (MFA) for securing applications and services. It is designed to act as a reverse proxy that sits in front of your services and enforces security policies before granting access. Authelia supports a variety of authentication methods, including password-based login, TOTP (Time-Based One-Time Password), and Duo Push.

Authelia is typically deployed in combination with a reverse proxy such as NGINX, Traefik, or Caddy, and it is capable of handling both single-factor and multi-factor authentication for web applications.

### Key Features of Authelia:

1. **Single Sign-On (SSO):** Allows users to sign in once and access multiple services.
2. **Multi-Factor Authentication (MFA):** Adds a second layer of security with options such as TOTP, WebAuthn, and Duo Push.
3. **Authorization Policies:** Define who can access what based on rules, IP addresses, groups, and more.
4. **Extensive Protocol Support:** Integrates with OIDC, LDAP, OAuth2, and more.
5. **Integrates with Reverse Proxies:** Works with NGINX, Traefik, Caddy, and more for seamless integration.

### How Authelia Works:

1. **Reverse Proxy Configuration:**
   - Authelia is deployed behind a reverse proxy, which intercepts all incoming requests to your services.
   - The reverse proxy redirects unauthenticated users to Authelia for authentication.

2. **Authentication Flow:**
   - **Initial Request:** When a user tries to access a protected resource, the reverse proxy forwards the request to Authelia.
   - **Authentication Check:** Authelia checks if the user is authenticated. If not, the user is redirected to the login portal.
   - **MFA (if enabled):** After entering the correct username and password, the user is prompted for MFA (e.g., TOTP).
   - **Authorization Check:** Once authenticated, Authelia checks authorization policies to determine if the user has access to the requested resource.
   - **Access Granted/Denied:** The user is granted access or redirected based on the policy decision.

3. **Authorization Policies:**
   - Authelia allows you to define fine-grained policies that specify who can access what services.
   - Policies can be based on IP addresses, user groups, or specific conditions.

### How to Set Up Authelia:

1. **Install Docker and Docker Compose:**
   - Authelia is often deployed using Docker. Ensure that Docker and Docker Compose are installed on your server.

   ```bash
   sudo apt update
   sudo apt install docker.io docker-compose -y
   ```

2. **Create the Directory Structure:**
   - Create a directory to store your Authelia configuration files.

   ```bash
   mkdir -p /etc/authelia
   cd /etc/authelia
   ```

3. **Create a Docker Compose File:**
   - In this directory, create a `docker-compose.yml` file for Authelia.

   ```yaml
   version: '3'
   services:
     authelia:
       image: authelia/authelia:latest
       container_name: authelia
       volumes:
         - ./config:/config
       ports:
         - 9091:9091
       restart: unless-stopped
   ```

4. **Create the Configuration File:**
   - Create a `configuration.yml` file inside the `config` directory to define your Authelia settings.

   **Example `configuration.yml`:**

   ```yaml
   server:
     host: 0.0.0.0
     port: 9091

   jwt_secret: your_jwt_secret

   default_redirection_url: https://coreledger.ca

   session:
     name: authelia_session
     secret: your_session_secret
     expiration: 3600s

   authentication_backend:
     file:
       path: /config/users_database.yml

   access_control:
     default_policy: deny
     rules:
       - domain: yourdomain.com
         policy: one_factor
   ```

5. **Set Up User Database:**
   - Create a `users_database.yml` file to store user credentials.

   **Example `users_database.yml`:**

   ```yaml
   users:
     your_username:
       password: $2y$12$hashed_password_here
       displayname: Your Name
       emails:
         - your_email@domain.com
   ```

   - Use bcrypt to hash passwords before adding them to this file.

6. **Configure Reverse Proxy:**
   - Configure your reverse proxy (e.g., NGINX, Traefik) to forward authentication requests to Authelia.

   **Example NGINX Configuration:**

   ```nginx
   server {
     listen 80;
     server_name yourdomain.com;

     location / {
       proxy_pass http://your_backend_service;
       include /etc/nginx/proxy_params;
       auth_request /auth-verify;
     }

     location /auth-verify {
       internal;
       proxy_pass http://authelia:9091/api/verify;
       proxy_set_header X-Original-URI $request_uri;
     }
   }
   ```

7. **Start Authelia:**
   - Start Authelia using Docker Compose.

   ```bash
   docker-compose up -d
   ```

### Summary of Key Points for Your Team:

- **Authelia** acts as an authentication and authorization server, providing multi-factor authentication (MFA) and enforcing security policies.
- **Deployment:** It is deployed behind a reverse proxy like NGINX or Traefik and works by intercepting requests to your services.
- **Configuration:** You'll configure Authelia via a YAML file, setting up authentication methods, user databases, and access control policies.
- **Integration:** Authelia integrates with your backend services via the reverse proxy, ensuring that only authenticated and authorized users can access them.

Authelia's official documentation provides comprehensive guidance on setting it up, including configuration details, deployment strategies, and integration with various reverse proxies: [Authelia Documentation](https://www.authelia.com/docs/).
