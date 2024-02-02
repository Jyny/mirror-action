# mirror action
A GitHub Action for mirroring a git repository as source to another git repository as destination.
Support both src/dst from private repository with ssh key or access token.

## Example workflows

### Mirror from private repository to private repository with ssh key
```yaml
name: example_workflow
on:
  workflow_dispatch:

jobs:
  example_job:
    runs-on: ubuntu-latest
    steps:
      - uses: jyny/mirror-action@v0.0.1
        with:
            SRC_REMOTE_URL: ${{ vars.SRC_REMOTE_URL }}
            SRC_SSH_KEY: ${{ secrets.SRC_SSH_KEY }}
            SRC_IGNORE_HOST_KEY: true
            DST_REMOTE_URL: ${{ vars.DST_REMOTE_URL }}
            DST_SSH_KEY: ${{ secrets.DST_SSH_KEY }}
            DST_IGNORE_HOST_KEY: true
```