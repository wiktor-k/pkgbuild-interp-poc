# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
pkgname=vncdotool
pkgver=1.4.2
pkgrel=1
pkgdesc='A command line VNC client'
arch=('any')
url='https://github.com/sibson/vncdotool'
license=('MIT')
depends=('python-pillow' 'python-twisted' 'python-pycryptodomex')
makedepends=('python-build' 'python-installer' 'python-setuptools' 'python-wheel')
source=($pkgname-$pkgver.tar.gz::https://github.com/sibson/vncdotool/archive/v$pkgver.tar.gz)
sha512sums=('7ca83a661a1d034d684ed77fc9a67a13e6f1321c393676d53be386b2b2d65bad173fd6d567f99335dfa1f351a3dda4df509c6f3d60b6f090d63e202ca90addfa')

build() {
    cd $pkgname-$pkgver
    python -m build --wheel --no-isolation
}

package() {
    cd $pkgname-$pkgver

    python -m installer --destdir="$pkgdir" dist/*.whl
    install -Dm 644 LICENSE.txt "$pkgdir/usr/share/licenses/$pkgname/LICENSE.txt"
}
