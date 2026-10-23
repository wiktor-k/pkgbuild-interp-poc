# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>

pkgname=python-aws-sam-translator
_pkgname=serverless-application-model
pkgver=1.113.0
pkgrel=1
pkgdesc='AWS Serverless Application Model (AWS SAM) prescribes rules for expressing Serverless applications on AWS'
arch=(any)
url='https://github.com/aws/serverless-application-model'
license=('Apache-2.0')
depends=(
  python
  python-boto3
  python-botocore
  python-jsonschema
  python-pydantic
  python-typing_extensions
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-parameterized
  python-pytest
  python-pytest-rerunfailures
  python-pytest-xdist
  python-requests
  # samtranslator.yaml_helper is used by tests and unpackaged upstream scripts.
  python-yaml
)
optdepends=('python-yaml: YAML parsing with samtranslator.yaml_helper')
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('08a25b1a13025af10eaccd7ed0a3ee466c603d72c0276c83eab63b0d2222650fcc6a0c1ddcb769e9df7424e6432d2a8de9dea2db5770524a9b024952028805ae')

prepare() {
  cd $_pkgname-$pkgver
  # Avoid very verbose testing logs
  sed -i -r 's#log_cli\s*=\s*1#log_cli = 0#' pytest.ini
}

build() {
  cd $_pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    -W ignore::UserWarning
    -n auto
    --override-ini="addopts="
  )
  cd $_pkgname-$pkgver
  AWS_DEFAULT_REGION=us-east-1 PYTHONPATH="$PWD" pytest "${pytest_options[@]}" tests
}

package() {
  cd $_pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
