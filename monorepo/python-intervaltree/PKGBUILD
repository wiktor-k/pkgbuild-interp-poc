# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: kpcyrd <kpcyrd[at]archlinux[dot]org>
# Contributor: Dobroslaw Kijowski [dobo] <dobo90_at_gmail.com>

pkgname=python-intervaltree
pkgver=3.2.1
pkgrel=1
pkgdesc='Library providing a mutable, self-balancing interval tree for Python'
arch=('any')
url="https://github.com/chaimleib/intervaltree"
license=('Apache-2.0')
depends=(
  'python'
  'python-sortedcontainers'
)
makedepends=(
  'python-build'
  'python-hatchling'
  'python-installer'
  'python-wheel'
)
checkdepends=('python-pytest')
source=(
  "$url/archive/$pkgver/$pkgname-$pkgver.tar.gz"
  "$pkgname-remove-uv-dynamic-version.patch"
)
b2sums=('7aa45f4d7908eee0d613e782bd4180e5e90111e9740f9cfc4025700ad0e6b3530cb9487275ce6959d5a0ea75fe74d1e9436e24c63f5abd4a4cf863a54381634c'
        '2cc99fdcc760e53579d3d2a603c8b78f85e21daeb09dac44431c4dc307355fd0237b173d162f68fbfceac59c870aad2730f31421ea46a328e39a93a560a267ea')

prepare() {
  cd ${pkgname#python-}-$pkgver
  patch -Np1 < ../$pkgname-remove-uv-dynamic-version.patch
  sed -i "s/@VERSION@/$pkgver/" pyproject.toml
}

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --no-isolation --skip-dependency-check
}

check() {
  cd ${pkgname#python-}-$pkgver
  pytest
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
}
