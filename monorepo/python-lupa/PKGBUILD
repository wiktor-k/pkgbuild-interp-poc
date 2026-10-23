# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Eli Schwartz <eschwartz@archlinux.org>

_pkgname=lupa
pkgname=python-lupa
pkgver=2.8
pkgrel=1
pkgdesc='Python wrapper around Lua and LuaJIT'
arch=(x86_64)
url="https://github.com/scoder/$_pkgname"
license=(MIT)
depends=(gcc-libs # libgcc_s.so
         glibc # libc.so
         python)
makedepends=(cython
             python-{build,installer,wheel}
             python-setuptools)
_archive="$_pkgname-$pkgver"
source=("https://files.pythonhosted.org/packages/source/${_pkgname::1}/$_pkgname/$_archive.tar.gz")
sha256sums=('d8022641b9ec8ecf2c5ecbe9f47e5a70e0b87c4b5ae921b92cb02a638e0acd08')

build() {
	cd "$_archive"
	python -m build -wn
	# https://gitlab.archlinux.org/archlinux/packaging/packages/python-lupa/-/issues/1
	local _pyver=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
	cp -r lupa/tests "build/lib.linux-$CARCH-cpython-$_pyver/lupa"
}

check() {
	local _pyver=$(python -c 'import sys; print("".join(map(str, sys.version_info[:2])))')
	export PYTHONPATH="$PWD/$_archive/build/lib.linux-$CARCH-cpython-$_pyver"
	python -m unittest discover "$_archive/$_pkgname/tests"
}

package() {
	cd "$_archive"
	python -m installer -d "$pkgdir" dist/*.whl
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE.txt
}
