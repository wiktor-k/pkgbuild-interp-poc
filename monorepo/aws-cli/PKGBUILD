# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Chih-Hsuan Yen <yan12125@archlinux.org>
# Contributor: Jonathan Steel <jsteel at archlinux.org>
# Contributor: Chris Severance aur.severach AatT spamgourmet.com
# Contributor: Alper KANAT <alperkanat@raptiye.org>

pkgname=aws-cli
pkgver=1.46.1
pkgrel=1
pkgdesc='Universal Command Line Interface for Amazon Web Services'
arch=('any')
url="https://github.com/aws/aws-cli"
license=('Apache-2.0')
depends=(
  'python'
  'python-certifi'
  'python-colorama'
  'python-dateutil'
  'python-docutils'
  'python-jmespath'
  'python-pyasn1'
  'python-rsa'
  'python-urllib3'
  'python-yaml'
)
makedepends=(
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
checkdepends=(
  'procps-ng'
  'python-awscrt'
  'python-jsonschema'
  'python-pytest'
  'python-pytest-xdist'
)
optdepends=('python-awscrt: use AWS Common Runtime')
source=("$url/archive/$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('3120c6768b749a85eea847040758f64120d5fba33be786566e178c36ed53d5d81676b8f9d632be7c09e65cc469d71366f897a4cafeb237b8fd2128361e3b4b1d')

build() {
  cd $pkgname-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $pkgname-$pkgver
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # Many integration tests need real credentials
  PATH=$PWD/test-env/bin:$PATH test-env/bin/python -m pytest tests -n auto \
    --ignore=tests/integration
}

package() {
  cd $pkgname-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm644 bin/aws_bash_completer "$pkgdir/usr/share/bash-completion/completions/aws"
  rm -v "$pkgdir/usr/bin/"{aws.cmd,aws_bash_completer}
}
