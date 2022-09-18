## Required env files and values

For UI submodule, a `.env.production` file is required, and should contain the following. These values are used at compile time. This should be placed in the bbdb root directory

```dotenv
VITE_POCKETBASE_URL=https://your-url-here.com
VITE_VIEWER_USER=user@your-url-here.com
VITE_VERSION=0.1.0
```

## Build Instructions

For local dev, run the following. A `.env` file is required and should contain the same variables mentioned above.

```bash
./build.sh
```