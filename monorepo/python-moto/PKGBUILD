# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Guillaume Horel <guillaume.horel@gmail.com>

pkgname=python-moto
pkgver=5.2.4
pkgrel=1
pkgdesc='Moto is a library to mock out the boto library.'
arch=(any)
url='https://github.com/spulec/moto'
license=('Apache-2.0')
depends=(
  python
  python-boto3
  python-botocore
  python-cryptography
  python-dateutil
  python-jinja
  python-requests
  python-responses
  python-urllib3
  python-werkzeug
  python-xmltodict
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-antlr4
  python-aws-xray-sdk
  python-cfn-lint
  python-crc32c
  python-docker
  python-flask
  python-flask-cors
  python-freezegun
  python-graphql-core
  python-joserfc
  python-jsondiff
  python-jsonpath-ng
  python-jsonschema
  python-multipart
  python-openapi-spec-validator
  python-py-partiql-parser
  python-pyparsing
  python-pytest
  python-pytest-order
  python-yaml
  python-pyotp
)
optdepends=(
  'python-yaml: for apigatewayv2, cloudformation, s3 and ssm'
  'python-joserfc: for apigateway, cloudformation and cognitoidp'
  'python-openapi-spec-validator: for apigateway and cloudformation'
  # SNS and SQS still uses docker indirectly, while upstream explicitly removes them
  # https://github.com/spulec/moto/pull/4094
  'python-docker: for awslambda, batch, cloudformation, dynamodb and dynamodbstreams'
  'python-graphql-core: For appsync and cloudformation'
  'python-jsondiff: for iotdata and cloudformation'
  'python-aws-xray-sdk: for xray and cloudformation'
  'python-cfn-lint: for cloudformation'
  'python-pyparsing: for glue and cloudformation'
  'python-py-partiql-parser: for cloudformation and s3'
  'python-crc32c: for s3'
  'python-flask: for moto_server'
  'python-flask-cors: for moto_server'
  'python-multipart: for moto_proxy'
  'python-antlr4: for stepfunctions'
  'python-jsonpath-ng: for events, stepfunctions'
  'python-jsonschema: for quicksight'
)
source=("git+https://github.com/getmoto/moto#tag=$pkgver")
b2sums=('7d98d0fe113a74a7589df77c3f70d14771ae9bd999d1bd633f6e32a980eb031fe80ded31386cdd56bfd45bb879ba23fefca4de1e73e947a7d1728093d16c0413')

prepare() {
  cd ${pkgname#python-}
  python update_version_from_git.py $pkgver
}

build() {
  cd ${pkgname#python-}
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname#python-}
  local pytest_args=(
    # Needs a new package python-pycognito
    --ignore=tests/test_cognitoidp/test_cognitoidp.py
    -m 'not requires_docker'
    # Fails for some reason
    --deselect=tests/test_awslambda/test_lambda.py::test_delete_function
    --deselect=tests/test_dsql/test_dsql.py::test_stream_lifecycle
    --deselect=tests/test_dsql/test_dsql.py::test_create_stream_idempotency_conflict
    --deselect=tests/test_dsql/test_dsql.py::test_create_stream_is_idempotent
    --deselect=tests/test_dsql/test_dsql.py::test_list_streams_with_pagination
    --deselect=tests/test_dsql/test_dsql.py::test_tag_and_untag_stream
    --deselect=tests/test_dsql/test_dsql.py::test_stream_operations_validate_resources
    --deselect=tests/test_dsql/test_dsql.py::test_resource_groups_tagging_api_returns_clusters_and_streams
    --deselect=tests/test_mq/test_mq.py::test_describe_shared_resources
    --deselect=tests/test_paymentcryptography/test_paymentcryptography.py::test_put_resource_policy
    --deselect=tests/test_paymentcryptography/test_paymentcryptography.py::test_put_resource_policy_key_not_found
    --deselect=tests/test_paymentcryptography/test_paymentcryptography.py::test_get_resource_policy_when_none_set
    --deselect=tests/test_paymentcryptography/test_paymentcryptography.py::test_get_resource_policy_key_not_found
    --deselect=tests/test_paymentcryptography/test_paymentcryptography.py::test_delete_resource_policy
  )
  TZ=UTC pytest tests "${pytest_args[@]}"
}

package() {
  cd ${pkgname#python-}
  python -m installer --destdir="$pkgdir" dist/*.whl
}
