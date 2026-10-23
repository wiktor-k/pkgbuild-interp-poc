# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Aleksandr Beliaev <trap000d@gmail.com>

_pkgname=zipstream-ng
pkgname=python-$_pkgname
pkgver=1.9.3
pkgrel=1
pkgdesc='A modern and easy to use streamable zip file generator'
arch=(any)
url="https://github.com/pR0Ps/$_pkgname"
license=(LGPL-3.0-only)
depends=(python)
provides=("python-zipstream=$pkgver")
conflicts=(python-zipstream)
makedepends=(python-{build,installer,wheel}
             python-hatchling)
_archive="$_pkgname-$pkgver"
source=("$url/archive/v$pkgver/$_archive.tar.gz")
sha256sums=('f1781c8d4110b30cb1d2b774d28f4a3e264698ed29be7aa37804b11cdee20da7')

build() {
	cd "$_archive"
	python -m build -wn
}

package() {
	cd "$_archive"
	python -m installer -d "$pkgdir" dist/*.whl
}
