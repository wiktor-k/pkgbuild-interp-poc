# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Angel Velasquez <angvp@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Letu Ren <fantasquex@gmail.com>

_pyname=astroid
pkgname=python-$_pyname
pkgver=4.3.4
pkgrel=1
pkgdesc='A common base representation of python source code'
arch=(any)
url="https://github.com/pylint-dev/$_pyname"
license=(LGPL-2.1-or-later)
depends=(python)
makedepends=(python-{build,installer,wheel}
             python-setuptools)
checkdepends=(mypy
              python-attrs
              python-dateutil
              python-pyqt6
              python-regex
              python-urllib3
              python-pip
              python-pytest)
replaces=(python-logilab-astng)
conflicts=(python-logilab-astng)
_archive="$_pyname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz"
        "astroid-4.2.0-pkg_resources-removal.patch")
sha256sums=('8cd3fb0e084c89d69688d3aa4a1abcaf22c84a8167ebb06895bf5f335779a4f5'
            'ac45d2fb054f9b911866192e1a4843786fb8f1eff8d869c4d0f1dccde54465b4')

prepare() {
	cd "$_archive"
	# Skip tests requiring pkg_resources
	patch -Np1 -i ../astroid-4.2.0-pkg_resources-removal.patch
}

build() {
	cd "$_archive"
	python -m build -wn
}

check() {
	cd "$_archive"
	pytest
}

package() {
	cd "$_archive"
	python -m installer -d "$pkgdir" dist/*.whl
}
