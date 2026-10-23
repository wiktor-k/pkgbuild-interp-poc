# Maintainer: BlackIkeEagle <ike DOT devolder AT gmail DOT com>
# Contributor: Philippe Cherel <philippe.cherel@mayenne.org>
# vim: ft=sh:

pkgname=libcec
pkgver=8.1.7
pkgrel=1
pkgdesc="Pulse-Eight's libcec for the Pulse-Eight USB-CEC adapter"
arch=('x86_64')
url="http://libcec.pulse-eight.com/"
license=('GPL-2.0-or-later')
makedepends=('cmake' 'swig' 'python')
depends=('udev' 'libxrandr')
optdepends=(
    'python: use cec in python applications'
)
source=("$pkgname-$pkgver.tar.gz::https://github.com/Pulse-Eight/$pkgname/archive/$pkgname-$pkgver.tar.gz")
sha256sums=('e4ac4d1dd3559cf83189d29ab40fc5537b5a2beceb82916ef173dac7e50dbfa0')

build() {
    cd "$pkgname-$pkgname-$pkgver"
    mkdir build
    cd build
    cmake .. \
        -DCMAKE_BUILD_TYPE=Release \
        -DBUILD_SHARED_LIBS=1 \
        -DDISABLE_STATIC=ON \
        -DHAVE_LINUX_API=1 \
        -DHAVE_EXYNOS_API=1 \
        -DHAVE_AOCEC_API=1 \
        -DCMAKE_INSTALL_PREFIX=/usr \
        -DCMAKE_INSTALL_LIBDIR=/usr/lib \
        -DCMAKE_INSTALL_LIBDIR_NOARCH=/usr/lib
    make
}

package() {
    cd "$pkgname-$pkgname-$pkgver/build"
    make DESTDIR="$pkgdir" install
}
