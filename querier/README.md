# Querier

Querier is responsible for performing ElasticSearch queries. In short, this app:

- provides a HTTP endpoint to ingest user-click events.
- supports AWS ElasticSearch authentication using signed-requests using Lambda Roles.
- integrates with Terraform-managed infrastructure using SSM Parameter Store.
- is managed using Serverless Framework.

### The Stack

1. **App:** Python with Serverless Framework
2. **Code Standard and Quality**: `autopep8`
3. **Unit Testing:** `unittest`
4. **Logging:** `jsonlogger`
5. **ES Client:** `elasticsearch`
6. **CI/CD:** GitHub Action to automatically test, build, and deploy on mainline branch.

### Developing

**Requirements:**

1. Python 3.8 or newer
2. [AutoPEP8](https://pypi.org/project/autopep8/) (code linting)
3. [Nose](https://pypi.org/project/nose/) (unit test helper)

**Code Formatting:**

```bash
# Run autopep8 is you do not have autopep8 integration with your IDE

$ autopep8 --in-place *.py
```

**Running Tests:**

```bash
# Install dependencies
pip install -r requirements.txt


# Without nosetests

$ python -m unittest handler_test.py


# Using nosetests

$ nosetests .
```

### Deploying

Requirements:

1. [Serverless Framework 2.x](https://www.serverless.com/framework/docs/providers/aws/guide/installation/)

Steps:

```bash
sls deploy --region=<REGION> --stage=<STAGE>
```

### Environment Variables

| Environment | Type     | Description                                      | Default Value |
| ----------- | -------- | ------------------------------------------------ | ------------- |
| `REGION`    | _string_ | AWS Region (used to sign ElasticSearch requests) |               |
| `ES_HOST`   | _string_ | ElasticSearch Host FQDN                          |               |
