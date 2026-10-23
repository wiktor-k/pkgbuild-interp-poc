# Maintainer: David Runge <dvzrv@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

pkgname=python-nitrokey
_name="${pkgname#python-}-sdk-py"
pkgver=0.5.0
pkgrel=1
pkgdesc="Python SDK for Nitrokey devices"
arch=(any)
url="https://github.com/nitrokey/nitrokey-sdk-py"
license=('Apache-2.0 OR MIT')
depends=(
  python
  python-crcmod
  python-cryptography
  python-fido2
  python-hidapi
  python-protobuf
  python-pyserial
  python-requests
  python-semver
  python-tlv8
  python-typing_extensions
)
makedepends=(
  python-build
  python-installer
  python-poetry-core
)
source=($pkgname-$pkgver.tar.gz::$url/archive/refs/tags/v$pkgver.tar.gz)
sha512sums=('e002c5a71853a38a82f77d840d7658d7b06687214b4b5eef91a336c6e03fe2082c547a48e64343d4b306ff138a675294e3aea7fcb562935bd85a43dc73db11a7')
b2sums=('3e43f30aa8cae7dbdfe5d9baa449ab16f59c9024e354dd7daaf946948e26b0209044f46d95345d59f11fd51ad78ead8e3d2ccdde601b4af9deacc9f334568332')

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  cd $_name-$pkgver
  PYTHONPATH="src:$PYTHONPATH" python -m unittest discover tests
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSES/*.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
}
