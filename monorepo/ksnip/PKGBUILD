# Maintainer : Daniel Bermond <dbermond@archlinux.org>
# Contributor: Balló György <ballogyor+arch at gmail dot com>
# Contributor: dracorp aka Piotr Rogoza <piotr.r.public at gmail.com>

pkgname=ksnip
pkgver=1.10.1
pkgrel=6
pkgdesc='Qt-based screenshot tool that provides many annotation features'
arch=('x86_64')
url='https://github.com/ksnip/ksnip/'
license=('GPL-3.0-or-later')
depends=(
    'glibc'
    'hicolor-icon-theme'
    'kimageannotator'
    'libgcc'
    'libstdc++'
    'libx11'
    'libxcb'
    'qt6-base')
makedepends=(
    'extra-cmake-modules'
    'git'
    'qt6-tools')
source=("git+https://github.com/ksnip/ksnip#tag=v${pkgver}"
        '010-ksnip-qt6-fix.patch')
sha256sums=('ef434afdef93f93690dea1f1bc2522b443b490fd87c5f3211110908174638bef'
            '8c15ae9d3740bb9ae163dc3e877ca155aa4ed0381b7c4d9b58263b05e37a0249')

prepare() {
    git -C ksnip cherry-pick -n 76f4b381971eead6ff31b8bf3bb64bb5717469c3 # Fix build with kimageannotator 0.7.1
    git -C ksnip cherry-pick -n 82499f6a8b3483f17fd74cc0e1293d82db276100 # Port to Qt6
    patch -d ksnip -Np1 -i "${srcdir}/010-ksnip-qt6-fix.patch"
}

build() {
    cmake -B build -S ksnip \
        -G 'Unix Makefiles' \
        -DCMAKE_BUILD_TYPE:STRING='None' \
        -DCMAKE_INSTALL_PREFIX:PATH='/usr' \
        -DBUILD_WITH_QT6:BOOL='ON' \
        -Wno-dev
    cmake --build build
}

package() {
    DESTDIR="$pkgdir" cmake --install build
}
